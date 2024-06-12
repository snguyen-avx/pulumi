// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CustomResourceDefinitionList is a list of CustomResourceDefinition objects.
type CustomResourceDefinitionList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// items list individual CustomResourceDefinition objects
	Items CustomResourceDefinitionTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewCustomResourceDefinitionList registers a new resource with the given unique name, arguments, and options.
func NewCustomResourceDefinitionList(ctx *pulumi.Context,
	name string, args *CustomResourceDefinitionListArgs, opts ...pulumi.ResourceOption) (*CustomResourceDefinitionList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("apiextensions.k8s.io/v1")
	args.Kind = pulumi.StringPtr("CustomResourceDefinitionList")
	var resource CustomResourceDefinitionList
	err := ctx.RegisterResource("kubernetes:apiextensions.k8s.io/v1:CustomResourceDefinitionList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomResourceDefinitionList gets an existing CustomResourceDefinitionList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomResourceDefinitionList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomResourceDefinitionListState, opts ...pulumi.ResourceOption) (*CustomResourceDefinitionList, error) {
	var resource CustomResourceDefinitionList
	err := ctx.ReadResource("kubernetes:apiextensions.k8s.io/v1:CustomResourceDefinitionList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomResourceDefinitionList resources.
type customResourceDefinitionListState struct {
}

type CustomResourceDefinitionListState struct {
}

func (CustomResourceDefinitionListState) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionListState)(nil)).Elem()
}

type customResourceDefinitionListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// items list individual CustomResourceDefinition objects
	Items []CustomResourceDefinitionType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a CustomResourceDefinitionList resource.
type CustomResourceDefinitionListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// items list individual CustomResourceDefinition objects
	Items CustomResourceDefinitionTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (CustomResourceDefinitionListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customResourceDefinitionListArgs)(nil)).Elem()
}

type CustomResourceDefinitionListInput interface {
	pulumi.Input

	ToCustomResourceDefinitionListOutput() CustomResourceDefinitionListOutput
	ToCustomResourceDefinitionListOutputWithContext(ctx context.Context) CustomResourceDefinitionListOutput
}

func (*CustomResourceDefinitionList) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomResourceDefinitionList)(nil)).Elem()
}

func (i *CustomResourceDefinitionList) ToCustomResourceDefinitionListOutput() CustomResourceDefinitionListOutput {
	return i.ToCustomResourceDefinitionListOutputWithContext(context.Background())
}

func (i *CustomResourceDefinitionList) ToCustomResourceDefinitionListOutputWithContext(ctx context.Context) CustomResourceDefinitionListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceDefinitionListOutput)
}

// CustomResourceDefinitionListArrayInput is an input type that accepts CustomResourceDefinitionListArray and CustomResourceDefinitionListArrayOutput values.
// You can construct a concrete instance of `CustomResourceDefinitionListArrayInput` via:
//
//	CustomResourceDefinitionListArray{ CustomResourceDefinitionListArgs{...} }
type CustomResourceDefinitionListArrayInput interface {
	pulumi.Input

	ToCustomResourceDefinitionListArrayOutput() CustomResourceDefinitionListArrayOutput
	ToCustomResourceDefinitionListArrayOutputWithContext(context.Context) CustomResourceDefinitionListArrayOutput
}

type CustomResourceDefinitionListArray []CustomResourceDefinitionListInput

func (CustomResourceDefinitionListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomResourceDefinitionList)(nil)).Elem()
}

func (i CustomResourceDefinitionListArray) ToCustomResourceDefinitionListArrayOutput() CustomResourceDefinitionListArrayOutput {
	return i.ToCustomResourceDefinitionListArrayOutputWithContext(context.Background())
}

func (i CustomResourceDefinitionListArray) ToCustomResourceDefinitionListArrayOutputWithContext(ctx context.Context) CustomResourceDefinitionListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceDefinitionListArrayOutput)
}

// CustomResourceDefinitionListMapInput is an input type that accepts CustomResourceDefinitionListMap and CustomResourceDefinitionListMapOutput values.
// You can construct a concrete instance of `CustomResourceDefinitionListMapInput` via:
//
//	CustomResourceDefinitionListMap{ "key": CustomResourceDefinitionListArgs{...} }
type CustomResourceDefinitionListMapInput interface {
	pulumi.Input

	ToCustomResourceDefinitionListMapOutput() CustomResourceDefinitionListMapOutput
	ToCustomResourceDefinitionListMapOutputWithContext(context.Context) CustomResourceDefinitionListMapOutput
}

type CustomResourceDefinitionListMap map[string]CustomResourceDefinitionListInput

func (CustomResourceDefinitionListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomResourceDefinitionList)(nil)).Elem()
}

func (i CustomResourceDefinitionListMap) ToCustomResourceDefinitionListMapOutput() CustomResourceDefinitionListMapOutput {
	return i.ToCustomResourceDefinitionListMapOutputWithContext(context.Background())
}

func (i CustomResourceDefinitionListMap) ToCustomResourceDefinitionListMapOutputWithContext(ctx context.Context) CustomResourceDefinitionListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomResourceDefinitionListMapOutput)
}

type CustomResourceDefinitionListOutput struct{ *pulumi.OutputState }

func (CustomResourceDefinitionListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomResourceDefinitionList)(nil)).Elem()
}

func (o CustomResourceDefinitionListOutput) ToCustomResourceDefinitionListOutput() CustomResourceDefinitionListOutput {
	return o
}

func (o CustomResourceDefinitionListOutput) ToCustomResourceDefinitionListOutputWithContext(ctx context.Context) CustomResourceDefinitionListOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CustomResourceDefinitionListOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinitionList) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// items list individual CustomResourceDefinition objects
func (o CustomResourceDefinitionListOutput) Items() CustomResourceDefinitionTypeArrayOutput {
	return o.ApplyT(func(v *CustomResourceDefinitionList) CustomResourceDefinitionTypeArrayOutput { return v.Items }).(CustomResourceDefinitionTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CustomResourceDefinitionListOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinitionList) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o CustomResourceDefinitionListOutput) Metadata() metav1.ListMetaPtrOutput {
	return o.ApplyT(func(v *CustomResourceDefinitionList) metav1.ListMetaPtrOutput { return v.Metadata }).(metav1.ListMetaPtrOutput)
}

type CustomResourceDefinitionListArrayOutput struct{ *pulumi.OutputState }

func (CustomResourceDefinitionListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomResourceDefinitionList)(nil)).Elem()
}

func (o CustomResourceDefinitionListArrayOutput) ToCustomResourceDefinitionListArrayOutput() CustomResourceDefinitionListArrayOutput {
	return o
}

func (o CustomResourceDefinitionListArrayOutput) ToCustomResourceDefinitionListArrayOutputWithContext(ctx context.Context) CustomResourceDefinitionListArrayOutput {
	return o
}

func (o CustomResourceDefinitionListArrayOutput) Index(i pulumi.IntInput) CustomResourceDefinitionListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomResourceDefinitionList {
		return vs[0].([]*CustomResourceDefinitionList)[vs[1].(int)]
	}).(CustomResourceDefinitionListOutput)
}

type CustomResourceDefinitionListMapOutput struct{ *pulumi.OutputState }

func (CustomResourceDefinitionListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomResourceDefinitionList)(nil)).Elem()
}

func (o CustomResourceDefinitionListMapOutput) ToCustomResourceDefinitionListMapOutput() CustomResourceDefinitionListMapOutput {
	return o
}

func (o CustomResourceDefinitionListMapOutput) ToCustomResourceDefinitionListMapOutputWithContext(ctx context.Context) CustomResourceDefinitionListMapOutput {
	return o
}

func (o CustomResourceDefinitionListMapOutput) MapIndex(k pulumi.StringInput) CustomResourceDefinitionListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomResourceDefinitionList {
		return vs[0].(map[string]*CustomResourceDefinitionList)[vs[1].(string)]
	}).(CustomResourceDefinitionListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomResourceDefinitionListInput)(nil)).Elem(), &CustomResourceDefinitionList{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomResourceDefinitionListArrayInput)(nil)).Elem(), CustomResourceDefinitionListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomResourceDefinitionListMapInput)(nil)).Elem(), CustomResourceDefinitionListMap{})
	pulumi.RegisterOutputType(CustomResourceDefinitionListOutput{})
	pulumi.RegisterOutputType(CustomResourceDefinitionListArrayOutput{})
	pulumi.RegisterOutputType(CustomResourceDefinitionListMapOutput{})
}