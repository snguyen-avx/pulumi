// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the Account ID of the [AWS Billing and Cost Management Service Account](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-getting-started.html#step-2) for the purpose of permitting in S3 bucket policy.
func GetBillingServiceAccount(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetBillingServiceAccountResult, error) {
	var rv GetBillingServiceAccountResult
	err := ctx.Invoke("aws:index/getBillingServiceAccount:getBillingServiceAccount", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getBillingServiceAccount.
type GetBillingServiceAccountResult struct {
	// ARN of the AWS billing service account.
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}