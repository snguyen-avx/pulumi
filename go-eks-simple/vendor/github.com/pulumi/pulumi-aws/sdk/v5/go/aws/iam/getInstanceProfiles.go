// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch information about all
// IAM instance profiles under a role. By using this data source, you can reference IAM
// instance profile properties without having to hard code ARNs as input.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.GetInstanceProfiles(ctx, &iam.GetInstanceProfilesArgs{
//				RoleName: "an_example_iam_role_name",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstanceProfiles(ctx *pulumi.Context, args *GetInstanceProfilesArgs, opts ...pulumi.InvokeOption) (*GetInstanceProfilesResult, error) {
	var rv GetInstanceProfilesResult
	err := ctx.Invoke("aws:iam/getInstanceProfiles:getInstanceProfiles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceProfiles.
type GetInstanceProfilesArgs struct {
	// IAM role name.
	RoleName string `pulumi:"roleName"`
}

// A collection of values returned by getInstanceProfiles.
type GetInstanceProfilesResult struct {
	// Set of ARNs of instance profiles.
	Arns []string `pulumi:"arns"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set of IAM instance profile names.
	Names []string `pulumi:"names"`
	// Set of IAM instance profile paths.
	Paths    []string `pulumi:"paths"`
	RoleName string   `pulumi:"roleName"`
}

func GetInstanceProfilesOutput(ctx *pulumi.Context, args GetInstanceProfilesOutputArgs, opts ...pulumi.InvokeOption) GetInstanceProfilesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceProfilesResult, error) {
			args := v.(GetInstanceProfilesArgs)
			r, err := GetInstanceProfiles(ctx, &args, opts...)
			var s GetInstanceProfilesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceProfilesResultOutput)
}

// A collection of arguments for invoking getInstanceProfiles.
type GetInstanceProfilesOutputArgs struct {
	// IAM role name.
	RoleName pulumi.StringInput `pulumi:"roleName"`
}

func (GetInstanceProfilesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceProfilesArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceProfiles.
type GetInstanceProfilesResultOutput struct{ *pulumi.OutputState }

func (GetInstanceProfilesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceProfilesResult)(nil)).Elem()
}

func (o GetInstanceProfilesResultOutput) ToGetInstanceProfilesResultOutput() GetInstanceProfilesResultOutput {
	return o
}

func (o GetInstanceProfilesResultOutput) ToGetInstanceProfilesResultOutputWithContext(ctx context.Context) GetInstanceProfilesResultOutput {
	return o
}

// Set of ARNs of instance profiles.
func (o GetInstanceProfilesResultOutput) Arns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceProfilesResult) []string { return v.Arns }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceProfilesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceProfilesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of IAM instance profile names.
func (o GetInstanceProfilesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceProfilesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

// Set of IAM instance profile paths.
func (o GetInstanceProfilesResultOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceProfilesResult) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o GetInstanceProfilesResultOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceProfilesResult) string { return v.RoleName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceProfilesResultOutput{})
}
