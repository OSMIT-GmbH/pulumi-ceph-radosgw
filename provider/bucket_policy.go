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
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
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
func (BucketPolicy) Create(ctx context.Context, req infer.CreateRequest[BucketPolicyArgs]) (infer.CreateResponse[BucketPolicyState], error) {
	// bail out now when we are in preview mode
	if req.DryRun {
		return infer.CreateResponse[BucketPolicyState]{
			ID: IdPreviewPrefix + req.Name,
			Output: BucketPolicyState{
				BucketPolicyArgs: req.Inputs,
			},
		}, nil
	}

	retErr := func(err error) (infer.CreateResponse[BucketPolicyState], error) {
		return infer.CreateResponse[BucketPolicyState]{Output: BucketPolicyState{BucketPolicyArgs: req.Inputs}}, err
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return retErr(err)
	}

	// Put PutBucketPolicyInput
	s3req := &s3.PutBucketPolicyInput{
		Bucket: aws.String(req.Inputs.Bucket),
		Policy: aws.String(req.Inputs.Policy),
	}

	_, err = ce.s3.PutBucketPolicy(ctx, s3req)
	if err != nil {
		return retErr(err)
	}

	state, err := UpdateBucketPolicy(ctx, ce, req.Inputs)

	return infer.CreateResponse[BucketPolicyState]{
		ID:     req.Name,
		Output: state,
	}, err
}

func (BucketPolicy) Diff(ctx context.Context, req infer.DiffRequest[BucketPolicyArgs, BucketPolicyState]) (infer.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}
	if req.Inputs.Bucket != req.State.Bucket {
		diff["bucket"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}
	if req.Inputs.Policy != req.State.Policy {
		diff["policy"] = p.PropertyDiff{Kind: p.Update}
	}

	if len(diff) > 0 {
		fmt.Printf("DIFF on BucketPolicy %s/%s: Found %d diffs: %v", req.Inputs.Bucket, req.ID, len(diff), diff)
	}
	return infer.DiffResponse{
		DeleteBeforeReplace: hasReplaceDiff(diff),
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (BucketPolicy) Read(ctx context.Context, req infer.ReadRequest[BucketPolicyArgs, BucketPolicyState]) (infer.ReadResponse[BucketPolicyArgs, BucketPolicyState], error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return infer.ReadResponse[BucketPolicyArgs, BucketPolicyState]{
			ID:     req.ID,
			Inputs: req.Inputs,
			State:  req.State,
		}, err
	}
	ret, err := ReadBucketPolicyState(ctx, ce, req.Inputs)
	return infer.ReadResponse[BucketPolicyArgs, BucketPolicyState]{
		ID:     req.ID,
		Inputs: req.Inputs,
		State:  ret,
	}, err
}

func (BucketPolicy) Update(ctx context.Context, req infer.UpdateRequest[BucketPolicyArgs, BucketPolicyState]) (infer.UpdateResponse[BucketPolicyState], error) {
	// bail out now when we are in preview mode
	if req.DryRun {
		return infer.UpdateResponse[BucketPolicyState]{Output: BucketPolicyState{BucketPolicyArgs: req.Inputs}}, nil
	}
	ce, _, err := initClient(ctx)
	if err != nil {
		return infer.UpdateResponse[BucketPolicyState]{Output: req.State}, err
	}

	result, err := UpdateBucketPolicy(ctx, ce, req.Inputs)
	return infer.UpdateResponse[BucketPolicyState]{Output: result}, err
}

func UpdateBucketPolicy(ctx context.Context, ce *CacheEntry, input BucketPolicyArgs) (BucketPolicyState, error) {
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

func ReadBucketPolicyState(ctx context.Context, ce *CacheEntry, input BucketPolicyArgs) (BucketPolicyState, error) {

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

func (BucketPolicy) Delete(ctx context.Context, req infer.DeleteRequest[BucketPolicyState]) (infer.DeleteResponse, error) {
	ce, _, err := initClient(ctx)
	if err != nil {
		return infer.DeleteResponse{}, err
	}
	s3req := &s3.DeleteBucketPolicyInput{
		Bucket: aws.String(req.State.Bucket),
	}
	_, err = ce.s3.DeleteBucketPolicy(ctx, s3req)
	return infer.DeleteResponse{}, err
}
