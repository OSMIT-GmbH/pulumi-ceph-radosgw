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
	"reflect"
	"strings"

	"github.com/ceph/go-ceph/rgw/admin"
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
type User struct{}

// https://docs.ceph.com/en/latest/radosgw/admin/#add-remove-admin-capabilities
// --caps="[users|buckets|metadata|usage|zone|amz-cache|info|bilog|mdlog|datalog|user-policy|oidc-provider|roles|ratelimit]=[*|read|write|read, write]"

/*
type Capability string

const (

	CapabilityUsers        Capability = "users"
	CapabilityBuckets      Capability = "buckets"
	CapabilityMetadata     Capability = "metadata"
	CapabilityUsage        Capability = "usage"
	CapabilityZone         Capability = "zone"
	CapabilityAmzCache     Capability = "amz-cache"
	CapabilityInfo         Capability = "info"
	CapabilityBilog        Capability = "bilog"
	CapabilityMdlog        Capability = "mdlog"
	CapabilityDatalog      Capability = "datalog"
	CapabilityUserPolicy   Capability = "user-policy"
	CapabilityOidcProvider Capability = "oidc-provider"
	CapabilityRoles        Capability = "roles"
	CapabilityRatelimit    Capability = "ratelimit"

)

	func (Capability) Values() []infer.EnumValue[Capability] {
		return []infer.EnumValue[Capability]{
			{Name: "users", Value: CapabilityUsers},
			{Name: "buckets", Value: CapabilityBuckets},
			{Name: "metadata", Value: CapabilityMetadata},
			{Name: "usage", Value: CapabilityUsage},
			{Name: "zone", Value: CapabilityZone},
			{Name: "amz-cache", Value: CapabilityAmzCache},
			{Name: "info", Value: CapabilityInfo},
			{Name: "bilog", Value: CapabilityBilog},
			{Name: "mdlog", Value: CapabilityMdlog},
			{Name: "datalog", Value: CapabilityDatalog},
			{Name: "user-policy", Value: CapabilityUserPolicy},
			{Name: "oidc-provider", Value: CapabilityOidcProvider},
			{Name: "roles", Value: CapabilityRoles},
			{Name: "ratelimit", Value: CapabilityRatelimit},
		}
	}
*/
type CapabilityPermission string

const (
	CapabilityPermissionRead  CapabilityPermission = "read"
	CapabilityPermissionWrite CapabilityPermission = "write"
	// CapabilityPermissionReadWrite CapabilityPermission = "read,write" seems to map to "Full"
	CapabilityPermissionFull CapabilityPermission = "*"
)

func (CapabilityPermission) Values() []infer.EnumValue[CapabilityPermission] {
	return []infer.EnumValue[CapabilityPermission]{
		{Name: "read", Value: CapabilityPermissionRead},
		{Name: "write", Value: CapabilityPermissionWrite},
		// {Name: "read,write", Value: CapabilityPermissionReadWrite},
		{Name: "*", Value: CapabilityPermissionFull},
	}
}

type Capabilities struct {
	Users        CapabilityPermission `pulumi:"users,optional" capname:"users"`
	Buckets      CapabilityPermission `pulumi:"buckets,optional" capname:"buckets"`
	Metadata     CapabilityPermission `pulumi:"metadata,optional" capname:"metadata"`
	Usage        CapabilityPermission `pulumi:"usage,optional" capname:"usage"`
	Zone         CapabilityPermission `pulumi:"zone,optional" capname:"zone"`
	AmzCache     CapabilityPermission `pulumi:"amzCache,optional" capname:"amz-cache"`
	Info         CapabilityPermission `pulumi:"info,optional" capname:"info"`
	Bilog        CapabilityPermission `pulumi:"bilog,optional" capname:"bilog"`
	Mdlog        CapabilityPermission `pulumi:"mdlog,optional" capname:"mdlog"`
	Datalog      CapabilityPermission `pulumi:"datalog,optional" capname:"datalog"`
	UserPolicy   CapabilityPermission `pulumi:"userPolicy,optional" capname:"user-policy"`
	OidcProvider CapabilityPermission `pulumi:"oidcProvider,optional" capname:"oidc-provider"`
	Roles        CapabilityPermission `pulumi:"roles,optional" capname:"roles"`
	Ratelimit    CapabilityPermission `pulumi:"ratelimit,optional" capname:"ratelimit"`
}

// Each resource has in input struct, defining what arguments it accepts.
type UserArgs struct {
	// user id
	UserID string `pulumi:"userId"`
	// display name
	DisplayName string `pulumi:"displayName,optional"`
	// email Address
	Email        string       `pulumi:"email,optional"`
	Suspended    bool         `pulumi:"suspended,optional"`
	MaxBuckets   *int         `pulumi:"maxBuckets,optional"`
	Capabilities Capabilities `pulumi:"capabilities,optional"`
}

type KeyType string

const (
	KeyTypeS3    KeyType = "s3"
	KeyTypeSwift KeyType = "swift"
)

func (KeyType) Values() []infer.EnumValue[KeyType] {
	return []infer.EnumValue[KeyType]{
		{Name: "s3", Value: KeyTypeS3},
		{Name: "swift", Value: KeyTypeSwift},
	}
}

// keys for user/subUser
type KeyEntry struct {
	SecretKey string `pulumi:"secretKey" provider:"secret"`
	AccessKey string `pulumi:"accessKey,optional"`
	KeyType   string `pulumi:"keyType"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type UserState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	UserArgs

	Keys []KeyEntry `pulumi:"keys"`

	// this config element was assimilated...
	// Required: true
	Assimilated bool `pulumi:"_assimilated"`
}

func userArgsToAPI(input UserArgs) (admin.User, map[string]int) {
	capMap := make(map[string]int)
	capType := reflect.TypeOf(input.Capabilities)
	capVal := reflect.ValueOf(input.Capabilities)
	var sb strings.Builder
	for b := 0; b < capType.NumField(); b++ {
		v2 := capVal.Field(b)
		tag := capType.Field(b).Tag.Get("capname")
		tagList := strings.Split(tag, ",")
		name := tagList[0]
		capMap[name] = b
		if v2.String() == "" {
			continue
		}
		if sb.Len() > 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(name)
		sb.WriteString("=")
		sb.WriteString(v2.String())
	}
	// func getReflect(i interface{}, acceptableFields []string, values *url.Values) {
	// 	t := reflect.TypeOf(i)
	// 	v := reflect.ValueOf(i)

	// 	for b := 0; b < v.NumField(); b++ {
	// 		v2 := v.Field(b)
	// 		tag := t.Field(b).Tag.Get("url")
	// 		if tag == "-" {
	// 			continue
	// 		}
	// 		tagList := strings.Split(tag, ",")
	// 		name := tagList[0]
	// 		if len(name) == 0 {
	// 			name = t.Field(b).Name
	// 		}

	// 		if v2.Kind() == reflect.Struct {
	// 			getReflect(v2.Interface(), acceptableFields, values)
	// 			continue
	// 		}

	// 		if v2.Kind() == reflect.Slice {
	// 			for i := 0; i < v2.Len(); i++ {
	// 				item := v2.Index(i)
	// 				getReflect(item.Interface(), acceptableFields, values)
	// 			}
	// 			continue
	// 		}

	// 		if v2.Kind() == reflect.String ||
	// 			v2.Kind() == reflect.Bool ||
	// 			v2.Kind() == reflect.Int {

	// 			_v2 := fmt.Sprint(v2)
	// 			if len(_v2) > 0 && contains(acceptableFields, name) {
	// 				values.Add(name, _v2)
	// 			}
	// 			continue
	// 		}

	// 		if v2.Kind() == reflect.Ptr && v2.IsValid() && !v2.IsNil() {
	// 			_v2 := fmt.Sprint(v2.Elem())
	// 			if len(_v2) > 0 && contains(acceptableFields, name) {
	// 				values.Add(name, _v2)
	// 			}
	// 			continue
	// 		}
	// 	}
	// }
	return admin.User{
		ID:          input.UserID,
		DisplayName: input.DisplayName,
		Email:       input.Email,
		Suspended:   ifted[int](input.Suspended, 1, 0),
		MaxBuckets:  input.MaxBuckets,
		UserCaps:    sb.String(),
	}, capMap
}

func APItoUserArgs(ctx p.Context, capMap map[string]int, resp admin.User, assimilated bool) UserState {
	user := UserArgs{
		UserID:      resp.ID,
		DisplayName: resp.DisplayName,
		Email:       resp.Email,
		MaxBuckets:  resp.MaxBuckets,
		Suspended:   *resp.Suspended == 1,
	}
	// ctx.Logf(diag.Info, "User Caps for %s: %s", resp.ID, resp.Caps)

	ncv := reflect.ValueOf(&user.Capabilities).Elem()
	// also add subuser keys to response
	for _, cap := range resp.Caps {
		// ctx.Logf(diag.Info, "cap  %s=%s", cap.Type, cap.Perm)
		idx, ok := capMap[cap.Type]
		if !ok {
			ctx.Logf(diag.Error, "Unknown cap while looking up %s => %s on user %s", cap.Type, cap.Perm, user.UserID)
			continue
		}
		ncv.Field(idx).SetString(cap.Perm)
	}

	userState := UserState{UserArgs: user, Assimilated: assimilated}
	// also add subuser keys to response
	for _, key := range resp.Keys {
		// ctx.Logf(diag.Info, "found s3 key for %s (%s) [%s -> %s]", key.User, id, key.AccessKey, key.SecretKey)
		if key.User == resp.ID {
			userState.Keys = append(userState.Keys, KeyEntry{
				AccessKey: key.AccessKey,
				SecretKey: key.SecretKey,
				KeyType:   string(KeyTypeS3),
			})
		}
	}

	for _, key := range resp.SwiftKeys {
		if key.User == resp.ID {
			userState.Keys = append(userState.Keys, KeyEntry{
				AccessKey: "",
				SecretKey: key.SecretKey,
				KeyType:   string(KeyTypeSwift),
			})
		}
	}
	return userState
}

// All resources must implement Create at a minumum.
func (thiz *User) Create(ctx p.Context, name string, input UserArgs, preview bool) (string, UserState, error) {
	// bail out now when we are in preview mode
	if preview {
		return IdPreviewPrefix + name, UserState{
			UserArgs: input,
		}, nil
	}

	retErr := func(err error) (string, UserState, error) {
		return "", UserState{UserArgs: input}, err
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}

	user, capMap := userArgsToAPI(input)

	user, err = ce.client.CreateUser(ctx, user)
	if err != nil {
		// ctx.Logf(diag.Error, "Assimilate failed: List failed with %s", err.Error())
		return retErr(err)
	}

	return user.ID, APItoUserArgs(ctx, capMap, user, false), err
}

func (*User) Diff(ctx p.Context, id string, olds UserState, news UserArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if news.UserID != olds.UserID {
		diff["userId"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if news.DisplayName != olds.DisplayName {
		diff["displayName"] = p.PropertyDiff{Kind: p.Update}
	}
	if news.Email != olds.Email {
		diff["email"] = p.PropertyDiff{Kind: p.Update}
	}
	diffWalk(ctx, diff, "capabilities", reflect.ValueOf(olds.Capabilities), reflect.ValueOf(news.Capabilities))

	if len(diff) > 0 {
		ctx.Log(diag.Info, fmt.Sprintf("DIFF on User %s/%s: Found %d diffs: %v", news.UserID, id, len(diff), diff))
	}
	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*User) Read(ctx p.Context, id string, inputs UserArgs, state UserState) (string, UserArgs, UserState, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return id, inputs, state, err
	}
	resp, err := ce.client.GetUser(ctx, admin.User{ID: inputs.UserID})
	if err != nil {
		return id, inputs, state, err
	}
	_, capMap := userArgsToAPI(inputs)
	return id, inputs, APItoUserArgs(ctx, capMap, resp, state.Assimilated), err
}

func (*User) Update(ctx p.Context, id string, olds UserState, news UserArgs, preview bool) (UserState, error) {
	// bail out now when we are in preview mode
	if preview {
		return UserState{UserArgs: news, Assimilated: olds.Assimilated}, nil
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return olds, err
	}

	user, capMap := userArgsToAPI(news)
	user, err = ce.client.ModifyUser(ctx, user)
	if err != nil {
		return olds, err
	}
	ret := APItoUserArgs(ctx, capMap, user, olds.Assimilated)

	capVal := reflect.ValueOf(news.Capabilities)
	capValOld := reflect.ValueOf(olds.Capabilities)
	var sbAdd strings.Builder
	var sbRemove strings.Builder
	for name, b := range capMap {
		newPerm := capVal.Field(b).String()
		oldPerm := capValOld.Field(b).String()
		if newPerm == oldPerm {
			continue
		}
		if oldPerm != "" {
			if sbRemove.Len() > 0 {
				sbRemove.WriteString("; ")
			}
			sbRemove.WriteString(name)
			sbRemove.WriteString("=")
			sbRemove.WriteString(oldPerm)
		}
		if newPerm != "" {
			if sbAdd.Len() > 0 {
				sbAdd.WriteString("; ")
			}
			sbAdd.WriteString(name)
			sbAdd.WriteString("=")
			sbAdd.WriteString(newPerm)
		}
	}
	updateCaps := func(caps []admin.UserCapSpec) {
		nc := Capabilities{}
		ncv := reflect.ValueOf(&nc).Elem()
		for _, cap := range caps {
			// ctx.Logf(diag.Info, "cap  %s=%s", cap.Type, cap.Perm)
			idx, ok := capMap[cap.Type]
			if !ok {
				ctx.Logf(diag.Error, "Unknown cap while looking up %s => %s on user %s", cap.Type, cap.Perm, user.ID)
				continue
			}
			ncv.Field(idx).SetString(cap.Perm)
			// nc.Buckets = "*a"
		}
		ret.Capabilities = nc
	}
	if sbRemove.Len() > 0 {
		caps, err := ce.client.RemoveUserCap(ctx, user.ID, sbRemove.String())
		if err != nil {
			return olds, err
		}
		// ctx.Logf(diag.Info, "removeCaps:  %s => %s", sbRemove.String(), fmt.Sprint(caps))
		updateCaps(caps)
	}
	if sbAdd.Len() > 0 {
		caps, err := ce.client.AddUserCap(ctx, user.ID, sbAdd.String())
		if err != nil {
			return olds, err
		}
		// ctx.Logf(diag.Info, "addCaps:  %s => %s", sbAdd.String(), fmt.Sprint(caps))
		updateCaps(caps)
	}

	return ret, err
}

func (*User) Delete(ctx p.Context, id string, state UserState) error {
	ce, c, err := initClient(ctx)
	if err != nil {
		return err
	}
	if state.Assimilated && !c.deleteAssimilated {
		ctx.Logf(diag.Info, "DELETE on %s[%s]: Keeping as this object was assimilated!", "User", id)
		return nil
	}
	return ce.client.RemoveUser(ctx, admin.User{ID: id})
}
