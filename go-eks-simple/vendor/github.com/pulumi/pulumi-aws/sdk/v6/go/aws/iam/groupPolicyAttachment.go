// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attaches a Managed IAM Policy to an IAM group
//
// > **NOTE:** The usage of this resource conflicts with the `iam.PolicyAttachment` resource and will permanently show a difference if both are defined.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			group, err := iam.NewGroup(ctx, "group", &iam.GroupArgs{
//				Name: pulumi.String("test-group"),
//			})
//			if err != nil {
//				return err
//			}
//			policy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
//				Name:        pulumi.String("test-policy"),
//				Description: pulumi.String("A test policy"),
//				Policy:      pulumi.Any("{ ... policy JSON ... }"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewGroupPolicyAttachment(ctx, "test-attach", &iam.GroupPolicyAttachmentArgs{
//				Group:     group.Name,
//				PolicyArn: policy.Arn,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import IAM group policy attachments using the group name and policy arn separated by `/`. For example:
//
// ```sh
// $ pulumi import aws:iam/groupPolicyAttachment:GroupPolicyAttachment test-attach test-group/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
// ```
type GroupPolicyAttachment struct {
	pulumi.CustomResourceState

	// The group the policy should be applied to
	Group pulumi.StringOutput `pulumi:"group"`
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringOutput `pulumi:"policyArn"`
}

// NewGroupPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewGroupPolicyAttachment(ctx *pulumi.Context,
	name string, args *GroupPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*GroupPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.PolicyArn == nil {
		return nil, errors.New("invalid value for required argument 'PolicyArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupPolicyAttachment
	err := ctx.RegisterResource("aws:iam/groupPolicyAttachment:GroupPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupPolicyAttachment gets an existing GroupPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupPolicyAttachmentState, opts ...pulumi.ResourceOption) (*GroupPolicyAttachment, error) {
	var resource GroupPolicyAttachment
	err := ctx.ReadResource("aws:iam/groupPolicyAttachment:GroupPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupPolicyAttachment resources.
type groupPolicyAttachmentState struct {
	// The group the policy should be applied to
	Group interface{} `pulumi:"group"`
	// The ARN of the policy you want to apply
	PolicyArn *string `pulumi:"policyArn"`
}

type GroupPolicyAttachmentState struct {
	// The group the policy should be applied to
	Group pulumi.Input
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringPtrInput
}

func (GroupPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyAttachmentState)(nil)).Elem()
}

type groupPolicyAttachmentArgs struct {
	// The group the policy should be applied to
	Group interface{} `pulumi:"group"`
	// The ARN of the policy you want to apply
	PolicyArn string `pulumi:"policyArn"`
}

// The set of arguments for constructing a GroupPolicyAttachment resource.
type GroupPolicyAttachmentArgs struct {
	// The group the policy should be applied to
	Group pulumi.Input
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringInput
}

func (GroupPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyAttachmentArgs)(nil)).Elem()
}

type GroupPolicyAttachmentInput interface {
	pulumi.Input

	ToGroupPolicyAttachmentOutput() GroupPolicyAttachmentOutput
	ToGroupPolicyAttachmentOutputWithContext(ctx context.Context) GroupPolicyAttachmentOutput
}

func (*GroupPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicyAttachment)(nil)).Elem()
}

func (i *GroupPolicyAttachment) ToGroupPolicyAttachmentOutput() GroupPolicyAttachmentOutput {
	return i.ToGroupPolicyAttachmentOutputWithContext(context.Background())
}

func (i *GroupPolicyAttachment) ToGroupPolicyAttachmentOutputWithContext(ctx context.Context) GroupPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyAttachmentOutput)
}

// GroupPolicyAttachmentArrayInput is an input type that accepts GroupPolicyAttachmentArray and GroupPolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `GroupPolicyAttachmentArrayInput` via:
//
//	GroupPolicyAttachmentArray{ GroupPolicyAttachmentArgs{...} }
type GroupPolicyAttachmentArrayInput interface {
	pulumi.Input

	ToGroupPolicyAttachmentArrayOutput() GroupPolicyAttachmentArrayOutput
	ToGroupPolicyAttachmentArrayOutputWithContext(context.Context) GroupPolicyAttachmentArrayOutput
}

type GroupPolicyAttachmentArray []GroupPolicyAttachmentInput

func (GroupPolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupPolicyAttachment)(nil)).Elem()
}

func (i GroupPolicyAttachmentArray) ToGroupPolicyAttachmentArrayOutput() GroupPolicyAttachmentArrayOutput {
	return i.ToGroupPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i GroupPolicyAttachmentArray) ToGroupPolicyAttachmentArrayOutputWithContext(ctx context.Context) GroupPolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyAttachmentArrayOutput)
}

// GroupPolicyAttachmentMapInput is an input type that accepts GroupPolicyAttachmentMap and GroupPolicyAttachmentMapOutput values.
// You can construct a concrete instance of `GroupPolicyAttachmentMapInput` via:
//
//	GroupPolicyAttachmentMap{ "key": GroupPolicyAttachmentArgs{...} }
type GroupPolicyAttachmentMapInput interface {
	pulumi.Input

	ToGroupPolicyAttachmentMapOutput() GroupPolicyAttachmentMapOutput
	ToGroupPolicyAttachmentMapOutputWithContext(context.Context) GroupPolicyAttachmentMapOutput
}

type GroupPolicyAttachmentMap map[string]GroupPolicyAttachmentInput

func (GroupPolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupPolicyAttachment)(nil)).Elem()
}

func (i GroupPolicyAttachmentMap) ToGroupPolicyAttachmentMapOutput() GroupPolicyAttachmentMapOutput {
	return i.ToGroupPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i GroupPolicyAttachmentMap) ToGroupPolicyAttachmentMapOutputWithContext(ctx context.Context) GroupPolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPolicyAttachmentMapOutput)
}

type GroupPolicyAttachmentOutput struct{ *pulumi.OutputState }

func (GroupPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicyAttachment)(nil)).Elem()
}

func (o GroupPolicyAttachmentOutput) ToGroupPolicyAttachmentOutput() GroupPolicyAttachmentOutput {
	return o
}

func (o GroupPolicyAttachmentOutput) ToGroupPolicyAttachmentOutputWithContext(ctx context.Context) GroupPolicyAttachmentOutput {
	return o
}

// The group the policy should be applied to
func (o GroupPolicyAttachmentOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupPolicyAttachment) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// The ARN of the policy you want to apply
func (o GroupPolicyAttachmentOutput) PolicyArn() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupPolicyAttachment) pulumi.StringOutput { return v.PolicyArn }).(pulumi.StringOutput)
}

type GroupPolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (GroupPolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupPolicyAttachment)(nil)).Elem()
}

func (o GroupPolicyAttachmentArrayOutput) ToGroupPolicyAttachmentArrayOutput() GroupPolicyAttachmentArrayOutput {
	return o
}

func (o GroupPolicyAttachmentArrayOutput) ToGroupPolicyAttachmentArrayOutputWithContext(ctx context.Context) GroupPolicyAttachmentArrayOutput {
	return o
}

func (o GroupPolicyAttachmentArrayOutput) Index(i pulumi.IntInput) GroupPolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupPolicyAttachment {
		return vs[0].([]*GroupPolicyAttachment)[vs[1].(int)]
	}).(GroupPolicyAttachmentOutput)
}

type GroupPolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (GroupPolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupPolicyAttachment)(nil)).Elem()
}

func (o GroupPolicyAttachmentMapOutput) ToGroupPolicyAttachmentMapOutput() GroupPolicyAttachmentMapOutput {
	return o
}

func (o GroupPolicyAttachmentMapOutput) ToGroupPolicyAttachmentMapOutputWithContext(ctx context.Context) GroupPolicyAttachmentMapOutput {
	return o
}

func (o GroupPolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) GroupPolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupPolicyAttachment {
		return vs[0].(map[string]*GroupPolicyAttachment)[vs[1].(string)]
	}).(GroupPolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPolicyAttachmentInput)(nil)).Elem(), &GroupPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPolicyAttachmentArrayInput)(nil)).Elem(), GroupPolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPolicyAttachmentMapInput)(nil)).Elem(), GroupPolicyAttachmentMap{})
	pulumi.RegisterOutputType(GroupPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(GroupPolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(GroupPolicyAttachmentMapOutput{})
}
