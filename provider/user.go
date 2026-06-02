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
	"reflect"
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
	Users        *CapabilityPermission `pulumi:"users,optional" capname:"users"`
	Buckets      *CapabilityPermission `pulumi:"buckets,optional" capname:"buckets"`
	Metadata     *CapabilityPermission `pulumi:"metadata,optional" capname:"metadata"`
	Usage        *CapabilityPermission `pulumi:"usage,optional" capname:"usage"`
	Zone         *CapabilityPermission `pulumi:"zone,optional" capname:"zone"`
	AmzCache     *CapabilityPermission `pulumi:"amzCache,optional" capname:"amz-cache"`
	Info         *CapabilityPermission `pulumi:"info,optional" capname:"info"`
	Bilog        *CapabilityPermission `pulumi:"bilog,optional" capname:"bilog"`
	Mdlog        *CapabilityPermission `pulumi:"mdlog,optional" capname:"mdlog"`
	Datalog      *CapabilityPermission `pulumi:"datalog,optional" capname:"datalog"`
	UserPolicy   *CapabilityPermission `pulumi:"userPolicy,optional" capname:"user-policy"`
	OidcProvider *CapabilityPermission `pulumi:"oidcProvider,optional" capname:"oidc-provider"`
	Roles        *CapabilityPermission `pulumi:"roles,optional" capname:"roles"`
	Ratelimit    *CapabilityPermission `pulumi:"ratelimit,optional" capname:"ratelimit"`
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
	Capabilities *Capabilities `pulumi:"capabilities,optional"`
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
	capType := reflect.TypeOf(Capabilities{})
	for b := 0; b < capType.NumField(); b++ {
		tag := capType.Field(b).Tag.Get("capname")
		tagList := strings.Split(tag, ",")
		capMap[tagList[0]] = b
	}
	var sb strings.Builder
	if input.Capabilities != nil {
		capVal := reflect.ValueOf(*input.Capabilities)
		for b := 0; b < capType.NumField(); b++ {
			v2 := capVal.Field(b)
			name := capType.Field(b).Tag.Get("capname")
			tagList := strings.Split(name, ",")
			name = tagList[0]
			if v2.IsNil() || v2.Elem().String() == "" {
				continue
			}
			if sb.Len() > 0 {
				sb.WriteString("; ")
			}
			sb.WriteString(name)
			sb.WriteString("=")
			sb.WriteString(v2.Elem().String())
		}
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

func APItoUserArgs(ctx context.Context, capMap map[string]int, resp admin.User, assimilated bool) UserState {
	user := UserArgs{
		UserID:      resp.ID,
		DisplayName: resp.DisplayName,
		Email:       resp.Email,
		MaxBuckets:  resp.MaxBuckets,
		Suspended:   *resp.Suspended == 1,
	}
	// p.GetLogger(ctx).Infof("User Caps for %s: %s\n", resp.ID, resp.Caps)

	if user.Capabilities == nil {
		user.Capabilities = &Capabilities{}
	}
	ncv := reflect.ValueOf(user.Capabilities).Elem()
	// also add subuser keys to response
	for _, cap := range resp.Caps {
		// p.GetLogger(ctx).Infof("cap  %s=%s\n", cap.Type, cap.Perm)
		idx, ok := capMap[cap.Type]
		if !ok {
			p.GetLogger(ctx).Errorf("ERROR: Unknown cap while looking up %s => %s on user %s\n", cap.Type, cap.Perm, user.UserID)
			continue
		}
		f := ncv.Field(idx)
		if f.IsNil() {
			cp := CapabilityPermission(cap.Perm)
			f.Set(reflect.ValueOf(&cp))
		} else {
			f.Elem().SetString(cap.Perm)
		}
	}

	userState := UserState{UserArgs: user, Assimilated: assimilated}
	// also add subuser keys to response
	for _, key := range resp.Keys {
		// p.GetLogger(ctx).Infof("found s3 key for %s (%s) [%s -> %s]\n", key.User, id, key.AccessKey, key.SecretKey)
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
func (User) Create(ctx context.Context, req infer.CreateRequest[UserArgs]) (infer.CreateResponse[UserState], error) {
	// bail out now when we are in preview mode
	if req.DryRun {
		return infer.CreateResponse[UserState]{
			ID: IdPreviewPrefix + req.Name,
			Output: UserState{
				UserArgs: req.Inputs,
			},
		}, nil
	}

	retErr := func(err error) (infer.CreateResponse[UserState], error) {
		return infer.CreateResponse[UserState]{Output: UserState{UserArgs: req.Inputs}}, err
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}

	user, capMap := userArgsToAPI(req.Inputs)

	user, err = ce.client.CreateUser(ctx, user)
	if err != nil {
		// p.GetLogger(ctx).Errorf("Assimilate failed: List failed with %s\n", err.Error())
		return retErr(err)
	}

	return infer.CreateResponse[UserState]{
		ID:     user.ID,
		Output: APItoUserArgs(ctx, capMap, user, false),
	}, err
}

func (User) Diff(ctx context.Context, req infer.DiffRequest[UserArgs, UserState]) (infer.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if req.Inputs.UserID != req.State.UserID {
		diff["userId"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if req.Inputs.DisplayName != req.State.DisplayName {
		diff["displayName"] = p.PropertyDiff{Kind: p.Update}
	}
	if req.Inputs.Email != req.State.Email {
		diff["email"] = p.PropertyDiff{Kind: p.Update}
	}
	diffWalk(ctx, diff, "capabilities", reflect.ValueOf(req.State.Capabilities), reflect.ValueOf(req.Inputs.Capabilities))

	if len(diff) > 0 {
		p.GetLogger(ctx).Infof("DIFF on User %s/%s: Found %d diffs: %v\n", req.Inputs.UserID, req.ID, len(diff), diff)
	}
	return infer.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (User) Read(ctx context.Context, req infer.ReadRequest[UserArgs, UserState]) (infer.ReadResponse[UserArgs, UserState], error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return infer.ReadResponse[UserArgs, UserState]{
			ID:     req.ID,
			Inputs: req.Inputs,
			State:  req.State,
		}, err
	}
	resp, err := ce.client.GetUser(ctx, admin.User{ID: req.Inputs.UserID})
	if err != nil {
		return infer.ReadResponse[UserArgs, UserState]{
			ID:     req.ID,
			Inputs: req.Inputs,
			State:  req.State,
		}, err
	}
	_, capMap := userArgsToAPI(req.Inputs)
	return infer.ReadResponse[UserArgs, UserState]{
		ID:     req.ID,
		Inputs: req.Inputs,
		State:  APItoUserArgs(ctx, capMap, resp, req.State.Assimilated),
	}, err
}

func (User) Update(ctx context.Context, req infer.UpdateRequest[UserArgs, UserState]) (infer.UpdateResponse[UserState], error) {
	// bail out now when we are in preview mode
	if req.DryRun {
		return infer.UpdateResponse[UserState]{
			Output: UserState{UserArgs: req.Inputs, Assimilated: req.State.Assimilated},
		}, nil
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return infer.UpdateResponse[UserState]{Output: req.State}, err
	}

	user, capMap := userArgsToAPI(req.Inputs)
	user, err = ce.client.ModifyUser(ctx, user)
	if err != nil {
		return infer.UpdateResponse[UserState]{Output: req.State}, err
	}
	ret := APItoUserArgs(ctx, capMap, user, req.State.Assimilated)

	capVal := reflect.ValueOf(req.Inputs.Capabilities)
	capValOld := reflect.ValueOf(req.State.Capabilities)
	getCapStr := func(v reflect.Value, b int) string {
		if v.IsNil() {
			return ""
		}
		f := v.Elem().Field(b)
		if f.IsNil() {
			return ""
		}
		return f.Elem().String()
	}
	var sbAdd strings.Builder
	var sbRemove strings.Builder
	for name, b := range capMap {
		newPerm := getCapStr(capVal, b)
		oldPerm := getCapStr(capValOld, b)
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
			// p.GetLogger(ctx).Infof("cap  %s=%s\n", cap.Type, cap.Perm)
			idx, ok := capMap[cap.Type]
			if !ok {
				p.GetLogger(ctx).Errorf("ERROR: Unknown cap while looking up %s => %s on user %s\n", cap.Type, cap.Perm, user.ID)
				continue
			}
			f := ncv.Field(idx)
			if f.IsNil() {
				cp := CapabilityPermission(cap.Perm)
				f.Set(reflect.ValueOf(&cp))
			} else {
				f.Elem().SetString(cap.Perm)
			}
		}
		ret.Capabilities = &nc
	}
	if sbRemove.Len() > 0 {
		caps, err := ce.client.RemoveUserCap(ctx, user.ID, sbRemove.String())
		if err != nil {
			return infer.UpdateResponse[UserState]{Output: req.State}, err
		}
		// p.GetLogger(ctx).Infof("removeCaps:  %s => %s\n", sbRemove.String(), fmt.Sprint(caps))
		updateCaps(caps)
	}
	if sbAdd.Len() > 0 {
		caps, err := ce.client.AddUserCap(ctx, user.ID, sbAdd.String())
		if err != nil {
			return infer.UpdateResponse[UserState]{Output: req.State}, err
		}
		// p.GetLogger(ctx).Infof("addCaps:  %s => %s\n", sbAdd.String(), fmt.Sprint(caps))
		updateCaps(caps)
	}

	return infer.UpdateResponse[UserState]{Output: ret}, err
}

func (User) Delete(ctx context.Context, req infer.DeleteRequest[UserState]) (infer.DeleteResponse, error) {
	ce, c, err := initClient(ctx)
	if err != nil {
		return infer.DeleteResponse{}, err
	}
	if req.State.Assimilated && !c.deleteAssimilated {
		p.GetLogger(ctx).Infof("DELETE on User[%s]: Keeping as this object was assimilated!\n", req.ID)
		return infer.DeleteResponse{}, nil
	}
	return infer.DeleteResponse{}, ce.client.RemoveUser(ctx, admin.User{ID: req.ID})
}
