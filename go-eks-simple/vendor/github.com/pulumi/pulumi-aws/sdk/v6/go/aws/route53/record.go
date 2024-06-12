// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Route53 record resource.
//
// ## Example Usage
//
// ### Simple routing policy
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
//			_, err := route53.NewRecord(ctx, "www", &route53.RecordArgs{
//				ZoneId: pulumi.Any(primary.ZoneId),
//				Name:   pulumi.String("www.example.com"),
//				Type:   pulumi.String(route53.RecordTypeA),
//				Ttl:    pulumi.Int(300),
//				Records: pulumi.StringArray{
//					lb.PublicIp,
//				},
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
// ### Weighted routing policy
//
// Other routing policies are configured similarly. See [Amazon Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html) for details.
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
//			_, err := route53.NewRecord(ctx, "www-dev", &route53.RecordArgs{
//				ZoneId: pulumi.Any(primary.ZoneId),
//				Name:   pulumi.String("www"),
//				Type:   pulumi.String(route53.RecordTypeCNAME),
//				Ttl:    pulumi.Int(5),
//				WeightedRoutingPolicies: route53.RecordWeightedRoutingPolicyArray{
//					&route53.RecordWeightedRoutingPolicyArgs{
//						Weight: pulumi.Int(10),
//					},
//				},
//				SetIdentifier: pulumi.String("dev"),
//				Records: pulumi.StringArray{
//					pulumi.String("dev.example.com"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = route53.NewRecord(ctx, "www-live", &route53.RecordArgs{
//				ZoneId: pulumi.Any(primary.ZoneId),
//				Name:   pulumi.String("www"),
//				Type:   pulumi.String(route53.RecordTypeCNAME),
//				Ttl:    pulumi.Int(5),
//				WeightedRoutingPolicies: route53.RecordWeightedRoutingPolicyArray{
//					&route53.RecordWeightedRoutingPolicyArgs{
//						Weight: pulumi.Int(90),
//					},
//				},
//				SetIdentifier: pulumi.String("live"),
//				Records: pulumi.StringArray{
//					pulumi.String("live.example.com"),
//				},
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
// ### Geoproximity routing policy
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
//			_, err := route53.NewRecord(ctx, "www", &route53.RecordArgs{
//				ZoneId: pulumi.Any(primary.ZoneId),
//				Name:   pulumi.String("www.example.com"),
//				Type:   pulumi.String(route53.RecordTypeCNAME),
//				Ttl:    pulumi.Int(300),
//				GeoproximityRoutingPolicy: &route53.RecordGeoproximityRoutingPolicyArgs{
//					Coordinates: route53.RecordGeoproximityRoutingPolicyCoordinateArray{
//						&route53.RecordGeoproximityRoutingPolicyCoordinateArgs{
//							Latitude:  pulumi.String("49.22"),
//							Longitude: pulumi.String("-74.01"),
//						},
//					},
//				},
//				SetIdentifier: pulumi.String("dev"),
//				Records: pulumi.StringArray{
//					pulumi.String("dev.example.com"),
//				},
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
// ### Alias record
//
// See [related part of Amazon Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-choosing-alias-non-alias.html)
// to understand differences between alias and non-alias records.
//
// TTL for all alias records is [60 seconds](https://aws.amazon.com/route53/faqs/#dns_failover_do_i_need_to_adjust),
// you cannot change this, therefore `ttl` has to be omitted in alias records.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elb"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := elb.NewLoadBalancer(ctx, "main", &elb.LoadBalancerArgs{
//				Name: pulumi.String("foobar-elb"),
//				AvailabilityZones: pulumi.StringArray{
//					pulumi.String("us-east-1c"),
//				},
//				Listeners: elb.LoadBalancerListenerArray{
//					&elb.LoadBalancerListenerArgs{
//						InstancePort:     pulumi.Int(80),
//						InstanceProtocol: pulumi.String("http"),
//						LbPort:           pulumi.Int(80),
//						LbProtocol:       pulumi.String("http"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = route53.NewRecord(ctx, "www", &route53.RecordArgs{
//				ZoneId: pulumi.Any(primary.ZoneId),
//				Name:   pulumi.String("example.com"),
//				Type:   pulumi.String(route53.RecordTypeA),
//				Aliases: route53.RecordAliasArray{
//					&route53.RecordAliasArgs{
//						Name:                 main.DnsName,
//						ZoneId:               main.ZoneId,
//						EvaluateTargetHealth: pulumi.Bool(true),
//					},
//				},
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
// ### NS and SOA Record Management
//
// When creating Route 53 zones, the `NS` and `SOA` records for the zone are automatically created. Enabling the `allowOverwrite` argument will allow managing these records in a single deployment without the requirement for `import`.
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
//			example, err := route53.NewZone(ctx, "example", &route53.ZoneArgs{
//				Name: pulumi.String("test.example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = route53.NewRecord(ctx, "example", &route53.RecordArgs{
//				AllowOverwrite: pulumi.Bool(true),
//				Name:           pulumi.String("test.example.com"),
//				Ttl:            pulumi.Int(172800),
//				Type:           pulumi.String(route53.RecordTypeNS),
//				ZoneId:         example.ZoneId,
//				Records: pulumi.StringArray{
//					example.NameServers.ApplyT(func(nameServers []string) (string, error) {
//						return nameServers[0], nil
//					}).(pulumi.StringOutput),
//					example.NameServers.ApplyT(func(nameServers []string) (string, error) {
//						return nameServers[1], nil
//					}).(pulumi.StringOutput),
//					example.NameServers.ApplyT(func(nameServers []string) (string, error) {
//						return nameServers[2], nil
//					}).(pulumi.StringOutput),
//					example.NameServers.ApplyT(func(nameServers []string) (string, error) {
//						return nameServers[3], nil
//					}).(pulumi.StringOutput),
//				},
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
// If the record also contains a set identifier, append it:
//
// If the record name is the empty string, it can be omitted:
//
// __Using `pulumi import` to import__ Route53 Records using the ID of the record, record name, record type, and set identifier. For example:
//
// Using the ID of the record, which is the zone identifier, record name, and record type, separated by underscores (`_`):
//
// ```sh
// $ pulumi import aws:route53/record:Record myrecord Z4KAPRWWNC7JR_dev.example.com_NS
// ```
// If the record also contains a set identifier, append it:
//
// ```sh
// $ pulumi import aws:route53/record:Record myrecord Z4KAPRWWNC7JR_dev.example.com_NS_dev
// ```
type Record struct {
	pulumi.CustomResourceState

	// An alias block. Conflicts with `ttl` & `records`.
	// Documented below.
	Aliases RecordAliasArrayOutput `pulumi:"aliases"`
	// Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.
	//
	// Exactly one of `records` or `alias` must be specified: this determines whether it's an alias record.
	AllowOverwrite pulumi.BoolOutput `pulumi:"allowOverwrite"`
	// A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
	CidrRoutingPolicy RecordCidrRoutingPolicyPtrOutput `pulumi:"cidrRoutingPolicy"`
	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	FailoverRoutingPolicies RecordFailoverRoutingPolicyArrayOutput `pulumi:"failoverRoutingPolicies"`
	// [FQDN](https://en.wikipedia.org/wiki/Fully_qualified_domain_name) built using the zone domain and `name`.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	GeolocationRoutingPolicies RecordGeolocationRoutingPolicyArrayOutput `pulumi:"geolocationRoutingPolicies"`
	// A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
	GeoproximityRoutingPolicy RecordGeoproximityRoutingPolicyPtrOutput `pulumi:"geoproximityRoutingPolicy"`
	// The health check the record should be associated with.
	HealthCheckId pulumi.StringPtrOutput `pulumi:"healthCheckId"`
	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	LatencyRoutingPolicies RecordLatencyRoutingPolicyArrayOutput `pulumi:"latencyRoutingPolicies"`
	// Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	MultivalueAnswerRoutingPolicy pulumi.BoolPtrOutput `pulumi:"multivalueAnswerRoutingPolicy"`
	// The name of the record.
	Name pulumi.StringOutput `pulumi:"name"`
	// A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the provider configuration string (e.g., `"first255characters\"\"morecharacters"`).
	Records pulumi.StringArrayOutput `pulumi:"records"`
	// Unique identifier to differentiate records with routing policies from one another. Required if using `cidrRoutingPolicy`, `failoverRoutingPolicy`, `geolocationRoutingPolicy`,`geoproximityRoutingPolicy`, `latencyRoutingPolicy`, `multivalueAnswerRoutingPolicy`, or `weightedRoutingPolicy`.
	SetIdentifier pulumi.StringPtrOutput `pulumi:"setIdentifier"`
	// The TTL of the record.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// The record type. Valid values are `A`, `AAAA`, `CAA`, `CNAME`, `DS`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV` and `TXT`.
	Type pulumi.StringOutput `pulumi:"type"`
	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	WeightedRoutingPolicies RecordWeightedRoutingPolicyArrayOutput `pulumi:"weightedRoutingPolicies"`
	// The ID of the hosted zone to contain this record.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewRecord registers a new resource with the given unique name, arguments, and options.
func NewRecord(ctx *pulumi.Context,
	name string, args *RecordArgs, opts ...pulumi.ResourceOption) (*Record, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Record
	err := ctx.RegisterResource("aws:route53/record:Record", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecord gets an existing Record resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordState, opts ...pulumi.ResourceOption) (*Record, error) {
	var resource Record
	err := ctx.ReadResource("aws:route53/record:Record", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Record resources.
type recordState struct {
	// An alias block. Conflicts with `ttl` & `records`.
	// Documented below.
	Aliases []RecordAlias `pulumi:"aliases"`
	// Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.
	//
	// Exactly one of `records` or `alias` must be specified: this determines whether it's an alias record.
	AllowOverwrite *bool `pulumi:"allowOverwrite"`
	// A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
	CidrRoutingPolicy *RecordCidrRoutingPolicy `pulumi:"cidrRoutingPolicy"`
	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	FailoverRoutingPolicies []RecordFailoverRoutingPolicy `pulumi:"failoverRoutingPolicies"`
	// [FQDN](https://en.wikipedia.org/wiki/Fully_qualified_domain_name) built using the zone domain and `name`.
	Fqdn *string `pulumi:"fqdn"`
	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	GeolocationRoutingPolicies []RecordGeolocationRoutingPolicy `pulumi:"geolocationRoutingPolicies"`
	// A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
	GeoproximityRoutingPolicy *RecordGeoproximityRoutingPolicy `pulumi:"geoproximityRoutingPolicy"`
	// The health check the record should be associated with.
	HealthCheckId *string `pulumi:"healthCheckId"`
	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	LatencyRoutingPolicies []RecordLatencyRoutingPolicy `pulumi:"latencyRoutingPolicies"`
	// Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	MultivalueAnswerRoutingPolicy *bool `pulumi:"multivalueAnswerRoutingPolicy"`
	// The name of the record.
	Name *string `pulumi:"name"`
	// A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the provider configuration string (e.g., `"first255characters\"\"morecharacters"`).
	Records []string `pulumi:"records"`
	// Unique identifier to differentiate records with routing policies from one another. Required if using `cidrRoutingPolicy`, `failoverRoutingPolicy`, `geolocationRoutingPolicy`,`geoproximityRoutingPolicy`, `latencyRoutingPolicy`, `multivalueAnswerRoutingPolicy`, or `weightedRoutingPolicy`.
	SetIdentifier *string `pulumi:"setIdentifier"`
	// The TTL of the record.
	Ttl *int `pulumi:"ttl"`
	// The record type. Valid values are `A`, `AAAA`, `CAA`, `CNAME`, `DS`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV` and `TXT`.
	Type *string `pulumi:"type"`
	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	WeightedRoutingPolicies []RecordWeightedRoutingPolicy `pulumi:"weightedRoutingPolicies"`
	// The ID of the hosted zone to contain this record.
	ZoneId *string `pulumi:"zoneId"`
}

type RecordState struct {
	// An alias block. Conflicts with `ttl` & `records`.
	// Documented below.
	Aliases RecordAliasArrayInput
	// Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.
	//
	// Exactly one of `records` or `alias` must be specified: this determines whether it's an alias record.
	AllowOverwrite pulumi.BoolPtrInput
	// A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
	CidrRoutingPolicy RecordCidrRoutingPolicyPtrInput
	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	FailoverRoutingPolicies RecordFailoverRoutingPolicyArrayInput
	// [FQDN](https://en.wikipedia.org/wiki/Fully_qualified_domain_name) built using the zone domain and `name`.
	Fqdn pulumi.StringPtrInput
	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	GeolocationRoutingPolicies RecordGeolocationRoutingPolicyArrayInput
	// A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
	GeoproximityRoutingPolicy RecordGeoproximityRoutingPolicyPtrInput
	// The health check the record should be associated with.
	HealthCheckId pulumi.StringPtrInput
	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	LatencyRoutingPolicies RecordLatencyRoutingPolicyArrayInput
	// Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	MultivalueAnswerRoutingPolicy pulumi.BoolPtrInput
	// The name of the record.
	Name pulumi.StringPtrInput
	// A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the provider configuration string (e.g., `"first255characters\"\"morecharacters"`).
	Records pulumi.StringArrayInput
	// Unique identifier to differentiate records with routing policies from one another. Required if using `cidrRoutingPolicy`, `failoverRoutingPolicy`, `geolocationRoutingPolicy`,`geoproximityRoutingPolicy`, `latencyRoutingPolicy`, `multivalueAnswerRoutingPolicy`, or `weightedRoutingPolicy`.
	SetIdentifier pulumi.StringPtrInput
	// The TTL of the record.
	Ttl pulumi.IntPtrInput
	// The record type. Valid values are `A`, `AAAA`, `CAA`, `CNAME`, `DS`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV` and `TXT`.
	Type pulumi.StringPtrInput
	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	WeightedRoutingPolicies RecordWeightedRoutingPolicyArrayInput
	// The ID of the hosted zone to contain this record.
	ZoneId pulumi.StringPtrInput
}

func (RecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordState)(nil)).Elem()
}

type recordArgs struct {
	// An alias block. Conflicts with `ttl` & `records`.
	// Documented below.
	Aliases []RecordAlias `pulumi:"aliases"`
	// Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.
	//
	// Exactly one of `records` or `alias` must be specified: this determines whether it's an alias record.
	AllowOverwrite *bool `pulumi:"allowOverwrite"`
	// A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
	CidrRoutingPolicy *RecordCidrRoutingPolicy `pulumi:"cidrRoutingPolicy"`
	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	FailoverRoutingPolicies []RecordFailoverRoutingPolicy `pulumi:"failoverRoutingPolicies"`
	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	GeolocationRoutingPolicies []RecordGeolocationRoutingPolicy `pulumi:"geolocationRoutingPolicies"`
	// A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
	GeoproximityRoutingPolicy *RecordGeoproximityRoutingPolicy `pulumi:"geoproximityRoutingPolicy"`
	// The health check the record should be associated with.
	HealthCheckId *string `pulumi:"healthCheckId"`
	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	LatencyRoutingPolicies []RecordLatencyRoutingPolicy `pulumi:"latencyRoutingPolicies"`
	// Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	MultivalueAnswerRoutingPolicy *bool `pulumi:"multivalueAnswerRoutingPolicy"`
	// The name of the record.
	Name string `pulumi:"name"`
	// A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the provider configuration string (e.g., `"first255characters\"\"morecharacters"`).
	Records []string `pulumi:"records"`
	// Unique identifier to differentiate records with routing policies from one another. Required if using `cidrRoutingPolicy`, `failoverRoutingPolicy`, `geolocationRoutingPolicy`,`geoproximityRoutingPolicy`, `latencyRoutingPolicy`, `multivalueAnswerRoutingPolicy`, or `weightedRoutingPolicy`.
	SetIdentifier *string `pulumi:"setIdentifier"`
	// The TTL of the record.
	Ttl *int `pulumi:"ttl"`
	// The record type. Valid values are `A`, `AAAA`, `CAA`, `CNAME`, `DS`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV` and `TXT`.
	Type string `pulumi:"type"`
	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	WeightedRoutingPolicies []RecordWeightedRoutingPolicy `pulumi:"weightedRoutingPolicies"`
	// The ID of the hosted zone to contain this record.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Record resource.
type RecordArgs struct {
	// An alias block. Conflicts with `ttl` & `records`.
	// Documented below.
	Aliases RecordAliasArrayInput
	// Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.
	//
	// Exactly one of `records` or `alias` must be specified: this determines whether it's an alias record.
	AllowOverwrite pulumi.BoolPtrInput
	// A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
	CidrRoutingPolicy RecordCidrRoutingPolicyPtrInput
	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	FailoverRoutingPolicies RecordFailoverRoutingPolicyArrayInput
	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	GeolocationRoutingPolicies RecordGeolocationRoutingPolicyArrayInput
	// A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
	GeoproximityRoutingPolicy RecordGeoproximityRoutingPolicyPtrInput
	// The health check the record should be associated with.
	HealthCheckId pulumi.StringPtrInput
	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	LatencyRoutingPolicies RecordLatencyRoutingPolicyArrayInput
	// Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	MultivalueAnswerRoutingPolicy pulumi.BoolPtrInput
	// The name of the record.
	Name pulumi.StringInput
	// A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the provider configuration string (e.g., `"first255characters\"\"morecharacters"`).
	Records pulumi.StringArrayInput
	// Unique identifier to differentiate records with routing policies from one another. Required if using `cidrRoutingPolicy`, `failoverRoutingPolicy`, `geolocationRoutingPolicy`,`geoproximityRoutingPolicy`, `latencyRoutingPolicy`, `multivalueAnswerRoutingPolicy`, or `weightedRoutingPolicy`.
	SetIdentifier pulumi.StringPtrInput
	// The TTL of the record.
	Ttl pulumi.IntPtrInput
	// The record type. Valid values are `A`, `AAAA`, `CAA`, `CNAME`, `DS`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV` and `TXT`.
	Type pulumi.StringInput
	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	WeightedRoutingPolicies RecordWeightedRoutingPolicyArrayInput
	// The ID of the hosted zone to contain this record.
	ZoneId pulumi.StringInput
}

func (RecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordArgs)(nil)).Elem()
}

type RecordInput interface {
	pulumi.Input

	ToRecordOutput() RecordOutput
	ToRecordOutputWithContext(ctx context.Context) RecordOutput
}

func (*Record) ElementType() reflect.Type {
	return reflect.TypeOf((**Record)(nil)).Elem()
}

func (i *Record) ToRecordOutput() RecordOutput {
	return i.ToRecordOutputWithContext(context.Background())
}

func (i *Record) ToRecordOutputWithContext(ctx context.Context) RecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordOutput)
}

// RecordArrayInput is an input type that accepts RecordArray and RecordArrayOutput values.
// You can construct a concrete instance of `RecordArrayInput` via:
//
//	RecordArray{ RecordArgs{...} }
type RecordArrayInput interface {
	pulumi.Input

	ToRecordArrayOutput() RecordArrayOutput
	ToRecordArrayOutputWithContext(context.Context) RecordArrayOutput
}

type RecordArray []RecordInput

func (RecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Record)(nil)).Elem()
}

func (i RecordArray) ToRecordArrayOutput() RecordArrayOutput {
	return i.ToRecordArrayOutputWithContext(context.Background())
}

func (i RecordArray) ToRecordArrayOutputWithContext(ctx context.Context) RecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordArrayOutput)
}

// RecordMapInput is an input type that accepts RecordMap and RecordMapOutput values.
// You can construct a concrete instance of `RecordMapInput` via:
//
//	RecordMap{ "key": RecordArgs{...} }
type RecordMapInput interface {
	pulumi.Input

	ToRecordMapOutput() RecordMapOutput
	ToRecordMapOutputWithContext(context.Context) RecordMapOutput
}

type RecordMap map[string]RecordInput

func (RecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Record)(nil)).Elem()
}

func (i RecordMap) ToRecordMapOutput() RecordMapOutput {
	return i.ToRecordMapOutputWithContext(context.Background())
}

func (i RecordMap) ToRecordMapOutputWithContext(ctx context.Context) RecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordMapOutput)
}

type RecordOutput struct{ *pulumi.OutputState }

func (RecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Record)(nil)).Elem()
}

func (o RecordOutput) ToRecordOutput() RecordOutput {
	return o
}

func (o RecordOutput) ToRecordOutputWithContext(ctx context.Context) RecordOutput {
	return o
}

// An alias block. Conflicts with `ttl` & `records`.
// Documented below.
func (o RecordOutput) Aliases() RecordAliasArrayOutput {
	return o.ApplyT(func(v *Record) RecordAliasArrayOutput { return v.Aliases }).(RecordAliasArrayOutput)
}

// Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.
//
// Exactly one of `records` or `alias` must be specified: this determines whether it's an alias record.
func (o RecordOutput) AllowOverwrite() pulumi.BoolOutput {
	return o.ApplyT(func(v *Record) pulumi.BoolOutput { return v.AllowOverwrite }).(pulumi.BoolOutput)
}

// A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
func (o RecordOutput) CidrRoutingPolicy() RecordCidrRoutingPolicyPtrOutput {
	return o.ApplyT(func(v *Record) RecordCidrRoutingPolicyPtrOutput { return v.CidrRoutingPolicy }).(RecordCidrRoutingPolicyPtrOutput)
}

// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
func (o RecordOutput) FailoverRoutingPolicies() RecordFailoverRoutingPolicyArrayOutput {
	return o.ApplyT(func(v *Record) RecordFailoverRoutingPolicyArrayOutput { return v.FailoverRoutingPolicies }).(RecordFailoverRoutingPolicyArrayOutput)
}

// [FQDN](https://en.wikipedia.org/wiki/Fully_qualified_domain_name) built using the zone domain and `name`.
func (o RecordOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Record) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
func (o RecordOutput) GeolocationRoutingPolicies() RecordGeolocationRoutingPolicyArrayOutput {
	return o.ApplyT(func(v *Record) RecordGeolocationRoutingPolicyArrayOutput { return v.GeolocationRoutingPolicies }).(RecordGeolocationRoutingPolicyArrayOutput)
}

// A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
func (o RecordOutput) GeoproximityRoutingPolicy() RecordGeoproximityRoutingPolicyPtrOutput {
	return o.ApplyT(func(v *Record) RecordGeoproximityRoutingPolicyPtrOutput { return v.GeoproximityRoutingPolicy }).(RecordGeoproximityRoutingPolicyPtrOutput)
}

// The health check the record should be associated with.
func (o RecordOutput) HealthCheckId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Record) pulumi.StringPtrOutput { return v.HealthCheckId }).(pulumi.StringPtrOutput)
}

// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
func (o RecordOutput) LatencyRoutingPolicies() RecordLatencyRoutingPolicyArrayOutput {
	return o.ApplyT(func(v *Record) RecordLatencyRoutingPolicyArrayOutput { return v.LatencyRoutingPolicies }).(RecordLatencyRoutingPolicyArrayOutput)
}

// Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
func (o RecordOutput) MultivalueAnswerRoutingPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Record) pulumi.BoolPtrOutput { return v.MultivalueAnswerRoutingPolicy }).(pulumi.BoolPtrOutput)
}

// The name of the record.
func (o RecordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Record) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the provider configuration string (e.g., `"first255characters\"\"morecharacters"`).
func (o RecordOutput) Records() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Record) pulumi.StringArrayOutput { return v.Records }).(pulumi.StringArrayOutput)
}

// Unique identifier to differentiate records with routing policies from one another. Required if using `cidrRoutingPolicy`, `failoverRoutingPolicy`, `geolocationRoutingPolicy`,`geoproximityRoutingPolicy`, `latencyRoutingPolicy`, `multivalueAnswerRoutingPolicy`, or `weightedRoutingPolicy`.
func (o RecordOutput) SetIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Record) pulumi.StringPtrOutput { return v.SetIdentifier }).(pulumi.StringPtrOutput)
}

// The TTL of the record.
func (o RecordOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Record) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The record type. Valid values are `A`, `AAAA`, `CAA`, `CNAME`, `DS`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV` and `TXT`.
func (o RecordOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Record) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
func (o RecordOutput) WeightedRoutingPolicies() RecordWeightedRoutingPolicyArrayOutput {
	return o.ApplyT(func(v *Record) RecordWeightedRoutingPolicyArrayOutput { return v.WeightedRoutingPolicies }).(RecordWeightedRoutingPolicyArrayOutput)
}

// The ID of the hosted zone to contain this record.
func (o RecordOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Record) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type RecordArrayOutput struct{ *pulumi.OutputState }

func (RecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Record)(nil)).Elem()
}

func (o RecordArrayOutput) ToRecordArrayOutput() RecordArrayOutput {
	return o
}

func (o RecordArrayOutput) ToRecordArrayOutputWithContext(ctx context.Context) RecordArrayOutput {
	return o
}

func (o RecordArrayOutput) Index(i pulumi.IntInput) RecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Record {
		return vs[0].([]*Record)[vs[1].(int)]
	}).(RecordOutput)
}

type RecordMapOutput struct{ *pulumi.OutputState }

func (RecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Record)(nil)).Elem()
}

func (o RecordMapOutput) ToRecordMapOutput() RecordMapOutput {
	return o
}

func (o RecordMapOutput) ToRecordMapOutputWithContext(ctx context.Context) RecordMapOutput {
	return o
}

func (o RecordMapOutput) MapIndex(k pulumi.StringInput) RecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Record {
		return vs[0].(map[string]*Record)[vs[1].(string)]
	}).(RecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RecordInput)(nil)).Elem(), &Record{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecordArrayInput)(nil)).Elem(), RecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecordMapInput)(nil)).Elem(), RecordMap{})
	pulumi.RegisterOutputType(RecordOutput{})
	pulumi.RegisterOutputType(RecordArrayOutput{})
	pulumi.RegisterOutputType(RecordMapOutput{})
}