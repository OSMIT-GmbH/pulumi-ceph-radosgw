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
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ceph/go-ceph/rgw/admin"
	"github.com/pulumi/pulumi-go-provider/infer"

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
type SubUser struct{}

type SubUserPermission string

const (
	SubUserPermissionNone            SubUserPermission = "none"
	SubUserPermissionRead            SubUserPermission = "read"
	SubUserPermissionWrite           SubUserPermission = "write"
	SubUserPermissionReadWrite       SubUserPermission = "readWrite"
	SubUserPermissionReadFullControl SubUserPermission = "fullControl"
)

func (SubUserPermission) Values() []infer.EnumValue[SubUserPermission] {
	return []infer.EnumValue[SubUserPermission]{
		{Name: "none", Value: SubUserPermissionNone},
		{Name: "read", Value: SubUserPermissionRead},
		{Name: "write", Value: SubUserPermissionWrite},
		{Name: "readWrite", Value: SubUserPermissionReadWrite},
		{Name: "fullControl", Value: SubUserPermissionReadFullControl},
	}
}

// Each resource has in input struct, defining what arguments it accepts.
type SubUserArgs struct {
	// User id (parent)
	UserID string `pulumi:"userId"`
	// Sub user name
	SubUserName string `pulumi:"subUserName"`

	Permissions SubUserPermission `pulumi:"permissions"`

	// these are always nil in answers, they are only relevant in requests
	GenerateKey *bool   `pulumi:"generateKey,optional"`
	SecretKey   *string `pulumi:"secretKey,optional" provider:"secret"`
	// Secret      *string  `pulumi:"secret,optional" provider:"secret"`
	PurgeKeys *bool    `pulumi:"purgeKeys,optional"`
	KeyType   *KeyType `pulumi:"keyType,optional"`
}

func (m *SubUserArgs) Annotate(a infer.Annotator) {
	a.Describe(&m.UserID, "User-ID of 'parent' user")
	a.Describe(&m.SubUserName, "Name of this subuser")
}

// Each resource has a state, describing the fields that exist on the created resource.
type SubUserState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	SubUserArgs

	FullName string `pulumi:"fullname"`

	Keys []KeyEntry `pulumi:"keys"`
	// this config element was assimilated...
	// Required: true
	Assimilated bool `pulumi:"_assimilated"`
}

func subUserArgsToAPI(input SubUserArgs) (admin.User, admin.SubuserSpec, error) {
	var access admin.SubuserAccess
	switch input.Permissions {
	case SubUserPermissionRead:
		access = admin.SubuserAccessRead
	case SubUserPermissionWrite:
		access = admin.SubuserAccessWrite
	case SubUserPermissionReadWrite:
		access = admin.SubuserAccessReadWrite
	case SubUserPermissionReadFullControl:
		access = admin.SubuserAccessFull
	case SubUserPermissionNone:
		access = admin.SubuserAccessNone
	default:
		access = admin.SubuserAccessNone
	}
	var keyType *string
	if input.KeyType != nil {
		var kt string
		switch *input.KeyType {
		case KeyTypeS3:
			kt = "s3"
		case KeyTypeSwift:
			kt = "swift"
		default:
			kt = "s3"
		}
		keyType = &kt
	} else {
		keyType = nil
	}
	return admin.User{ID: input.UserID}, admin.SubuserSpec{
		Name:        input.SubUserName,
		Access:      access,
		GenerateKey: input.GenerateKey,
		SecretKey:   input.SecretKey,
		// Secret:      input.Secret,
		PurgeKeys: input.PurgeKeys,
		KeyType:   keyType,
	}, nil
}

func APItoSubUserArgs(input admin.SubuserSpec, orig SubUserArgs) (SubUserArgs, error) {
	var perm SubUserPermission
	switch input.Access {
	case admin.SubuserAccessReplyNone:
		perm = SubUserPermissionNone
	case admin.SubuserAccessReplyRead:
		perm = SubUserPermissionRead
	case admin.SubuserAccessReplyWrite:
		perm = SubUserPermissionWrite
	case admin.SubuserAccessReplyReadWrite:
		perm = SubUserPermissionReadWrite
	case admin.SubuserAccessReplyFull:
		perm = SubUserPermissionReadFullControl
	}
	return SubUserArgs{
		SubUserName: input.Name,
		Permissions: perm,
		UserID:      orig.UserID,
		GenerateKey: orig.GenerateKey,
		SecretKey:   orig.SecretKey,
		// Secret:      orig.Secret,
		PurgeKeys: orig.PurgeKeys,
		KeyType:   orig.KeyType,
	}, nil
}

// All resources must implement Create at a minumum.
func (SubUser) Create(ctx context.Context, req infer.CreateRequest[SubUserArgs]) (infer.CreateResponse[SubUserState], error) {
	// bail out now when we are in preview mode
	if req.DryRun {
		return infer.CreateResponse[SubUserState]{
			ID: IdPreviewPrefix + req.Name,
			Output: SubUserState{
				SubUserArgs: req.Inputs,
			},
		}, nil
	}

	retErr := func(err error) (infer.CreateResponse[SubUserState], error) {
		return infer.CreateResponse[SubUserState]{Output: SubUserState{SubUserArgs: req.Inputs}}, err
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}

	user, subUser, err := subUserArgsToAPI(req.Inputs)
	if err != nil {
		return retErr(err)
	}

	err = ce.client.CreateSubuser(ctx, user, subUser)
	if err != nil {
		// p.GetLogger(ctx).Errorf(Assimilate failed: List failed with %s", err.Error())
		return retErr(err)
	}

	id := fmt.Sprintf("%s:%s", req.Inputs.UserID, req.Inputs.SubUserName)
	subUserState, err := ReadSubUserState(ctx, ce, id, req.Inputs, false)
	return infer.CreateResponse[SubUserState]{ID: id, Output: subUserState}, err
}

func (SubUser) Diff(ctx context.Context, req infer.DiffRequest[SubUserArgs, SubUserState]) (infer.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if req.Inputs.SubUserName != req.State.SubUserName {
		diff["subUserId"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if req.Inputs.Permissions != req.State.Permissions {
		diff["permissions"] = p.PropertyDiff{Kind: p.Update}
	}
	// if news.Email != olds.Email {
	// 	diff["email"] = p.PropertyDiff{Kind: p.Update}
	// }
	//diffWalk(ctx, diff, "tags", reflect.ValueOf(olds.Tags), reflect.ValueOf(news.Tags))

	if len(diff) > 0 {
		p.GetLogger(ctx).Infof("DIFF on SubUser %s/%s: Found %d diffs: %v", req.Inputs.SubUserName, req.ID, len(diff), diff)
	}
	return infer.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (SubUser) Read(ctx context.Context, req infer.ReadRequest[SubUserArgs, SubUserState]) (infer.ReadResponse[SubUserArgs, SubUserState], error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return infer.ReadResponse[SubUserArgs, SubUserState]{ID: req.ID, Inputs: req.Inputs, State: req.State}, err
	}
	subUserState, err := ReadSubUserState(ctx, ce, req.State.FullName, req.Inputs, req.State.Assimilated)
	return infer.ReadResponse[SubUserArgs, SubUserState]{ID: req.ID, Inputs: req.Inputs, State: subUserState}, err
}

func ReadSubUserState(ctx context.Context, ce *CacheEntry, id string, inputs SubUserArgs, assimilated bool) (SubUserState, error) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 {
		return SubUserState{SubUserArgs: inputs, Assimilated: assimilated}, errors.New(fmt.Sprintf("Error splitting up id `%s` into parts", id))
	}

	user, err := ce.client.GetUser(ctx, admin.User{ID: parts[0]})
	if err != nil {
		return SubUserState{SubUserArgs: inputs, Assimilated: assimilated}, err
	}
	var found bool
	var matchingSubuser admin.SubuserSpec
	for _, subuser := range user.Subusers {
		if subuser.Name == id {
			found = true
			matchingSubuser = subuser
			matchingSubuser.Name = strings.TrimPrefix(subuser.Name, user.ID+":")
			break
		}
	}
	if !found {
		return SubUserState{SubUserArgs: inputs, Assimilated: assimilated}, errors.New(fmt.Sprintf("Could not find subuser `%s` within user `%s`", parts[1], parts[0]))
	}

	subUser, err := APItoSubUserArgs(matchingSubuser, inputs)
	fullName := fmt.Sprintf("%s:%s", subUser.UserID, subUser.SubUserName)
	subUserState := SubUserState{SubUserArgs: subUser, FullName: fullName, Assimilated: assimilated}

	// also add subuser keys to response
	for _, key := range user.Keys {
		// p.GetLogger(ctx).Infof("found s3 key for %s (%s) [%s -> %s]", key.User, id, key.AccessKey, key.SecretKey)
		if key.User == id {
			subUserState.Keys = append(subUserState.Keys, KeyEntry{
				AccessKey: key.AccessKey,
				SecretKey: key.SecretKey,
				KeyType:   string(KeyTypeS3),
			})
		}
	}

	for _, key := range user.SwiftKeys {
		if key.User == id {
			subUserState.Keys = append(subUserState.Keys, KeyEntry{
				AccessKey: "",
				SecretKey: key.SecretKey,
				KeyType:   string(KeyTypeSwift),
			})
		}
	}
	return subUserState, err
}

func (SubUser) Update(ctx context.Context, req infer.UpdateRequest[SubUserArgs, SubUserState]) (infer.UpdateResponse[SubUserState], error) {
	user, subUser, err := subUserArgsToAPI(req.Inputs)
	if err != nil {
		return infer.UpdateResponse[SubUserState]{Output: req.State}, err
	}
	// bail out now when we are in preview mode
	if req.DryRun {
		fullName := fmt.Sprintf("%s:%s", req.Inputs.UserID, req.Inputs.SubUserName)
		return infer.UpdateResponse[SubUserState]{Output: SubUserState{SubUserArgs: req.Inputs, FullName: fullName, Assimilated: req.State.Assimilated}}, nil
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return infer.UpdateResponse[SubUserState]{Output: req.State}, err
	}

	err = ce.client.ModifySubuser(ctx, user, subUser)
	if err != nil {
		return infer.UpdateResponse[SubUserState]{Output: req.State}, err
	}

	subUserState, err := ReadSubUserState(ctx, ce, req.State.FullName, req.Inputs, req.State.Assimilated)
	return infer.UpdateResponse[SubUserState]{Output: subUserState}, err
}

func (SubUser) Delete(ctx context.Context, req infer.DeleteRequest[SubUserState]) (infer.DeleteResponse, error) {
	ce, c, err := initClient(ctx)
	if err != nil {
		return infer.DeleteResponse{}, err
	}
	if req.State.Assimilated && !c.deleteAssimilated {
		p.GetLogger(ctx).Infof("DELETE on SubUser[%s]: Keeping as this object was assimilated!", req.ID)
		return infer.DeleteResponse{}, nil
	}
	return infer.DeleteResponse{}, ce.client.RemoveSubuser(ctx, admin.User{ID: req.State.UserID}, admin.SubuserSpec{Name: req.State.SubUserName, PurgeKeys: req.State.PurgeKeys})
}
