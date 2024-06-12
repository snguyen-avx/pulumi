// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ConfigMap holds configuration data for pods to consume.
type ConfigMap struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData pulumi.StringMapOutput `pulumi:"binaryData"`
	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
	Data pulumi.StringMapOutput `pulumi:"data"`
	// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable pulumi.BoolOutput `pulumi:"immutable"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
}

// NewConfigMap registers a new resource with the given unique name, arguments, and options.
func NewConfigMap(ctx *pulumi.Context,
	name string, args *ConfigMapArgs, opts ...pulumi.ResourceOption) (*ConfigMap, error) {
	if args == nil {
		args = &ConfigMapArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("ConfigMap")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigMap
	err := ctx.RegisterResource("kubernetes:core/v1:ConfigMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigMap gets an existing ConfigMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigMapState, opts ...pulumi.ResourceOption) (*ConfigMap, error) {
	var resource ConfigMap
	err := ctx.ReadResource("kubernetes:core/v1:ConfigMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigMap resources.
type configMapState struct {
}

type ConfigMapState struct {
}

func (ConfigMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*configMapState)(nil)).Elem()
}

type configMapArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData map[string]string `pulumi:"binaryData"`
	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
	Data map[string]string `pulumi:"data"`
	// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable *bool `pulumi:"immutable"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a ConfigMap resource.
type ConfigMapArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
	BinaryData pulumi.StringMapInput
	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
	Data pulumi.StringMapInput
	// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable pulumi.BoolPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
}

func (ConfigMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configMapArgs)(nil)).Elem()
}

type ConfigMapInput interface {
	pulumi.Input

	ToConfigMapOutput() ConfigMapOutput
	ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput
}

func (*ConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigMap)(nil)).Elem()
}

func (i *ConfigMap) ToConfigMapOutput() ConfigMapOutput {
	return i.ToConfigMapOutputWithContext(context.Background())
}

func (i *ConfigMap) ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapOutput)
}

// ConfigMapArrayInput is an input type that accepts ConfigMapArray and ConfigMapArrayOutput values.
// You can construct a concrete instance of `ConfigMapArrayInput` via:
//
//	ConfigMapArray{ ConfigMapArgs{...} }
type ConfigMapArrayInput interface {
	pulumi.Input

	ToConfigMapArrayOutput() ConfigMapArrayOutput
	ToConfigMapArrayOutputWithContext(context.Context) ConfigMapArrayOutput
}

type ConfigMapArray []ConfigMapInput

func (ConfigMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigMap)(nil)).Elem()
}

func (i ConfigMapArray) ToConfigMapArrayOutput() ConfigMapArrayOutput {
	return i.ToConfigMapArrayOutputWithContext(context.Background())
}

func (i ConfigMapArray) ToConfigMapArrayOutputWithContext(ctx context.Context) ConfigMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapArrayOutput)
}

// ConfigMapMapInput is an input type that accepts ConfigMapMap and ConfigMapMapOutput values.
// You can construct a concrete instance of `ConfigMapMapInput` via:
//
//	ConfigMapMap{ "key": ConfigMapArgs{...} }
type ConfigMapMapInput interface {
	pulumi.Input

	ToConfigMapMapOutput() ConfigMapMapOutput
	ToConfigMapMapOutputWithContext(context.Context) ConfigMapMapOutput
}

type ConfigMapMap map[string]ConfigMapInput

func (ConfigMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigMap)(nil)).Elem()
}

func (i ConfigMapMap) ToConfigMapMapOutput() ConfigMapMapOutput {
	return i.ToConfigMapMapOutputWithContext(context.Background())
}

func (i ConfigMapMap) ToConfigMapMapOutputWithContext(ctx context.Context) ConfigMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigMapMapOutput)
}

type ConfigMapOutput struct{ *pulumi.OutputState }

func (ConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigMap)(nil)).Elem()
}

func (o ConfigMapOutput) ToConfigMapOutput() ConfigMapOutput {
	return o
}

func (o ConfigMapOutput) ToConfigMapOutputWithContext(ctx context.Context) ConfigMapOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ConfigMapOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigMap) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
func (o ConfigMapOutput) BinaryData() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigMap) pulumi.StringMapOutput { return v.BinaryData }).(pulumi.StringMapOutput)
}

// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
func (o ConfigMapOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigMap) pulumi.StringMapOutput { return v.Data }).(pulumi.StringMapOutput)
}

// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
func (o ConfigMapOutput) Immutable() pulumi.BoolOutput {
	return o.ApplyT(func(v *ConfigMap) pulumi.BoolOutput { return v.Immutable }).(pulumi.BoolOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ConfigMapOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigMap) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o ConfigMapOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v *ConfigMap) metav1.ObjectMetaOutput { return v.Metadata }).(metav1.ObjectMetaOutput)
}

type ConfigMapArrayOutput struct{ *pulumi.OutputState }

func (ConfigMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigMap)(nil)).Elem()
}

func (o ConfigMapArrayOutput) ToConfigMapArrayOutput() ConfigMapArrayOutput {
	return o
}

func (o ConfigMapArrayOutput) ToConfigMapArrayOutputWithContext(ctx context.Context) ConfigMapArrayOutput {
	return o
}

func (o ConfigMapArrayOutput) Index(i pulumi.IntInput) ConfigMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigMap {
		return vs[0].([]*ConfigMap)[vs[1].(int)]
	}).(ConfigMapOutput)
}

type ConfigMapMapOutput struct{ *pulumi.OutputState }

func (ConfigMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigMap)(nil)).Elem()
}

func (o ConfigMapMapOutput) ToConfigMapMapOutput() ConfigMapMapOutput {
	return o
}

func (o ConfigMapMapOutput) ToConfigMapMapOutputWithContext(ctx context.Context) ConfigMapMapOutput {
	return o
}

func (o ConfigMapMapOutput) MapIndex(k pulumi.StringInput) ConfigMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigMap {
		return vs[0].(map[string]*ConfigMap)[vs[1].(string)]
	}).(ConfigMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapInput)(nil)).Elem(), &ConfigMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapArrayInput)(nil)).Elem(), ConfigMapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigMapMapInput)(nil)).Elem(), ConfigMapMap{})
	pulumi.RegisterOutputType(ConfigMapOutput{})
	pulumi.RegisterOutputType(ConfigMapArrayOutput{})
	pulumi.RegisterOutputType(ConfigMapMapOutput{})
}
