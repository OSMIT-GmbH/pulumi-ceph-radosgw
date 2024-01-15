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
	"github.com/aws/aws-sdk-go/aws"
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
type BucketPolicy struct{}

// Each resource has in input struct, defining what arguments it accepts.
type BucketPolicyArgs struct {
	Bucket string `pulumi:"bucket"`
	Policy string `pulumi:"policy"`
}

// Annotate the nested fields for the Metadata receiver
func (m *BucketPolicyArgs) Annotate(a infer.Annotator) {
	a.Describe(&m.Bucket, "Bucket name")
	a.Describe(&m.Policy, "Bucket policy")
}

// Each resource has a state, describing the fields that exist on the created resource.
type BucketPolicyState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	BucketPolicyArgs
}

// All resources must implement Create at a minumum.
func (thiz *BucketPolicy) Create(ctx p.Context, name string, input BucketPolicyArgs, preview bool) (string, BucketPolicyState, error) {
	// bail out now when we are in preview mode
	if preview {
		return IdPreviewPrefix + name, BucketPolicyState{
			BucketPolicyArgs: input,
		}, nil
	}

	retErr := func(err error) (string, BucketPolicyState, error) {
		return "", BucketPolicyState{BucketPolicyArgs: input}, err
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}

	// Put PutBucketPolicyInput
	s3req := &s3.PutBucketPolicyInput{
		Bucket: aws.String(input.Bucket),
		Policy: aws.String(input.Policy),
	}

	_, err = ce.s3.PutBucketPolicy(ctx, s3req)
	if err != nil {
		return retErr(err)
	}

	state, err := UpdateBucketPolicy(ctx, ce, input)

	return name, state, err
}

func (*BucketPolicy) Diff(ctx p.Context, id string, olds BucketPolicyState, news BucketPolicyArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if news.Bucket != olds.Bucket {
		diff["bucket"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if news.Policy != olds.Policy {
		diff["policy"] = p.PropertyDiff{Kind: p.Update}
	}

	if len(diff) > 0 {
		ctx.Log(diag.Info, fmt.Sprintf("DIFF on BucketPolicy %s/%s: Found %d diffs: %v", news.Bucket, id, len(diff), diff))
	}
	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*BucketPolicy) Read(ctx p.Context, id string, inputs BucketPolicyArgs, state BucketPolicyState) (string, BucketPolicyArgs, BucketPolicyState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return id, inputs, state, err
	}
	ret, err := ReadBucketPolicyState(ctx, ce, inputs)
	return id, inputs, ret, err
}

func (*BucketPolicy) Update(ctx p.Context, id string, olds BucketPolicyState, news BucketPolicyArgs, preview bool) (BucketPolicyState, error) {
	// bail out now when we are in preview mode
	if preview {
		return BucketPolicyState{BucketPolicyArgs: news}, nil
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return olds, err
	}

	return UpdateBucketPolicy(ctx, ce, news)
}

func UpdateBucketPolicy(ctx p.Context, ce *CacheEntry, input BucketPolicyArgs) (BucketPolicyState, error) {
	retErr := func(err error) (BucketPolicyState, error) {
		return BucketPolicyState{BucketPolicyArgs: input}, err
	}

	// Put PutBucketPolicyInput
	s3req := &s3.PutBucketPolicyInput{
		Bucket: aws.String(input.Bucket),
		Policy: aws.String(input.Policy),
	}

	_, err := ce.s3.PutBucketPolicy(ctx, s3req)
	if err != nil {
		return retErr(err)
	}

	return ReadBucketPolicyState(ctx, ce, input)
}

func ReadBucketPolicyState(ctx p.Context, ce *CacheEntry, input BucketPolicyArgs) (BucketPolicyState, error) {

	s3req := &s3.GetBucketPolicyInput{
		Bucket: aws.String(input.Bucket),
	}

	s3res, err := ce.s3.GetBucketPolicy(ctx, s3req)
	if err != nil {
		return BucketPolicyState{BucketPolicyArgs: input}, err
	}

	state := BucketPolicyState{BucketPolicyArgs: input}
	state.Policy = *s3res.Policy
	return state, nil
}

func (*BucketPolicy) Delete(ctx p.Context, id string, state BucketPolicyState) error {
	ce, _, err := initClient(ctx)
	if err != nil {
		return err
	}
	s3req := &s3.DeleteBucketPolicyInput{
		Bucket: aws.String(state.Bucket),
	}
	_, err = ce.s3.DeleteBucketPolicy(ctx, s3req)
	return err
}
