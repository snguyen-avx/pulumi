// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can be useful for getting back a list of managed prefix list ids to be referenced elsewhere.
func GetManagedPrefixLists(ctx *pulumi.Context, args *GetManagedPrefixListsArgs, opts ...pulumi.InvokeOption) (*GetManagedPrefixListsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetManagedPrefixListsResult
	err := ctx.Invoke("aws:ec2/getManagedPrefixLists:getManagedPrefixLists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getManagedPrefixLists.
type GetManagedPrefixListsArgs struct {
	// Custom filter block as described below.
	Filters []GetManagedPrefixListsFilter `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired .
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getManagedPrefixLists.
type GetManagedPrefixListsResult struct {
	Filters []GetManagedPrefixListsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of all the managed prefix list ids found.
	Ids  []string          `pulumi:"ids"`
	Tags map[string]string `pulumi:"tags"`
}

func GetManagedPrefixListsOutput(ctx *pulumi.Context, args GetManagedPrefixListsOutputArgs, opts ...pulumi.InvokeOption) GetManagedPrefixListsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetManagedPrefixListsResult, error) {
			args := v.(GetManagedPrefixListsArgs)
			r, err := GetManagedPrefixLists(ctx, &args, opts...)
			var s GetManagedPrefixListsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetManagedPrefixListsResultOutput)
}

// A collection of arguments for invoking getManagedPrefixLists.
type GetManagedPrefixListsOutputArgs struct {
	// Custom filter block as described below.
	Filters GetManagedPrefixListsFilterArrayInput `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired .
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetManagedPrefixListsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetManagedPrefixListsArgs)(nil)).Elem()
}

// A collection of values returned by getManagedPrefixLists.
type GetManagedPrefixListsResultOutput struct{ *pulumi.OutputState }

func (GetManagedPrefixListsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetManagedPrefixListsResult)(nil)).Elem()
}

func (o GetManagedPrefixListsResultOutput) ToGetManagedPrefixListsResultOutput() GetManagedPrefixListsResultOutput {
	return o
}

func (o GetManagedPrefixListsResultOutput) ToGetManagedPrefixListsResultOutputWithContext(ctx context.Context) GetManagedPrefixListsResultOutput {
	return o
}

func (o GetManagedPrefixListsResultOutput) Filters() GetManagedPrefixListsFilterArrayOutput {
	return o.ApplyT(func(v GetManagedPrefixListsResult) []GetManagedPrefixListsFilter { return v.Filters }).(GetManagedPrefixListsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetManagedPrefixListsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetManagedPrefixListsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of all the managed prefix list ids found.
func (o GetManagedPrefixListsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetManagedPrefixListsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetManagedPrefixListsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetManagedPrefixListsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetManagedPrefixListsResultOutput{})
}
