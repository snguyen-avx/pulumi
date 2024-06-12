// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ClusterTrustBundle is a cluster-scoped container for X.509 trust anchors (root certificates).
//
// ClusterTrustBundle objects are considered to be readable by any authenticated user in the cluster, because they can be mounted by pods using the `clusterTrustBundle` projection.  All service accounts have read access to ClusterTrustBundles by default.  Users who only have namespace-level access to a cluster can read ClusterTrustBundles by impersonating a serviceaccount that they have access to.
//
// It can be optionally associated with a particular assigner, in which case it contains one valid set of trust anchors for that signer. Signers may have multiple associated ClusterTrustBundles; each is an independent set of trust anchors for that signer. Admission control is used to enforce that only users with permissions on the signer can create or modify the corresponding bundle.
type ClusterTrustBundle struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// metadata contains the object metadata.
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// spec contains the signer (if any) and trust anchors.
	Spec ClusterTrustBundleSpecOutput `pulumi:"spec"`
}

// NewClusterTrustBundle registers a new resource with the given unique name, arguments, and options.
func NewClusterTrustBundle(ctx *pulumi.Context,
	name string, args *ClusterTrustBundleArgs, opts ...pulumi.ResourceOption) (*ClusterTrustBundle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("certificates.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ClusterTrustBundle")
	var resource ClusterTrustBundle
	err := ctx.RegisterResource("kubernetes:certificates.k8s.io/v1alpha1:ClusterTrustBundle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterTrustBundle gets an existing ClusterTrustBundle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterTrustBundle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterTrustBundleState, opts ...pulumi.ResourceOption) (*ClusterTrustBundle, error) {
	var resource ClusterTrustBundle
	err := ctx.ReadResource("kubernetes:certificates.k8s.io/v1alpha1:ClusterTrustBundle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterTrustBundle resources.
type clusterTrustBundleState struct {
}

type ClusterTrustBundleState struct {
}

func (ClusterTrustBundleState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterTrustBundleState)(nil)).Elem()
}

type clusterTrustBundleArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// metadata contains the object metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec contains the signer (if any) and trust anchors.
	Spec ClusterTrustBundleSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ClusterTrustBundle resource.
type ClusterTrustBundleArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata contains the object metadata.
	Metadata metav1.ObjectMetaPtrInput
	// spec contains the signer (if any) and trust anchors.
	Spec ClusterTrustBundleSpecInput
}

func (ClusterTrustBundleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterTrustBundleArgs)(nil)).Elem()
}

type ClusterTrustBundleInput interface {
	pulumi.Input

	ToClusterTrustBundleOutput() ClusterTrustBundleOutput
	ToClusterTrustBundleOutputWithContext(ctx context.Context) ClusterTrustBundleOutput
}

func (*ClusterTrustBundle) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterTrustBundle)(nil)).Elem()
}

func (i *ClusterTrustBundle) ToClusterTrustBundleOutput() ClusterTrustBundleOutput {
	return i.ToClusterTrustBundleOutputWithContext(context.Background())
}

func (i *ClusterTrustBundle) ToClusterTrustBundleOutputWithContext(ctx context.Context) ClusterTrustBundleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTrustBundleOutput)
}

// ClusterTrustBundleArrayInput is an input type that accepts ClusterTrustBundleArray and ClusterTrustBundleArrayOutput values.
// You can construct a concrete instance of `ClusterTrustBundleArrayInput` via:
//
//	ClusterTrustBundleArray{ ClusterTrustBundleArgs{...} }
type ClusterTrustBundleArrayInput interface {
	pulumi.Input

	ToClusterTrustBundleArrayOutput() ClusterTrustBundleArrayOutput
	ToClusterTrustBundleArrayOutputWithContext(context.Context) ClusterTrustBundleArrayOutput
}

type ClusterTrustBundleArray []ClusterTrustBundleInput

func (ClusterTrustBundleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterTrustBundle)(nil)).Elem()
}

func (i ClusterTrustBundleArray) ToClusterTrustBundleArrayOutput() ClusterTrustBundleArrayOutput {
	return i.ToClusterTrustBundleArrayOutputWithContext(context.Background())
}

func (i ClusterTrustBundleArray) ToClusterTrustBundleArrayOutputWithContext(ctx context.Context) ClusterTrustBundleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTrustBundleArrayOutput)
}

// ClusterTrustBundleMapInput is an input type that accepts ClusterTrustBundleMap and ClusterTrustBundleMapOutput values.
// You can construct a concrete instance of `ClusterTrustBundleMapInput` via:
//
//	ClusterTrustBundleMap{ "key": ClusterTrustBundleArgs{...} }
type ClusterTrustBundleMapInput interface {
	pulumi.Input

	ToClusterTrustBundleMapOutput() ClusterTrustBundleMapOutput
	ToClusterTrustBundleMapOutputWithContext(context.Context) ClusterTrustBundleMapOutput
}

type ClusterTrustBundleMap map[string]ClusterTrustBundleInput

func (ClusterTrustBundleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterTrustBundle)(nil)).Elem()
}

func (i ClusterTrustBundleMap) ToClusterTrustBundleMapOutput() ClusterTrustBundleMapOutput {
	return i.ToClusterTrustBundleMapOutputWithContext(context.Background())
}

func (i ClusterTrustBundleMap) ToClusterTrustBundleMapOutputWithContext(ctx context.Context) ClusterTrustBundleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTrustBundleMapOutput)
}

type ClusterTrustBundleOutput struct{ *pulumi.OutputState }

func (ClusterTrustBundleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterTrustBundle)(nil)).Elem()
}

func (o ClusterTrustBundleOutput) ToClusterTrustBundleOutput() ClusterTrustBundleOutput {
	return o
}

func (o ClusterTrustBundleOutput) ToClusterTrustBundleOutputWithContext(ctx context.Context) ClusterTrustBundleOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o ClusterTrustBundleOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterTrustBundle) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o ClusterTrustBundleOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterTrustBundle) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// metadata contains the object metadata.
func (o ClusterTrustBundleOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *ClusterTrustBundle) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// spec contains the signer (if any) and trust anchors.
func (o ClusterTrustBundleOutput) Spec() ClusterTrustBundleSpecOutput {
	return o.ApplyT(func(v *ClusterTrustBundle) ClusterTrustBundleSpecOutput { return v.Spec }).(ClusterTrustBundleSpecOutput)
}

type ClusterTrustBundleArrayOutput struct{ *pulumi.OutputState }

func (ClusterTrustBundleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterTrustBundle)(nil)).Elem()
}

func (o ClusterTrustBundleArrayOutput) ToClusterTrustBundleArrayOutput() ClusterTrustBundleArrayOutput {
	return o
}

func (o ClusterTrustBundleArrayOutput) ToClusterTrustBundleArrayOutputWithContext(ctx context.Context) ClusterTrustBundleArrayOutput {
	return o
}

func (o ClusterTrustBundleArrayOutput) Index(i pulumi.IntInput) ClusterTrustBundleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterTrustBundle {
		return vs[0].([]*ClusterTrustBundle)[vs[1].(int)]
	}).(ClusterTrustBundleOutput)
}

type ClusterTrustBundleMapOutput struct{ *pulumi.OutputState }

func (ClusterTrustBundleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterTrustBundle)(nil)).Elem()
}

func (o ClusterTrustBundleMapOutput) ToClusterTrustBundleMapOutput() ClusterTrustBundleMapOutput {
	return o
}

func (o ClusterTrustBundleMapOutput) ToClusterTrustBundleMapOutputWithContext(ctx context.Context) ClusterTrustBundleMapOutput {
	return o
}

func (o ClusterTrustBundleMapOutput) MapIndex(k pulumi.StringInput) ClusterTrustBundleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterTrustBundle {
		return vs[0].(map[string]*ClusterTrustBundle)[vs[1].(string)]
	}).(ClusterTrustBundleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTrustBundleInput)(nil)).Elem(), &ClusterTrustBundle{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTrustBundleArrayInput)(nil)).Elem(), ClusterTrustBundleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTrustBundleMapInput)(nil)).Elem(), ClusterTrustBundleMap{})
	pulumi.RegisterOutputType(ClusterTrustBundleOutput{})
	pulumi.RegisterOutputType(ClusterTrustBundleArrayOutput{})
	pulumi.RegisterOutputType(ClusterTrustBundleMapOutput{})
}
