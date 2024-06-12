// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

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
// MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change the object. Deprecated in v1.16, planned for removal in v1.19. Use admissionregistration.k8s.io/v1 MutatingWebhookConfiguration instead.
type MutatingWebhookConfigurationPatch struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPatchPtrOutput `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks MutatingWebhookPatchArrayOutput `pulumi:"webhooks"`
}

// NewMutatingWebhookConfigurationPatch registers a new resource with the given unique name, arguments, and options.
func NewMutatingWebhookConfigurationPatch(ctx *pulumi.Context,
	name string, args *MutatingWebhookConfigurationPatchArgs, opts ...pulumi.ResourceOption) (*MutatingWebhookConfigurationPatch, error) {
	if args == nil {
		args = &MutatingWebhookConfigurationPatchArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("admissionregistration.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("MutatingWebhookConfiguration")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:admissionregistration.k8s.io/v1:MutatingWebhookConfigurationPatch"),
		},
	})
	opts = append(opts, aliases)
	var resource MutatingWebhookConfigurationPatch
	err := ctx.RegisterResource("kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfigurationPatch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMutatingWebhookConfigurationPatch gets an existing MutatingWebhookConfigurationPatch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMutatingWebhookConfigurationPatch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MutatingWebhookConfigurationPatchState, opts ...pulumi.ResourceOption) (*MutatingWebhookConfigurationPatch, error) {
	var resource MutatingWebhookConfigurationPatch
	err := ctx.ReadResource("kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfigurationPatch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MutatingWebhookConfigurationPatch resources.
type mutatingWebhookConfigurationPatchState struct {
}

type MutatingWebhookConfigurationPatchState struct {
}

func (MutatingWebhookConfigurationPatchState) ElementType() reflect.Type {
	return reflect.TypeOf((*mutatingWebhookConfigurationPatchState)(nil)).Elem()
}

type mutatingWebhookConfigurationPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *metav1.ObjectMetaPatch `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks []MutatingWebhookPatch `pulumi:"webhooks"`
}

// The set of arguments for constructing a MutatingWebhookConfigurationPatch resource.
type MutatingWebhookConfigurationPatchArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPatchPtrInput
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks MutatingWebhookPatchArrayInput
}

func (MutatingWebhookConfigurationPatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mutatingWebhookConfigurationPatchArgs)(nil)).Elem()
}

type MutatingWebhookConfigurationPatchInput interface {
	pulumi.Input

	ToMutatingWebhookConfigurationPatchOutput() MutatingWebhookConfigurationPatchOutput
	ToMutatingWebhookConfigurationPatchOutputWithContext(ctx context.Context) MutatingWebhookConfigurationPatchOutput
}

func (*MutatingWebhookConfigurationPatch) ElementType() reflect.Type {
	return reflect.TypeOf((**MutatingWebhookConfigurationPatch)(nil)).Elem()
}

func (i *MutatingWebhookConfigurationPatch) ToMutatingWebhookConfigurationPatchOutput() MutatingWebhookConfigurationPatchOutput {
	return i.ToMutatingWebhookConfigurationPatchOutputWithContext(context.Background())
}

func (i *MutatingWebhookConfigurationPatch) ToMutatingWebhookConfigurationPatchOutputWithContext(ctx context.Context) MutatingWebhookConfigurationPatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MutatingWebhookConfigurationPatchOutput)
}

// MutatingWebhookConfigurationPatchArrayInput is an input type that accepts MutatingWebhookConfigurationPatchArray and MutatingWebhookConfigurationPatchArrayOutput values.
// You can construct a concrete instance of `MutatingWebhookConfigurationPatchArrayInput` via:
//
//	MutatingWebhookConfigurationPatchArray{ MutatingWebhookConfigurationPatchArgs{...} }
type MutatingWebhookConfigurationPatchArrayInput interface {
	pulumi.Input

	ToMutatingWebhookConfigurationPatchArrayOutput() MutatingWebhookConfigurationPatchArrayOutput
	ToMutatingWebhookConfigurationPatchArrayOutputWithContext(context.Context) MutatingWebhookConfigurationPatchArrayOutput
}

type MutatingWebhookConfigurationPatchArray []MutatingWebhookConfigurationPatchInput

func (MutatingWebhookConfigurationPatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MutatingWebhookConfigurationPatch)(nil)).Elem()
}

func (i MutatingWebhookConfigurationPatchArray) ToMutatingWebhookConfigurationPatchArrayOutput() MutatingWebhookConfigurationPatchArrayOutput {
	return i.ToMutatingWebhookConfigurationPatchArrayOutputWithContext(context.Background())
}

func (i MutatingWebhookConfigurationPatchArray) ToMutatingWebhookConfigurationPatchArrayOutputWithContext(ctx context.Context) MutatingWebhookConfigurationPatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MutatingWebhookConfigurationPatchArrayOutput)
}

// MutatingWebhookConfigurationPatchMapInput is an input type that accepts MutatingWebhookConfigurationPatchMap and MutatingWebhookConfigurationPatchMapOutput values.
// You can construct a concrete instance of `MutatingWebhookConfigurationPatchMapInput` via:
//
//	MutatingWebhookConfigurationPatchMap{ "key": MutatingWebhookConfigurationPatchArgs{...} }
type MutatingWebhookConfigurationPatchMapInput interface {
	pulumi.Input

	ToMutatingWebhookConfigurationPatchMapOutput() MutatingWebhookConfigurationPatchMapOutput
	ToMutatingWebhookConfigurationPatchMapOutputWithContext(context.Context) MutatingWebhookConfigurationPatchMapOutput
}

type MutatingWebhookConfigurationPatchMap map[string]MutatingWebhookConfigurationPatchInput

func (MutatingWebhookConfigurationPatchMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MutatingWebhookConfigurationPatch)(nil)).Elem()
}

func (i MutatingWebhookConfigurationPatchMap) ToMutatingWebhookConfigurationPatchMapOutput() MutatingWebhookConfigurationPatchMapOutput {
	return i.ToMutatingWebhookConfigurationPatchMapOutputWithContext(context.Background())
}

func (i MutatingWebhookConfigurationPatchMap) ToMutatingWebhookConfigurationPatchMapOutputWithContext(ctx context.Context) MutatingWebhookConfigurationPatchMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MutatingWebhookConfigurationPatchMapOutput)
}

type MutatingWebhookConfigurationPatchOutput struct{ *pulumi.OutputState }

func (MutatingWebhookConfigurationPatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MutatingWebhookConfigurationPatch)(nil)).Elem()
}

func (o MutatingWebhookConfigurationPatchOutput) ToMutatingWebhookConfigurationPatchOutput() MutatingWebhookConfigurationPatchOutput {
	return o
}

func (o MutatingWebhookConfigurationPatchOutput) ToMutatingWebhookConfigurationPatchOutputWithContext(ctx context.Context) MutatingWebhookConfigurationPatchOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o MutatingWebhookConfigurationPatchOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MutatingWebhookConfigurationPatch) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o MutatingWebhookConfigurationPatchOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MutatingWebhookConfigurationPatch) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
func (o MutatingWebhookConfigurationPatchOutput) Metadata() metav1.ObjectMetaPatchPtrOutput {
	return o.ApplyT(func(v *MutatingWebhookConfigurationPatch) metav1.ObjectMetaPatchPtrOutput { return v.Metadata }).(metav1.ObjectMetaPatchPtrOutput)
}

// Webhooks is a list of webhooks and the affected resources and operations.
func (o MutatingWebhookConfigurationPatchOutput) Webhooks() MutatingWebhookPatchArrayOutput {
	return o.ApplyT(func(v *MutatingWebhookConfigurationPatch) MutatingWebhookPatchArrayOutput { return v.Webhooks }).(MutatingWebhookPatchArrayOutput)
}

type MutatingWebhookConfigurationPatchArrayOutput struct{ *pulumi.OutputState }

func (MutatingWebhookConfigurationPatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MutatingWebhookConfigurationPatch)(nil)).Elem()
}

func (o MutatingWebhookConfigurationPatchArrayOutput) ToMutatingWebhookConfigurationPatchArrayOutput() MutatingWebhookConfigurationPatchArrayOutput {
	return o
}

func (o MutatingWebhookConfigurationPatchArrayOutput) ToMutatingWebhookConfigurationPatchArrayOutputWithContext(ctx context.Context) MutatingWebhookConfigurationPatchArrayOutput {
	return o
}

func (o MutatingWebhookConfigurationPatchArrayOutput) Index(i pulumi.IntInput) MutatingWebhookConfigurationPatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MutatingWebhookConfigurationPatch {
		return vs[0].([]*MutatingWebhookConfigurationPatch)[vs[1].(int)]
	}).(MutatingWebhookConfigurationPatchOutput)
}

type MutatingWebhookConfigurationPatchMapOutput struct{ *pulumi.OutputState }

func (MutatingWebhookConfigurationPatchMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MutatingWebhookConfigurationPatch)(nil)).Elem()
}

func (o MutatingWebhookConfigurationPatchMapOutput) ToMutatingWebhookConfigurationPatchMapOutput() MutatingWebhookConfigurationPatchMapOutput {
	return o
}

func (o MutatingWebhookConfigurationPatchMapOutput) ToMutatingWebhookConfigurationPatchMapOutputWithContext(ctx context.Context) MutatingWebhookConfigurationPatchMapOutput {
	return o
}

func (o MutatingWebhookConfigurationPatchMapOutput) MapIndex(k pulumi.StringInput) MutatingWebhookConfigurationPatchOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MutatingWebhookConfigurationPatch {
		return vs[0].(map[string]*MutatingWebhookConfigurationPatch)[vs[1].(string)]
	}).(MutatingWebhookConfigurationPatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MutatingWebhookConfigurationPatchInput)(nil)).Elem(), &MutatingWebhookConfigurationPatch{})
	pulumi.RegisterInputType(reflect.TypeOf((*MutatingWebhookConfigurationPatchArrayInput)(nil)).Elem(), MutatingWebhookConfigurationPatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MutatingWebhookConfigurationPatchMapInput)(nil)).Elem(), MutatingWebhookConfigurationPatchMap{})
	pulumi.RegisterOutputType(MutatingWebhookConfigurationPatchOutput{})
	pulumi.RegisterOutputType(MutatingWebhookConfigurationPatchArrayOutput{})
	pulumi.RegisterOutputType(MutatingWebhookConfigurationPatchMapOutput{})
}