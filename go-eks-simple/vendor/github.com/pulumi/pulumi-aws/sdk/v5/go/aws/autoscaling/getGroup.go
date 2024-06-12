// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an existing autoscaling group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/autoscaling"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := autoscaling.LookupGroup(ctx, &autoscaling.LookupGroupArgs{
//				Name: "foo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("aws:autoscaling/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// Specify the exact name of the desired autoscaling group.
	Name string `pulumi:"name"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// ARN of the Auto Scaling group.
	Arn string `pulumi:"arn"`
	// One or more Availability Zones for the group.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	DefaultCooldown   int      `pulumi:"defaultCooldown"`
	// Desired size of the group.
	DesiredCapacity int `pulumi:"desiredCapacity"`
	// The unit of measurement for the value returned for `desiredCapacity`.
	DesiredCapacityType string `pulumi:"desiredCapacityType"`
	// List of metrics enabled for collection.
	EnabledMetrics []string `pulumi:"enabledMetrics"`
	// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
	HealthCheckGracePeriod int `pulumi:"healthCheckGracePeriod"`
	// Service to use for the health checks. The valid values are EC2 and ELB.
	HealthCheckType string `pulumi:"healthCheckType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the associated launch configuration.
	LaunchConfiguration string                   `pulumi:"launchConfiguration"`
	LaunchTemplates     []GetGroupLaunchTemplate `pulumi:"launchTemplates"`
	// One or more load balancers associated with the group.
	LoadBalancers []string `pulumi:"loadBalancers"`
	// Maximum size of the group.
	MaxSize int `pulumi:"maxSize"`
	// Minimum size of the group.
	MinSize int `pulumi:"minSize"`
	// Name of the Auto Scaling Group.
	Name                             string `pulumi:"name"`
	NewInstancesProtectedFromScaleIn bool   `pulumi:"newInstancesProtectedFromScaleIn"`
	// Name of the placement group into which to launch your instances, if any. For more information, see Placement Groups (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the Amazon Elastic Compute Cloud User Guide.
	PlacementGroup string `pulumi:"placementGroup"`
	// ARN of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.
	ServiceLinkedRoleArn string `pulumi:"serviceLinkedRoleArn"`
	// Current state of the group when DeleteAutoScalingGroup is in progress.
	Status string `pulumi:"status"`
	// ARNs of the target groups for your load balancer.
	TargetGroupArns []string `pulumi:"targetGroupArns"`
	// The termination policies for the group.
	TerminationPolicies []string `pulumi:"terminationPolicies"`
	// VPC ID for the group.
	VpcZoneIdentifier string `pulumi:"vpcZoneIdentifier"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

// A collection of arguments for invoking getGroup.
type LookupGroupOutputArgs struct {
	// Specify the exact name of the desired autoscaling group.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

// A collection of values returned by getGroup.
type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

// ARN of the Auto Scaling group.
func (o LookupGroupResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Arn }).(pulumi.StringOutput)
}

// One or more Availability Zones for the group.
func (o LookupGroupResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LookupGroupResultOutput) DefaultCooldown() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.DefaultCooldown }).(pulumi.IntOutput)
}

// Desired size of the group.
func (o LookupGroupResultOutput) DesiredCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.DesiredCapacity }).(pulumi.IntOutput)
}

// The unit of measurement for the value returned for `desiredCapacity`.
func (o LookupGroupResultOutput) DesiredCapacityType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.DesiredCapacityType }).(pulumi.StringOutput)
}

// List of metrics enabled for collection.
func (o LookupGroupResultOutput) EnabledMetrics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.EnabledMetrics }).(pulumi.StringArrayOutput)
}

// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
func (o LookupGroupResultOutput) HealthCheckGracePeriod() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.HealthCheckGracePeriod }).(pulumi.IntOutput)
}

// Service to use for the health checks. The valid values are EC2 and ELB.
func (o LookupGroupResultOutput) HealthCheckType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.HealthCheckType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the associated launch configuration.
func (o LookupGroupResultOutput) LaunchConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.LaunchConfiguration }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) LaunchTemplates() GetGroupLaunchTemplateArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []GetGroupLaunchTemplate { return v.LaunchTemplates }).(GetGroupLaunchTemplateArrayOutput)
}

// One or more load balancers associated with the group.
func (o LookupGroupResultOutput) LoadBalancers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.LoadBalancers }).(pulumi.StringArrayOutput)
}

// Maximum size of the group.
func (o LookupGroupResultOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.MaxSize }).(pulumi.IntOutput)
}

// Minimum size of the group.
func (o LookupGroupResultOutput) MinSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGroupResult) int { return v.MinSize }).(pulumi.IntOutput)
}

// Name of the Auto Scaling Group.
func (o LookupGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGroupResultOutput) NewInstancesProtectedFromScaleIn() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.NewInstancesProtectedFromScaleIn }).(pulumi.BoolOutput)
}

// Name of the placement group into which to launch your instances, if any. For more information, see Placement Groups (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the Amazon Elastic Compute Cloud User Guide.
func (o LookupGroupResultOutput) PlacementGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.PlacementGroup }).(pulumi.StringOutput)
}

// ARN of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.
func (o LookupGroupResultOutput) ServiceLinkedRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.ServiceLinkedRoleArn }).(pulumi.StringOutput)
}

// Current state of the group when DeleteAutoScalingGroup is in progress.
func (o LookupGroupResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Status }).(pulumi.StringOutput)
}

// ARNs of the target groups for your load balancer.
func (o LookupGroupResultOutput) TargetGroupArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.TargetGroupArns }).(pulumi.StringArrayOutput)
}

// The termination policies for the group.
func (o LookupGroupResultOutput) TerminationPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.TerminationPolicies }).(pulumi.StringArrayOutput)
}

// VPC ID for the group.
func (o LookupGroupResultOutput) VpcZoneIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.VpcZoneIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}