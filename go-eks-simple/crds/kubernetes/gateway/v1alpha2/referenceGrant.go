// Code generated by crd2pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha2

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ReferenceGrant identifies kinds of resources in other namespaces that are trusted to reference the specified kinds of resources in the same namespace as the policy.
//
//	Each ReferenceGrant can be used to represent a unique trust relationship. Additional Reference Grants can be used to add to the set of trusted sources of inbound references for the namespace they are defined within.
//	A ReferenceGrant is required for all cross-namespace references in Gateway API (with the exception of cross-namespace Route-Gateway attachment, which is governed by the AllowedRoutes configuration on the Gateway, and cross-namespace Service ParentRefs on a "consumer" mesh Route, which defines routing rules applicable only to workloads in the Route namespace). ReferenceGrants allowing a reference from a Route to a Service are only applicable to BackendRefs.
//	ReferenceGrant is a form of runtime verification allowing users to assert which cross-namespace object references are permitted. Implementations that support ReferenceGrant MUST NOT permit cross-namespace references which have no grant, and MUST respond to the removal of a grant by revoking the access that the grant allowed.
type ReferenceGrant struct {
	pulumi.CustomResourceState

	ApiVersion pulumi.StringPtrOutput     `pulumi:"apiVersion"`
	Kind       pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata   metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec defines the desired state of ReferenceGrant.
	Spec ReferenceGrantSpecPtrOutput `pulumi:"spec"`
}

// NewReferenceGrant registers a new resource with the given unique name, arguments, and options.
func NewReferenceGrant(ctx *pulumi.Context,
	name string, args *ReferenceGrantArgs, opts ...pulumi.ResourceOption) (*ReferenceGrant, error) {
	if args == nil {
		args = &ReferenceGrantArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("gateway.networking.k8s.io/v1alpha2")
	args.Kind = pulumi.StringPtr("ReferenceGrant")
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReferenceGrant
	err := ctx.RegisterResource("kubernetes:gateway.networking.k8s.io/v1alpha2:ReferenceGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReferenceGrant gets an existing ReferenceGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReferenceGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReferenceGrantState, opts ...pulumi.ResourceOption) (*ReferenceGrant, error) {
	var resource ReferenceGrant
	err := ctx.ReadResource("kubernetes:gateway.networking.k8s.io/v1alpha2:ReferenceGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReferenceGrant resources.
type referenceGrantState struct {
}

type ReferenceGrantState struct {
}

func (ReferenceGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceGrantState)(nil)).Elem()
}

type referenceGrantArgs struct {
	ApiVersion *string            `pulumi:"apiVersion"`
	Kind       *string            `pulumi:"kind"`
	Metadata   *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec defines the desired state of ReferenceGrant.
	Spec *ReferenceGrantSpec `pulumi:"spec"`
}

// The set of arguments for constructing a ReferenceGrant resource.
type ReferenceGrantArgs struct {
	ApiVersion pulumi.StringPtrInput
	Kind       pulumi.StringPtrInput
	Metadata   metav1.ObjectMetaPtrInput
	// Spec defines the desired state of ReferenceGrant.
	Spec ReferenceGrantSpecPtrInput
}

func (ReferenceGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceGrantArgs)(nil)).Elem()
}

type ReferenceGrantInput interface {
	pulumi.Input

	ToReferenceGrantOutput() ReferenceGrantOutput
	ToReferenceGrantOutputWithContext(ctx context.Context) ReferenceGrantOutput
}

func (*ReferenceGrant) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceGrant)(nil)).Elem()
}

func (i *ReferenceGrant) ToReferenceGrantOutput() ReferenceGrantOutput {
	return i.ToReferenceGrantOutputWithContext(context.Background())
}

func (i *ReferenceGrant) ToReferenceGrantOutputWithContext(ctx context.Context) ReferenceGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceGrantOutput)
}

type ReferenceGrantOutput struct{ *pulumi.OutputState }

func (ReferenceGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceGrant)(nil)).Elem()
}

func (o ReferenceGrantOutput) ToReferenceGrantOutput() ReferenceGrantOutput {
	return o
}

func (o ReferenceGrantOutput) ToReferenceGrantOutputWithContext(ctx context.Context) ReferenceGrantOutput {
	return o
}

func (o ReferenceGrantOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReferenceGrant) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o ReferenceGrantOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReferenceGrant) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ReferenceGrantOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *ReferenceGrant) metav1.ObjectMetaPtrOutput { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Spec defines the desired state of ReferenceGrant.
func (o ReferenceGrantOutput) Spec() ReferenceGrantSpecPtrOutput {
	return o.ApplyT(func(v *ReferenceGrant) ReferenceGrantSpecPtrOutput { return v.Spec }).(ReferenceGrantSpecPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceGrantInput)(nil)).Elem(), &ReferenceGrant{})
	pulumi.RegisterOutputType(ReferenceGrantOutput{})
}
