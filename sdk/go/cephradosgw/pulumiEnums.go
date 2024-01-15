// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cephradosgw

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type CapabilityPermission string

const (
	CapabilityPermissionRead     = CapabilityPermission("read")
	CapabilityPermissionWrite    = CapabilityPermission("write")
	CapabilityPermissionAsterisk = CapabilityPermission("*")
)

func (CapabilityPermission) ElementType() reflect.Type {
	return reflect.TypeOf((*CapabilityPermission)(nil)).Elem()
}

func (e CapabilityPermission) ToCapabilityPermissionOutput() CapabilityPermissionOutput {
	return pulumi.ToOutput(e).(CapabilityPermissionOutput)
}

func (e CapabilityPermission) ToCapabilityPermissionOutputWithContext(ctx context.Context) CapabilityPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CapabilityPermissionOutput)
}

func (e CapabilityPermission) ToCapabilityPermissionPtrOutput() CapabilityPermissionPtrOutput {
	return e.ToCapabilityPermissionPtrOutputWithContext(context.Background())
}

func (e CapabilityPermission) ToCapabilityPermissionPtrOutputWithContext(ctx context.Context) CapabilityPermissionPtrOutput {
	return CapabilityPermission(e).ToCapabilityPermissionOutputWithContext(ctx).ToCapabilityPermissionPtrOutputWithContext(ctx)
}

func (e CapabilityPermission) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CapabilityPermission) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CapabilityPermission) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CapabilityPermission) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CapabilityPermissionOutput struct{ *pulumi.OutputState }

func (CapabilityPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapabilityPermission)(nil)).Elem()
}

func (o CapabilityPermissionOutput) ToCapabilityPermissionOutput() CapabilityPermissionOutput {
	return o
}

func (o CapabilityPermissionOutput) ToCapabilityPermissionOutputWithContext(ctx context.Context) CapabilityPermissionOutput {
	return o
}

func (o CapabilityPermissionOutput) ToCapabilityPermissionPtrOutput() CapabilityPermissionPtrOutput {
	return o.ToCapabilityPermissionPtrOutputWithContext(context.Background())
}

func (o CapabilityPermissionOutput) ToCapabilityPermissionPtrOutputWithContext(ctx context.Context) CapabilityPermissionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CapabilityPermission) *CapabilityPermission {
		return &v
	}).(CapabilityPermissionPtrOutput)
}

func (o CapabilityPermissionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CapabilityPermissionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CapabilityPermission) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CapabilityPermissionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CapabilityPermissionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CapabilityPermission) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CapabilityPermissionPtrOutput struct{ *pulumi.OutputState }

func (CapabilityPermissionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapabilityPermission)(nil)).Elem()
}

func (o CapabilityPermissionPtrOutput) ToCapabilityPermissionPtrOutput() CapabilityPermissionPtrOutput {
	return o
}

func (o CapabilityPermissionPtrOutput) ToCapabilityPermissionPtrOutputWithContext(ctx context.Context) CapabilityPermissionPtrOutput {
	return o
}

func (o CapabilityPermissionPtrOutput) Elem() CapabilityPermissionOutput {
	return o.ApplyT(func(v *CapabilityPermission) CapabilityPermission {
		if v != nil {
			return *v
		}
		var ret CapabilityPermission
		return ret
	}).(CapabilityPermissionOutput)
}

func (o CapabilityPermissionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CapabilityPermissionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CapabilityPermission) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CapabilityPermissionInput is an input type that accepts values of the CapabilityPermission enum
// A concrete instance of `CapabilityPermissionInput` can be one of the following:
//
//	CapabilityPermissionRead
//	CapabilityPermissionWrite
//	CapabilityPermissionAsterisk
type CapabilityPermissionInput interface {
	pulumi.Input

	ToCapabilityPermissionOutput() CapabilityPermissionOutput
	ToCapabilityPermissionOutputWithContext(context.Context) CapabilityPermissionOutput
}

var capabilityPermissionPtrType = reflect.TypeOf((**CapabilityPermission)(nil)).Elem()

type CapabilityPermissionPtrInput interface {
	pulumi.Input

	ToCapabilityPermissionPtrOutput() CapabilityPermissionPtrOutput
	ToCapabilityPermissionPtrOutputWithContext(context.Context) CapabilityPermissionPtrOutput
}

type capabilityPermissionPtr string

func CapabilityPermissionPtr(v string) CapabilityPermissionPtrInput {
	return (*capabilityPermissionPtr)(&v)
}

func (*capabilityPermissionPtr) ElementType() reflect.Type {
	return capabilityPermissionPtrType
}

func (in *capabilityPermissionPtr) ToCapabilityPermissionPtrOutput() CapabilityPermissionPtrOutput {
	return pulumi.ToOutput(in).(CapabilityPermissionPtrOutput)
}

func (in *capabilityPermissionPtr) ToCapabilityPermissionPtrOutputWithContext(ctx context.Context) CapabilityPermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CapabilityPermissionPtrOutput)
}

func (in *capabilityPermissionPtr) ToOutput(ctx context.Context) pulumix.Output[*CapabilityPermission] {
	return pulumix.Output[*CapabilityPermission]{
		OutputState: in.ToCapabilityPermissionPtrOutputWithContext(ctx).OutputState,
	}
}

type KeyType string

const (
	KeyTypeS3    = KeyType("s3")
	KeyTypeSwift = KeyType("swift")
)

func (KeyType) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyType)(nil)).Elem()
}

func (e KeyType) ToKeyTypeOutput() KeyTypeOutput {
	return pulumi.ToOutput(e).(KeyTypeOutput)
}

func (e KeyType) ToKeyTypeOutputWithContext(ctx context.Context) KeyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeyTypeOutput)
}

func (e KeyType) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return e.ToKeyTypePtrOutputWithContext(context.Background())
}

func (e KeyType) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return KeyType(e).ToKeyTypeOutputWithContext(ctx).ToKeyTypePtrOutputWithContext(ctx)
}

func (e KeyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeyTypeOutput struct{ *pulumi.OutputState }

func (KeyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyType)(nil)).Elem()
}

func (o KeyTypeOutput) ToKeyTypeOutput() KeyTypeOutput {
	return o
}

func (o KeyTypeOutput) ToKeyTypeOutputWithContext(ctx context.Context) KeyTypeOutput {
	return o
}

func (o KeyTypeOutput) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return o.ToKeyTypePtrOutputWithContext(context.Background())
}

func (o KeyTypeOutput) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyType) *KeyType {
		return &v
	}).(KeyTypePtrOutput)
}

func (o KeyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeyTypePtrOutput struct{ *pulumi.OutputState }

func (KeyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyType)(nil)).Elem()
}

func (o KeyTypePtrOutput) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return o
}

func (o KeyTypePtrOutput) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return o
}

func (o KeyTypePtrOutput) Elem() KeyTypeOutput {
	return o.ApplyT(func(v *KeyType) KeyType {
		if v != nil {
			return *v
		}
		var ret KeyType
		return ret
	}).(KeyTypeOutput)
}

func (o KeyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// KeyTypeInput is an input type that accepts values of the KeyType enum
// A concrete instance of `KeyTypeInput` can be one of the following:
//
//	KeyTypeS3
//	KeyTypeSwift
type KeyTypeInput interface {
	pulumi.Input

	ToKeyTypeOutput() KeyTypeOutput
	ToKeyTypeOutputWithContext(context.Context) KeyTypeOutput
}

var keyTypePtrType = reflect.TypeOf((**KeyType)(nil)).Elem()

type KeyTypePtrInput interface {
	pulumi.Input

	ToKeyTypePtrOutput() KeyTypePtrOutput
	ToKeyTypePtrOutputWithContext(context.Context) KeyTypePtrOutput
}

type keyTypePtr string

func KeyTypePtr(v string) KeyTypePtrInput {
	return (*keyTypePtr)(&v)
}

func (*keyTypePtr) ElementType() reflect.Type {
	return keyTypePtrType
}

func (in *keyTypePtr) ToKeyTypePtrOutput() KeyTypePtrOutput {
	return pulumi.ToOutput(in).(KeyTypePtrOutput)
}

func (in *keyTypePtr) ToKeyTypePtrOutputWithContext(ctx context.Context) KeyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeyTypePtrOutput)
}

func (in *keyTypePtr) ToOutput(ctx context.Context) pulumix.Output[*KeyType] {
	return pulumix.Output[*KeyType]{
		OutputState: in.ToKeyTypePtrOutputWithContext(ctx).OutputState,
	}
}

type SubUserPermission string

const (
	SubUserPermissionNone        = SubUserPermission("none")
	SubUserPermissionRead        = SubUserPermission("read")
	SubUserPermissionWrite       = SubUserPermission("write")
	SubUserPermissionReadWrite   = SubUserPermission("readWrite")
	SubUserPermissionFullControl = SubUserPermission("fullControl")
)

func (SubUserPermission) ElementType() reflect.Type {
	return reflect.TypeOf((*SubUserPermission)(nil)).Elem()
}

func (e SubUserPermission) ToSubUserPermissionOutput() SubUserPermissionOutput {
	return pulumi.ToOutput(e).(SubUserPermissionOutput)
}

func (e SubUserPermission) ToSubUserPermissionOutputWithContext(ctx context.Context) SubUserPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SubUserPermissionOutput)
}

func (e SubUserPermission) ToSubUserPermissionPtrOutput() SubUserPermissionPtrOutput {
	return e.ToSubUserPermissionPtrOutputWithContext(context.Background())
}

func (e SubUserPermission) ToSubUserPermissionPtrOutputWithContext(ctx context.Context) SubUserPermissionPtrOutput {
	return SubUserPermission(e).ToSubUserPermissionOutputWithContext(ctx).ToSubUserPermissionPtrOutputWithContext(ctx)
}

func (e SubUserPermission) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubUserPermission) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubUserPermission) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SubUserPermission) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SubUserPermissionOutput struct{ *pulumi.OutputState }

func (SubUserPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubUserPermission)(nil)).Elem()
}

func (o SubUserPermissionOutput) ToSubUserPermissionOutput() SubUserPermissionOutput {
	return o
}

func (o SubUserPermissionOutput) ToSubUserPermissionOutputWithContext(ctx context.Context) SubUserPermissionOutput {
	return o
}

func (o SubUserPermissionOutput) ToSubUserPermissionPtrOutput() SubUserPermissionPtrOutput {
	return o.ToSubUserPermissionPtrOutputWithContext(context.Background())
}

func (o SubUserPermissionOutput) ToSubUserPermissionPtrOutputWithContext(ctx context.Context) SubUserPermissionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubUserPermission) *SubUserPermission {
		return &v
	}).(SubUserPermissionPtrOutput)
}

func (o SubUserPermissionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SubUserPermissionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubUserPermission) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SubUserPermissionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubUserPermissionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubUserPermission) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SubUserPermissionPtrOutput struct{ *pulumi.OutputState }

func (SubUserPermissionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubUserPermission)(nil)).Elem()
}

func (o SubUserPermissionPtrOutput) ToSubUserPermissionPtrOutput() SubUserPermissionPtrOutput {
	return o
}

func (o SubUserPermissionPtrOutput) ToSubUserPermissionPtrOutputWithContext(ctx context.Context) SubUserPermissionPtrOutput {
	return o
}

func (o SubUserPermissionPtrOutput) Elem() SubUserPermissionOutput {
	return o.ApplyT(func(v *SubUserPermission) SubUserPermission {
		if v != nil {
			return *v
		}
		var ret SubUserPermission
		return ret
	}).(SubUserPermissionOutput)
}

func (o SubUserPermissionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubUserPermissionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SubUserPermission) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SubUserPermissionInput is an input type that accepts values of the SubUserPermission enum
// A concrete instance of `SubUserPermissionInput` can be one of the following:
//
//	SubUserPermissionNone
//	SubUserPermissionRead
//	SubUserPermissionWrite
//	SubUserPermissionReadWrite
//	SubUserPermissionFullControl
type SubUserPermissionInput interface {
	pulumi.Input

	ToSubUserPermissionOutput() SubUserPermissionOutput
	ToSubUserPermissionOutputWithContext(context.Context) SubUserPermissionOutput
}

var subUserPermissionPtrType = reflect.TypeOf((**SubUserPermission)(nil)).Elem()

type SubUserPermissionPtrInput interface {
	pulumi.Input

	ToSubUserPermissionPtrOutput() SubUserPermissionPtrOutput
	ToSubUserPermissionPtrOutputWithContext(context.Context) SubUserPermissionPtrOutput
}

type subUserPermissionPtr string

func SubUserPermissionPtr(v string) SubUserPermissionPtrInput {
	return (*subUserPermissionPtr)(&v)
}

func (*subUserPermissionPtr) ElementType() reflect.Type {
	return subUserPermissionPtrType
}

func (in *subUserPermissionPtr) ToSubUserPermissionPtrOutput() SubUserPermissionPtrOutput {
	return pulumi.ToOutput(in).(SubUserPermissionPtrOutput)
}

func (in *subUserPermissionPtr) ToSubUserPermissionPtrOutputWithContext(ctx context.Context) SubUserPermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SubUserPermissionPtrOutput)
}

func (in *subUserPermissionPtr) ToOutput(ctx context.Context) pulumix.Output[*SubUserPermission] {
	return pulumix.Output[*SubUserPermission]{
		OutputState: in.ToSubUserPermissionPtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CapabilityPermissionInput)(nil)).Elem(), CapabilityPermission("read"))
	pulumi.RegisterInputType(reflect.TypeOf((*CapabilityPermissionPtrInput)(nil)).Elem(), CapabilityPermission("read"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeyTypeInput)(nil)).Elem(), KeyType("s3"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeyTypePtrInput)(nil)).Elem(), KeyType("s3"))
	pulumi.RegisterInputType(reflect.TypeOf((*SubUserPermissionInput)(nil)).Elem(), SubUserPermission("none"))
	pulumi.RegisterInputType(reflect.TypeOf((*SubUserPermissionPtrInput)(nil)).Elem(), SubUserPermission("none"))
	pulumi.RegisterOutputType(CapabilityPermissionOutput{})
	pulumi.RegisterOutputType(CapabilityPermissionPtrOutput{})
	pulumi.RegisterOutputType(KeyTypeOutput{})
	pulumi.RegisterOutputType(KeyTypePtrOutput{})
	pulumi.RegisterOutputType(SubUserPermissionOutput{})
	pulumi.RegisterOutputType(SubUserPermissionPtrOutput{})
}
