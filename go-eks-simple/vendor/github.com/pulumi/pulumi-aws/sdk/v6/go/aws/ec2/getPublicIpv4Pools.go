// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data source for getting information about AWS EC2 Public IPv4 Pools.
//
// ## Example Usage
//
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Returns all public IPv4 pools.
//			_, err := ec2.GetPublicIpv4Pools(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### Usage with Filter
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.GetPublicIpv4Pools(ctx, &ec2.GetPublicIpv4PoolsArgs{
//				Filters: []ec2.GetPublicIpv4PoolsFilter{
//					{
//						Name: "tag-key",
//						Values: []string{
//							"ExampleTagKey",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPublicIpv4Pools(ctx *pulumi.Context, args *GetPublicIpv4PoolsArgs, opts ...pulumi.InvokeOption) (*GetPublicIpv4PoolsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPublicIpv4PoolsResult
	err := ctx.Invoke("aws:ec2/getPublicIpv4Pools:getPublicIpv4Pools", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicIpv4Pools.
type GetPublicIpv4PoolsArgs struct {
	// Custom filter block as described below.
	Filters []GetPublicIpv4PoolsFilter `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match a pair on the desired pools.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getPublicIpv4Pools.
type GetPublicIpv4PoolsResult struct {
	Filters []GetPublicIpv4PoolsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of all the pool IDs found.
	PoolIds []string          `pulumi:"poolIds"`
	Tags    map[string]string `pulumi:"tags"`
}

func GetPublicIpv4PoolsOutput(ctx *pulumi.Context, args GetPublicIpv4PoolsOutputArgs, opts ...pulumi.InvokeOption) GetPublicIpv4PoolsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPublicIpv4PoolsResult, error) {
			args := v.(GetPublicIpv4PoolsArgs)
			r, err := GetPublicIpv4Pools(ctx, &args, opts...)
			var s GetPublicIpv4PoolsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPublicIpv4PoolsResultOutput)
}

// A collection of arguments for invoking getPublicIpv4Pools.
type GetPublicIpv4PoolsOutputArgs struct {
	// Custom filter block as described below.
	Filters GetPublicIpv4PoolsFilterArrayInput `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match a pair on the desired pools.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetPublicIpv4PoolsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicIpv4PoolsArgs)(nil)).Elem()
}

// A collection of values returned by getPublicIpv4Pools.
type GetPublicIpv4PoolsResultOutput struct{ *pulumi.OutputState }

func (GetPublicIpv4PoolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicIpv4PoolsResult)(nil)).Elem()
}

func (o GetPublicIpv4PoolsResultOutput) ToGetPublicIpv4PoolsResultOutput() GetPublicIpv4PoolsResultOutput {
	return o
}

func (o GetPublicIpv4PoolsResultOutput) ToGetPublicIpv4PoolsResultOutputWithContext(ctx context.Context) GetPublicIpv4PoolsResultOutput {
	return o
}

func (o GetPublicIpv4PoolsResultOutput) Filters() GetPublicIpv4PoolsFilterArrayOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolsResult) []GetPublicIpv4PoolsFilter { return v.Filters }).(GetPublicIpv4PoolsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPublicIpv4PoolsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of all the pool IDs found.
func (o GetPublicIpv4PoolsResultOutput) PoolIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolsResult) []string { return v.PoolIds }).(pulumi.StringArrayOutput)
}

func (o GetPublicIpv4PoolsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPublicIpv4PoolsResultOutput{})
}
