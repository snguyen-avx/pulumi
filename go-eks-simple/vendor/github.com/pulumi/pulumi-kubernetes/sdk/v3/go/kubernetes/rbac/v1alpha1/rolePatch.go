// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Patch resources are used to modify existing Kubernetes resources by using
// Server-Side Apply updates. The name of the resource must be specified, but all other properties are optional. More than
// one patch may be applied to the same resource, and a random FieldManager name will be used for each Patch resource.
// Conflicts will result in an error by default, but can be forced using the "pulumi.com/patchForce" annotation. See the
// [Server-Side Apply Docs](https://www.pulumi.com/registry/packages/kubernetes/how-to-guides/managing-resources-with-server-side-apply/) for
// additional information about using Server-Side Apply to manage Kubernetes resources with Pulumi.
// Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 Role, and will no longer be served in v1.20.
type RolePatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Rules holds all the PolicyRules for this Role
	Rules PolicyRulePatchArrayOutput `pulumi:"rules"`
}

// NewRolePatch registers a new resource with the given unique name, arguments, and options.
func NewRolePatch(ctx *pulumi.Context,
	name string, args *RolePatchArgs, opts ...pulumi.ResourceOption) (*RolePatch, error) {
	if args == nil {
		args = &RolePatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("Role")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1:RolePatch"),
		},
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1beta1:RolePatch"),
		},
	})
	opts = append(opts, aliases)
	var resource RolePatch
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1alpha1:RolePatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRolePatch gets an existing RolePatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRolePatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RolePatchState, opts ...pulumi.ResourceOption) (*RolePatch, error) {
	var resource RolePatch
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1alpha1:RolePatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RolePatch resources.
type rolePatchState struct {
}

type RolePatchState struct {
}

func (RolePatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*rolePatchState)(nil)).Elem()
}

type rolePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Rules holds all the PolicyRules for this Role
	Rules []PolicyRulePatch `pulumi:"rules"`
}

// The set of arguments for constructing a RolePatch resource.
type RolePatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPatchPtrInput
	// Rules holds all the PolicyRules for this Role
	Rules PolicyRulePatchArrayInput
}

func (RolePatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rolePatchArgs)(nil)).Elem()
}

type RolePatchInput interface {
	pulumi.Input

	ToRolePatchOutput() RolePatchOutput
	ToRolePatchOutputWithContext(ctx context.Context) RolePatchOutput
}

func (*RolePatch) ElementType() reflect.Type {
	return reflect.TypeOf((**RolePatch)(nil)).Elem()
}

func (i *RolePatch) ToRolePatchOutput() RolePatchOutput {
	return i.ToRolePatchOutputWithContext(context.Background())
}

func (i *RolePatch) ToRolePatchOutputWithContext(ctx context.Context) RolePatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePatchOutput)
}

// RolePatchArrayInput is an input type that accepts RolePatchArray and RolePatchArrayOutput values.
// You can construct a concrete instance of `RolePatchArrayInput` via:
//
//	RolePatchArray{ RolePatchArgs{...} }
type RolePatchArrayInput interface {
	pulumi.Input

	ToRolePatchArrayOutput() RolePatchArrayOutput
	ToRolePatchArrayOutputWithContext(context.Context) RolePatchArrayOutput
}

type RolePatchArray []RolePatchInput

func (RolePatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RolePatch)(nil)).Elem()
}

func (i RolePatchArray) ToRolePatchArrayOutput() RolePatchArrayOutput {
	return i.ToRolePatchArrayOutputWithContext(context.Background())
}

func (i RolePatchArray) ToRolePatchArrayOutputWithContext(ctx context.Context) RolePatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePatchArrayOutput)
}

// RolePatchMapInput is an input type that accepts RolePatchMap and RolePatchMapOutput values.
// You can construct a concrete instance of `RolePatchMapInput` via:
//
//	RolePatchMap{ "key": RolePatchArgs{...} }
type RolePatchMapInput interface {
	pulumi.Input

	ToRolePatchMapOutput() RolePatchMapOutput
	ToRolePatchMapOutputWithContext(context.Context) RolePatchMapOutput
}

type RolePatchMap map[string]RolePatchInput

func (RolePatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RolePatch)(nil)).Elem()
}

func (i RolePatchMap) ToRolePatchMapOutput() RolePatchMapOutput {
	return i.ToRolePatchMapOutputWithContext(context.Background())
}

func (i RolePatchMap) ToRolePatchMapOutputWithContext(ctx context.Context) RolePatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePatchMapOutput)
}

type RolePatchOutput struct{ *pulumi.OutputState }

func (RolePatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RolePatch)(nil)).Elem()
}

func (o RolePatchOutput) ToRolePatchOutput() RolePatchOutput {
	return o
}

func (o RolePatchOutput) ToRolePatchOutputWithContext(ctx context.Context) RolePatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o RolePatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RolePatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o RolePatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RolePatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata.
func (o RolePatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *RolePatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Rules holds all the PolicyRules for this Role
func (o RolePatchOutput) Rules() PolicyRulePatchArrayOutput {
	return o.ApplyT(func(v *RolePatch) PolicyRulePatchArrayOutput { return v.Rules }).(PolicyRulePatchArrayOutput)
}

type RolePatchArrayOutput struct{ *pulumi.OutputState }

func (RolePatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RolePatch)(nil)).Elem()
}

func (o RolePatchArrayOutput) ToRolePatchArrayOutput() RolePatchArrayOutput {
	return o
}

func (o RolePatchArrayOutput) ToRolePatchArrayOutputWithContext(ctx context.Context) RolePatchArrayOutput {
	return o
}

func (o RolePatchArrayOutput) Index(i pulumi.IntInput) RolePatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RolePatch {
		return vs[0].([]*RolePatch)[vs[1].(int)]
	}).(RolePatchOutput)
}

type RolePatchMapOutput struct{ *pulumi.OutputState }

func (RolePatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RolePatch)(nil)).Elem()
}

func (o RolePatchMapOutput) ToRolePatchMapOutput() RolePatchMapOutput {
	return o
}

func (o RolePatchMapOutput) ToRolePatchMapOutputWithContext(ctx context.Context) RolePatchMapOutput {
	return o
}

func (o RolePatchMapOutput) MapIndex(k pulumi.StringInput) RolePatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RolePatch {
		return vs[0].(map[string]*RolePatch)[vs[1].(string)]
	}).(RolePatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RolePatchInput)(nil)).Elem(), &RolePatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*RolePatchArrayInput)(nil)).Elem(), RolePatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RolePatchMapInput)(nil)).Elem(), RolePatchMap{})
	pulumi.RegisterOutputType(RolePatchOutput{})
	pulumi.RegisterOutputType(RolePatchArrayOutput{})
	pulumi.RegisterOutputType(RolePatchMapOutput{})
}
