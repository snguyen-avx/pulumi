// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `route53.DelegationSet` provides details about a specific Route 53 Delegation Set.
//
// This data source allows to find a list of name servers associated with a specific delegation set.
//
// ## Example Usage
//
// The following example shows how to get a delegation set from its id.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := route53.LookupDelegationSet(ctx, &route53.LookupDelegationSetArgs{
//				Id: "MQWGHCBFAKEID",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDelegationSet(ctx *pulumi.Context, args *LookupDelegationSetArgs, opts ...pulumi.InvokeOption) (*LookupDelegationSetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDelegationSetResult
	err := ctx.Invoke("aws:route53/getDelegationSet:getDelegationSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDelegationSet.
type LookupDelegationSetArgs struct {
	// Delegation set ID.
	//
	// The following attribute is additionally exported:
	Id string `pulumi:"id"`
}

// A collection of values returned by getDelegationSet.
type LookupDelegationSetResult struct {
	Arn             string   `pulumi:"arn"`
	CallerReference string   `pulumi:"callerReference"`
	Id              string   `pulumi:"id"`
	NameServers     []string `pulumi:"nameServers"`
}

func LookupDelegationSetOutput(ctx *pulumi.Context, args LookupDelegationSetOutputArgs, opts ...pulumi.InvokeOption) LookupDelegationSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDelegationSetResult, error) {
			args := v.(LookupDelegationSetArgs)
			r, err := LookupDelegationSet(ctx, &args, opts...)
			var s LookupDelegationSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDelegationSetResultOutput)
}

// A collection of arguments for invoking getDelegationSet.
type LookupDelegationSetOutputArgs struct {
	// Delegation set ID.
	//
	// The following attribute is additionally exported:
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDelegationSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDelegationSetArgs)(nil)).Elem()
}

// A collection of values returned by getDelegationSet.
type LookupDelegationSetResultOutput struct{ *pulumi.OutputState }

func (LookupDelegationSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDelegationSetResult)(nil)).Elem()
}

func (o LookupDelegationSetResultOutput) ToLookupDelegationSetResultOutput() LookupDelegationSetResultOutput {
	return o
}

func (o LookupDelegationSetResultOutput) ToLookupDelegationSetResultOutputWithContext(ctx context.Context) LookupDelegationSetResultOutput {
	return o
}

func (o LookupDelegationSetResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegationSetResult) string { return v.Arn }).(pulumi.StringOutput)
}

func (o LookupDelegationSetResultOutput) CallerReference() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegationSetResult) string { return v.CallerReference }).(pulumi.StringOutput)
}

func (o LookupDelegationSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDelegationSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDelegationSetResultOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDelegationSetResult) []string { return v.NameServers }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDelegationSetResultOutput{})
}