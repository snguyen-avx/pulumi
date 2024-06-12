// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Network Insights Path resource. Part of the "Reachability Analyzer" service in the AWS VPC console.
//
// ## Example Usage
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
//			_, err := ec2.NewNetworkInsightsPath(ctx, "test", &ec2.NetworkInsightsPathArgs{
//				Source:      pulumi.Any(source.Id),
//				Destination: pulumi.Any(destination.Id),
//				Protocol:    pulumi.String("tcp"),
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
// Using `pulumi import`, import Network Insights Paths using the `id`. For example:
//
// ```sh
// $ pulumi import aws:ec2/networkInsightsPath:NetworkInsightsPath test nip-00edfba169923aefd
// ```
type NetworkInsightsPath struct {
	pulumi.CustomResourceState

	// ARN of the Network Insights Path.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// ID or ARN of the resource which is the destination of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// ARN of the destination.
	DestinationArn pulumi.StringOutput `pulumi:"destinationArn"`
	// IP address of the destination resource.
	DestinationIp pulumi.StringPtrOutput `pulumi:"destinationIp"`
	// Destination port to analyze access to.
	DestinationPort pulumi.IntPtrOutput `pulumi:"destinationPort"`
	// Protocol to use for analysis. Valid options are `tcp` or `udp`.
	//
	// The following arguments are optional:
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// ID or ARN of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Source pulumi.StringOutput `pulumi:"source"`
	// ARN of the source.
	SourceArn pulumi.StringOutput `pulumi:"sourceArn"`
	// IP address of the source resource.
	SourceIp pulumi.StringPtrOutput `pulumi:"sourceIp"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewNetworkInsightsPath registers a new resource with the given unique name, arguments, and options.
func NewNetworkInsightsPath(ctx *pulumi.Context,
	name string, args *NetworkInsightsPathArgs, opts ...pulumi.ResourceOption) (*NetworkInsightsPath, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkInsightsPath
	err := ctx.RegisterResource("aws:ec2/networkInsightsPath:NetworkInsightsPath", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInsightsPath gets an existing NetworkInsightsPath resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInsightsPath(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInsightsPathState, opts ...pulumi.ResourceOption) (*NetworkInsightsPath, error) {
	var resource NetworkInsightsPath
	err := ctx.ReadResource("aws:ec2/networkInsightsPath:NetworkInsightsPath", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInsightsPath resources.
type networkInsightsPathState struct {
	// ARN of the Network Insights Path.
	Arn *string `pulumi:"arn"`
	// ID or ARN of the resource which is the destination of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Destination *string `pulumi:"destination"`
	// ARN of the destination.
	DestinationArn *string `pulumi:"destinationArn"`
	// IP address of the destination resource.
	DestinationIp *string `pulumi:"destinationIp"`
	// Destination port to analyze access to.
	DestinationPort *int `pulumi:"destinationPort"`
	// Protocol to use for analysis. Valid options are `tcp` or `udp`.
	//
	// The following arguments are optional:
	Protocol *string `pulumi:"protocol"`
	// ID or ARN of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Source *string `pulumi:"source"`
	// ARN of the source.
	SourceArn *string `pulumi:"sourceArn"`
	// IP address of the source resource.
	SourceIp *string `pulumi:"sourceIp"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type NetworkInsightsPathState struct {
	// ARN of the Network Insights Path.
	Arn pulumi.StringPtrInput
	// ID or ARN of the resource which is the destination of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Destination pulumi.StringPtrInput
	// ARN of the destination.
	DestinationArn pulumi.StringPtrInput
	// IP address of the destination resource.
	DestinationIp pulumi.StringPtrInput
	// Destination port to analyze access to.
	DestinationPort pulumi.IntPtrInput
	// Protocol to use for analysis. Valid options are `tcp` or `udp`.
	//
	// The following arguments are optional:
	Protocol pulumi.StringPtrInput
	// ID or ARN of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Source pulumi.StringPtrInput
	// ARN of the source.
	SourceArn pulumi.StringPtrInput
	// IP address of the source resource.
	SourceIp pulumi.StringPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (NetworkInsightsPathState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInsightsPathState)(nil)).Elem()
}

type networkInsightsPathArgs struct {
	// ID or ARN of the resource which is the destination of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Destination string `pulumi:"destination"`
	// IP address of the destination resource.
	DestinationIp *string `pulumi:"destinationIp"`
	// Destination port to analyze access to.
	DestinationPort *int `pulumi:"destinationPort"`
	// Protocol to use for analysis. Valid options are `tcp` or `udp`.
	//
	// The following arguments are optional:
	Protocol string `pulumi:"protocol"`
	// ID or ARN of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Source string `pulumi:"source"`
	// IP address of the source resource.
	SourceIp *string `pulumi:"sourceIp"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInsightsPath resource.
type NetworkInsightsPathArgs struct {
	// ID or ARN of the resource which is the destination of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Destination pulumi.StringInput
	// IP address of the destination resource.
	DestinationIp pulumi.StringPtrInput
	// Destination port to analyze access to.
	DestinationPort pulumi.IntPtrInput
	// Protocol to use for analysis. Valid options are `tcp` or `udp`.
	//
	// The following arguments are optional:
	Protocol pulumi.StringInput
	// ID or ARN of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
	Source pulumi.StringInput
	// IP address of the source resource.
	SourceIp pulumi.StringPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (NetworkInsightsPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInsightsPathArgs)(nil)).Elem()
}

type NetworkInsightsPathInput interface {
	pulumi.Input

	ToNetworkInsightsPathOutput() NetworkInsightsPathOutput
	ToNetworkInsightsPathOutputWithContext(ctx context.Context) NetworkInsightsPathOutput
}

func (*NetworkInsightsPath) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInsightsPath)(nil)).Elem()
}

func (i *NetworkInsightsPath) ToNetworkInsightsPathOutput() NetworkInsightsPathOutput {
	return i.ToNetworkInsightsPathOutputWithContext(context.Background())
}

func (i *NetworkInsightsPath) ToNetworkInsightsPathOutputWithContext(ctx context.Context) NetworkInsightsPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInsightsPathOutput)
}

// NetworkInsightsPathArrayInput is an input type that accepts NetworkInsightsPathArray and NetworkInsightsPathArrayOutput values.
// You can construct a concrete instance of `NetworkInsightsPathArrayInput` via:
//
//	NetworkInsightsPathArray{ NetworkInsightsPathArgs{...} }
type NetworkInsightsPathArrayInput interface {
	pulumi.Input

	ToNetworkInsightsPathArrayOutput() NetworkInsightsPathArrayOutput
	ToNetworkInsightsPathArrayOutputWithContext(context.Context) NetworkInsightsPathArrayOutput
}

type NetworkInsightsPathArray []NetworkInsightsPathInput

func (NetworkInsightsPathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkInsightsPath)(nil)).Elem()
}

func (i NetworkInsightsPathArray) ToNetworkInsightsPathArrayOutput() NetworkInsightsPathArrayOutput {
	return i.ToNetworkInsightsPathArrayOutputWithContext(context.Background())
}

func (i NetworkInsightsPathArray) ToNetworkInsightsPathArrayOutputWithContext(ctx context.Context) NetworkInsightsPathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInsightsPathArrayOutput)
}

// NetworkInsightsPathMapInput is an input type that accepts NetworkInsightsPathMap and NetworkInsightsPathMapOutput values.
// You can construct a concrete instance of `NetworkInsightsPathMapInput` via:
//
//	NetworkInsightsPathMap{ "key": NetworkInsightsPathArgs{...} }
type NetworkInsightsPathMapInput interface {
	pulumi.Input

	ToNetworkInsightsPathMapOutput() NetworkInsightsPathMapOutput
	ToNetworkInsightsPathMapOutputWithContext(context.Context) NetworkInsightsPathMapOutput
}

type NetworkInsightsPathMap map[string]NetworkInsightsPathInput

func (NetworkInsightsPathMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkInsightsPath)(nil)).Elem()
}

func (i NetworkInsightsPathMap) ToNetworkInsightsPathMapOutput() NetworkInsightsPathMapOutput {
	return i.ToNetworkInsightsPathMapOutputWithContext(context.Background())
}

func (i NetworkInsightsPathMap) ToNetworkInsightsPathMapOutputWithContext(ctx context.Context) NetworkInsightsPathMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInsightsPathMapOutput)
}

type NetworkInsightsPathOutput struct{ *pulumi.OutputState }

func (NetworkInsightsPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInsightsPath)(nil)).Elem()
}

func (o NetworkInsightsPathOutput) ToNetworkInsightsPathOutput() NetworkInsightsPathOutput {
	return o
}

func (o NetworkInsightsPathOutput) ToNetworkInsightsPathOutputWithContext(ctx context.Context) NetworkInsightsPathOutput {
	return o
}

// ARN of the Network Insights Path.
func (o NetworkInsightsPathOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// ID or ARN of the resource which is the destination of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
func (o NetworkInsightsPathOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// ARN of the destination.
func (o NetworkInsightsPathOutput) DestinationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.DestinationArn }).(pulumi.StringOutput)
}

// IP address of the destination resource.
func (o NetworkInsightsPathOutput) DestinationIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringPtrOutput { return v.DestinationIp }).(pulumi.StringPtrOutput)
}

// Destination port to analyze access to.
func (o NetworkInsightsPathOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.IntPtrOutput { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

// Protocol to use for analysis. Valid options are `tcp` or `udp`.
//
// The following arguments are optional:
func (o NetworkInsightsPathOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// ID or ARN of the resource which is the source of the path. Can be an Instance, Internet Gateway, Network Interface, Transit Gateway, VPC Endpoint, VPC Peering Connection or VPN Gateway. If the resource is in another account, you must specify an ARN.
func (o NetworkInsightsPathOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// ARN of the source.
func (o NetworkInsightsPathOutput) SourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringOutput { return v.SourceArn }).(pulumi.StringOutput)
}

// IP address of the source resource.
func (o NetworkInsightsPathOutput) SourceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringPtrOutput { return v.SourceIp }).(pulumi.StringPtrOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o NetworkInsightsPathOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o NetworkInsightsPathOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkInsightsPath) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type NetworkInsightsPathArrayOutput struct{ *pulumi.OutputState }

func (NetworkInsightsPathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkInsightsPath)(nil)).Elem()
}

func (o NetworkInsightsPathArrayOutput) ToNetworkInsightsPathArrayOutput() NetworkInsightsPathArrayOutput {
	return o
}

func (o NetworkInsightsPathArrayOutput) ToNetworkInsightsPathArrayOutputWithContext(ctx context.Context) NetworkInsightsPathArrayOutput {
	return o
}

func (o NetworkInsightsPathArrayOutput) Index(i pulumi.IntInput) NetworkInsightsPathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkInsightsPath {
		return vs[0].([]*NetworkInsightsPath)[vs[1].(int)]
	}).(NetworkInsightsPathOutput)
}

type NetworkInsightsPathMapOutput struct{ *pulumi.OutputState }

func (NetworkInsightsPathMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkInsightsPath)(nil)).Elem()
}

func (o NetworkInsightsPathMapOutput) ToNetworkInsightsPathMapOutput() NetworkInsightsPathMapOutput {
	return o
}

func (o NetworkInsightsPathMapOutput) ToNetworkInsightsPathMapOutputWithContext(ctx context.Context) NetworkInsightsPathMapOutput {
	return o
}

func (o NetworkInsightsPathMapOutput) MapIndex(k pulumi.StringInput) NetworkInsightsPathOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkInsightsPath {
		return vs[0].(map[string]*NetworkInsightsPath)[vs[1].(string)]
	}).(NetworkInsightsPathOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInsightsPathInput)(nil)).Elem(), &NetworkInsightsPath{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInsightsPathArrayInput)(nil)).Elem(), NetworkInsightsPathArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInsightsPathMapInput)(nil)).Elem(), NetworkInsightsPathMap{})
	pulumi.RegisterOutputType(NetworkInsightsPathOutput{})
	pulumi.RegisterOutputType(NetworkInsightsPathArrayOutput{})
	pulumi.RegisterOutputType(NetworkInsightsPathMapOutput{})
}
