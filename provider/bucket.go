// Copyright 2023, OSMIT GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/ceph/go-ceph/rgw/admin"
	"github.com/dustin/go-humanize"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"

	p "github.com/pulumi/pulumi-go-provider"
)

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type Bucket struct{}

// Each resource has in input struct, defining what arguments it accepts.
type QuotaArgs struct {
	Enabled       *bool   `pulumi:"enabled,optional"`
	MaxSizeHum    *string `pulumi:"maxSizeHum,optional"`
	MaxSize       *int64  `pulumi:"maxSize,optional"`
	MaxObjectsHum *string `pulumi:"maxObjectsHum,optional"`
	MaxObjects    *int64  `pulumi:"maxObjects,optional"`
}

// Annotate the nested fields for the Metadata receiver
func (m *QuotaArgs) Annotate(a infer.Annotator) {
	a.Describe(&m.Enabled, "Enable or disable quota enforcement quota")
	a.Describe(&m.MaxSize, "Maximum size - numeric value (alternate to MaxSizeHum)")
	a.Describe(&m.MaxSizeHum, "Maximum size - human readable format (alternate to MaxSize)")
	a.Describe(&m.MaxObjects, "Maximum object count - numeric value (alternate to MaxObjectsHum)")
	a.Describe(&m.MaxObjectsHum, "Maximum object count - human readable format (i.e. 10k) (alternate to MaxObjects)")
}

// Each resource has in input struct, defining what arguments it accepts.
type BucketArgs struct {
	// Bucket name
	Name          string    `pulumi:"name"`
	ObjectLocking bool      `pulumi:"objectLocking,optional"`
	Versioning    bool      `pulumi:"versioning,optional"`
	Quota         QuotaArgs `pulumi:"quota,optional"`
	PurgeOnDelete bool      `pulumi:"purgeOnDelete,optional"`
}

// Annotate the nested fields for the Metadata receiver
func (m *BucketArgs) Annotate(a infer.Annotator) {
	a.Describe(&m.Name, "Bucket name")
	a.Describe(&m.ObjectLocking, "Bucket object locking enabled")
	a.Describe(&m.Quota, "Bucket quota configuration")
	a.Describe(&m.PurgeOnDelete, "Purge bucket on delete")
}

// Each resource has a state, describing the fields that exist on the created resource.
type BucketState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	BucketArgs

	UBID string `pulumi:"ubid"`
	// this config element was assimilated...
	// Required: true
	Assimilated bool `pulumi:"_assimilated"`

	Location string `pulumi:"_location"`
}

// Annotate the nested fields for the Metadata receiver
func (m *BucketState) Annotate(a infer.Annotator) {
	a.Describe(&m.UBID, "The unique bucket id.")
	a.Describe(&m.Assimilated, "Bucket was 'assimilated' - managing an existing bucket.")
	a.Describe(&m.PurgeOnDelete, "Purge bucket on delete.")
}

func ToBucketState(input BucketArgs, olds BucketState) BucketState {
	return BucketState{BucketArgs: input, Assimilated: olds.Assimilated}
}

// All resources must implement Create at a minumum.
func (thiz *Bucket) Create(ctx p.Context, name string, input BucketArgs, preview bool) (string, BucketState, error) {
	// bail out now when we are in preview mode
	if preview {
		return IdPreviewPrefix + name, BucketState{
			BucketArgs: input,
		}, nil
	}

	retErr := func(err error) (string, BucketState, error) {
		return "", BucketState{BucketArgs: input}, err
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}

	// Configure CreateBucketInput
	s3req := &s3.CreateBucketInput{
		Bucket:                     aws.String(input.Name),
		ObjectLockEnabledForBucket: *aws.Bool(input.ObjectLocking),
	}

	_, err = ce.s3.CreateBucket(ctx, s3req)
	if err != nil {
		return retErr(err)
	}

	state, err := UpdateBucket(ctx, ce, input, false)

	return name, state, err
}

func (*Bucket) Diff(ctx p.Context, id string, olds BucketState, news BucketArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if news.Name != olds.Name {
		diff["name"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if news.ObjectLocking != olds.ObjectLocking {
		diff["objectLocking"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	// if news.DisplayName != olds.DisplayName {
	// 	diff["displayName"] = p.PropertyDiff{Kind: p.Update}
	// }
	// if news.Email != olds.Email {
	// 	diff["email"] = p.PropertyDiff{Kind: p.Update}
	// }
	if (olds.Quota.Enabled == nil) && (news.Quota.Enabled == nil) {
		// noop
	} else if (olds.Quota.Enabled == nil) && (news.Quota.Enabled != nil) {
		diff["quota.enabled"] = p.PropertyDiff{Kind: p.Add}
	} else if (olds.Quota.Enabled != nil) && (news.Quota.Enabled == nil) {
		diff["quota.enabled"] = p.PropertyDiff{Kind: p.Delete}
	} else if *olds.Quota.Enabled != *news.Quota.Enabled {
		diff["quota.enabled"] = p.PropertyDiff{Kind: p.Update}
	}
	if (olds.Quota.MaxObjects == nil) && (news.Quota.MaxObjects == nil) {
		// noop
	} else if (olds.Quota.MaxObjects == nil) && (news.Quota.MaxObjects != nil) {
		diff["quota.maxObjects"] = p.PropertyDiff{Kind: p.Add}
	} else if (olds.Quota.MaxObjects != nil) && (news.Quota.MaxObjects == nil) {
		diff["quota.maxObjects"] = p.PropertyDiff{Kind: p.Delete}
	} else if *olds.Quota.MaxObjects != *news.Quota.MaxObjects {
		diff["quota.maxObjects"] = p.PropertyDiff{Kind: p.Update}
	}
	if (olds.Quota.MaxObjectsHum == nil) && (news.Quota.MaxObjectsHum == nil) {
		// noop
	} else if (olds.Quota.MaxObjectsHum == nil) && (news.Quota.MaxObjectsHum != nil) {
		diff["quota.maxObjectsHum"] = p.PropertyDiff{Kind: p.Add}
	} else if (olds.Quota.MaxObjectsHum != nil) && (news.Quota.MaxObjectsHum == nil) {
		diff["quota.maxObjectsHum"] = p.PropertyDiff{Kind: p.Delete}
	} else if *olds.Quota.MaxObjectsHum != *news.Quota.MaxObjectsHum {
		diff["quota.maxObjectsHum"] = p.PropertyDiff{Kind: p.Update}
	}
	if (olds.Quota.MaxSize == nil) && (news.Quota.MaxSize == nil) {
		// noop
	} else if (olds.Quota.MaxSize == nil) && (news.Quota.MaxSize != nil) {
		diff["quota.maxSize"] = p.PropertyDiff{Kind: p.Add}
	} else if (olds.Quota.MaxSize != nil) && (news.Quota.MaxSize == nil) {
		diff["quota.maxSize"] = p.PropertyDiff{Kind: p.Delete}
	} else if *olds.Quota.MaxSize != *news.Quota.MaxSize {
		diff["quota.maxSize"] = p.PropertyDiff{Kind: p.Update}
	}
	if (olds.Quota.MaxSizeHum == nil) && (news.Quota.MaxSizeHum == nil) {
		// noop
	} else if (olds.Quota.MaxSizeHum == nil) && (news.Quota.MaxSizeHum != nil) {
		diff["quota.maxSizeHum"] = p.PropertyDiff{Kind: p.Add}
	} else if (olds.Quota.MaxSizeHum != nil) && (news.Quota.MaxSizeHum == nil) {
		diff["quota.maxSizeHum"] = p.PropertyDiff{Kind: p.Delete}
	} else if *olds.Quota.MaxSizeHum != *news.Quota.MaxSizeHum {
		diff["quota.maxSizeHum"] = p.PropertyDiff{Kind: p.Update}
	}
	// diffWalk(ctx, diff, "quota", reflect.ValueOf(olds.Quota), reflect.ValueOf(news.Quota))

	if len(diff) > 0 {
		ctx.Log(diag.Info, fmt.Sprintf("DIFF on Bucket %s/%s: Found %d diffs: %v", news.Name, id, len(diff), diff))
	}
	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Bucket) Read(ctx p.Context, id string, inputs BucketArgs, state BucketState) (string, BucketArgs, BucketState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return id, inputs, state, err
	}
	ret, err := ReadBucketState(ctx, ce, inputs, state.Assimilated)
	return id, inputs, ret, err
}

func UpdateBucket(ctx p.Context, ce *CacheEntry, input BucketArgs, assimilated bool) (BucketState, error) {
	retErr := func(err error) (BucketState, error) {
		return BucketState{BucketArgs: input, Assimilated: assimilated}, err
	}

	// read the bucket info - we need the owner
	info, err := ce.client.GetBucketInfo(ctx, admin.Bucket{Bucket: input.Name})
	if err != nil {
		return retErr(err)
	}

	var enabled bool = false
	var quotaBytes int64 = -1
	var quotaObjects int64 = -1
	if input.Quota.MaxSize != nil {
		quotaBytes = *input.Quota.MaxSize
		enabled = true
	} else if input.Quota.MaxSizeHum != nil && *input.Quota.MaxSizeHum != "" {
		quotaBytes64, err := humanize.ParseBytes(*input.Quota.MaxSizeHum)
		if err != nil {
			return retErr(err)
		}
		quotaBytes = int64(quotaBytes64)
		enabled = true
	}

	if input.Quota.MaxObjects != nil {
		quotaObjects = *input.Quota.MaxObjects
		enabled = true
	} else if input.Quota.MaxObjectsHum != nil && *input.Quota.MaxObjectsHum != "" {
		q64, unit, err := humanize.ParseSI(*input.Quota.MaxObjectsHum)
		if err != nil {
			return retErr(err)
		}
		if unit != "" {
			return retErr(fmt.Errorf("Error parsing quota.maxObjectsHum='%s': Unexpected unit '%s'", *input.Quota.MaxObjectsHum, unit))
		}
		quotaObjects = int64(q64)
		enabled = true
	}

	// overwrite with enabled value
	enabled = iftfe[bool](input.Quota.Enabled != nil, func() bool { return *input.Quota.Enabled }, enabled)
	quota := admin.QuotaSpec{
		UID:        info.Owner,
		Bucket:     input.Name,
		QuotaType:  "bucket",
		Enabled:    &enabled,
		MaxSize:    &quotaBytes,
		MaxObjects: &quotaObjects,
	}
	err = ce.client.SetIndividualBucketQuota(ctx, quota)
	if err != nil {
		return retErr(err)
	}

	// // Configure Bucket Locking
	// var s3LockConfig *s3types.ObjectLockConfiguration
	// if input.ObjectLocking {
	// 	s3LockConfig = &s3types.ObjectLockConfiguration{
	// 		ObjectLockEnabled: s3types.ObjectLockEnabledEnabled,
	// 	}
	// }
	// s3LockReq := &s3.PutObjectLockConfigurationInput{
	// 	Bucket:                  aws.String(input.Name),
	// 	ObjectLockConfiguration: s3LockConfig,
	// }
	// _, err = ce.s3.PutObjectLockConfiguration(ctx, s3LockReq)
	// if err != nil {
	// 	return retErr(err)
	// }

	// configure versioning
	s3VersReq := &s3.PutBucketVersioningInput{
		Bucket: aws.String(input.Name),
		VersioningConfiguration: &s3types.VersioningConfiguration{
			Status: ifte[s3types.BucketVersioningStatus](input.Versioning, s3types.BucketVersioningStatusEnabled, s3types.BucketVersioningStatusSuspended),
		},
	}
	_, err = ce.s3.PutBucketVersioning(ctx, s3VersReq)
	if err != nil {
		return retErr(err)
	}

	return ReadBucketState(ctx, ce, input, assimilated)
}

func ReadBucketState(ctx p.Context, ce *CacheEntry, input BucketArgs, assimilated bool) (BucketState, error) {
	// // Create Head Bucket Request
	// s3req := &s3.HeadBucketInput{
	// 	Bucket: aws.String(id),
	// }
	// // it just checks if the bucket exists and is accessible...
	// _, err := ce.s3.HeadBucket(ctx, s3req)
	// if err != nil {
	// 	// var ae smithy.APIError
	// 	// if errors.As(err, &ae) {
	// 	// 	switch ae.ErrorCode() {
	// 	// 	case "404":
	// 	// 		// deleted
	// 	// 		return
	// 	// 	case "403":
	// 	// 		// no permission to head bucket", err.Error()
	// 	// 		return
	// 	// 	}
	// 	// }
	// 	return BucketState{BucketArgs: input, Assimilated: assimilated}, err
	// }
	// ctx.Log(diag.Info, fmt.Sprintf("GetBuicketInfo %s", input.Name))
	info, err := ce.client.GetBucketInfo(ctx, admin.Bucket{Bucket: input.Name})
	if err != nil {
		return BucketState{BucketArgs: input, Assimilated: assimilated}, err
	}
	// ctx.Log(diag.Info, fmt.Sprintf("GetBuicketInfo id: %s owner: %s tenant: %s", info.ID, info.Owner, info.Tenant))
	state := BucketState{BucketArgs: input, UBID: info.ID, Assimilated: assimilated}
	state.Quota.MaxSize = info.BucketQuota.MaxSize
	if state.Quota.MaxSize != nil {
		if *state.Quota.MaxSize < 0 {
			state.Quota.MaxSize = input.Quota.MaxSize
		} else {
			str := humanize.SI(float64(*state.Quota.MaxSize), "B")
			state.Quota.MaxSizeHum = &str
		}
	}
	state.Quota.MaxObjects = info.BucketQuota.MaxObjects
	if state.Quota.MaxObjects != nil {
		if *state.Quota.MaxObjects < 0 {
			state.Quota.MaxObjects = input.Quota.MaxObjects
		} else {
			str := humanize.SI(float64(*state.Quota.MaxObjects), "")
			state.Quota.MaxObjectsHum = &str
		}
	}
	state.Quota.Enabled = info.BucketQuota.Enabled
	if state.Quota.MaxObjects == nil && state.Quota.MaxSize == nil {
		state.Quota.Enabled = input.Quota.Enabled
	}

	// // get object locking state
	// s3LockReq := &s3.GetObjectLockConfigurationInput{
	// 	Bucket: aws.String(input.Name),
	// }
	// s3LockResp, err := ce.s3.GetObjectLockConfiguration(ctx, s3LockReq)
	// if err != nil {
	// 	return BucketState{BucketArgs: input, Assimilated: assimilated}, err
	// }

	// if s3LockResp.ObjectLockConfiguration != nil && s3LockResp.ObjectLockConfiguration.ObjectLockEnabled == s3types.ObjectLockEnabledEnabled {
	// 	state.ObjectLocking = true
	// } else {
	// 	state.ObjectLocking = false
	// }

	// get object versioning state
	s3VersReq := &s3.GetBucketVersioningInput{
		Bucket: aws.String(input.Name),
	}
	s3VersResp, err := ce.s3.GetBucketVersioning(ctx, s3VersReq)
	if err != nil {
		return BucketState{BucketArgs: input, Assimilated: assimilated}, err
	}

	if s3VersResp.Status == s3types.BucketVersioningStatusEnabled {
		state.Versioning = true
	} else {
		state.Versioning = false
	}
	return state, nil
}

func (*Bucket) Update(ctx p.Context, id string, olds BucketState, news BucketArgs, preview bool) (BucketState, error) {
	// bail out now when we are in preview mode
	if preview {
		return ToBucketState(news, olds), nil
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return olds, err
	}
	ret, err := UpdateBucket(ctx, ce, news, olds.Assimilated)
	if err != nil {
		return olds, err
	}
	ret.Name = olds.BucketArgs.Name
	return ret, nil
}

func (*Bucket) Delete(ctx p.Context, id string, state BucketState) error {
	ce, c, err := initClient(ctx)
	if err != nil {
		return err
	}
	if state.Assimilated && !c.deleteAssimilated {
		ctx.Logf(diag.Info, "DELETE on %s[%s]: Keeping as this object was assimilated!", "Bucket", id)
		return nil
	}

	// s3req := &s3.DeleteBucketInput{
	// 	Bucket: aws.String(data.Id.ValueString()),
	// }
	// _, err := ce.s3.DeleteBucket(ctx, s3req)
	return ce.client.RemoveBucket(ctx, admin.Bucket{Bucket: state.Name, PurgeObject: &state.PurgeOnDelete})
}
