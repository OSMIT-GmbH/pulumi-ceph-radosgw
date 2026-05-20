// Copyright 2025, Pulumi Corporation.
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

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

// Name controls how this provider is referenced in package names and elsewhere.
const Name string = "ceph-radosgw"

// Provider creates a new instance of the provider.
func Provider() p.Provider {
	p, err := infer.NewProviderBuilder().
		WithDisplayName("Ceph Rados Gateway").
		WithDescription("A Pulumi provider for managing Ceph Rados Gateway resources.").
		WithHomepage("https://github.com/OSMIT-GmbH/pulumi-ceph-radosgw").
		WithNamespace("osmit-gmbh").
		WithGoImportPath("github.com/OSMIT-GmbH/pulumi-ceph-radosgw/sdk/go/pulumi-ceph-radosgw").
		WithRepository("https://github.com/OSMIT-GmbH/pulumi-ceph-radosgw").
		WithResources(
			infer.Resource[*Bucket, BucketArgs, BucketState](&Bucket{}),
			infer.Resource[*BucketPolicy, BucketPolicyArgs, BucketPolicyState](&BucketPolicy{}),
			infer.Resource[*SubUser, SubUserArgs, SubUserState](&SubUser{}),
			infer.Resource[*User, UserArgs, UserState](&User{}),
			infer.Resource[*Key, KeyArgs, KeyState](&Key{}),
		).
		// WithComponents(infer.ComponentF(NewRandomComponent)).
		WithConfig(infer.Config(&ProviderConfig{})).
		WithModuleMap(map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		}).Build()
	if err != nil {
		panic(fmt.Errorf("unable to build provider: %w", err))
	}
	return p
}

