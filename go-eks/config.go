package main

import (
	"log"
	"strconv"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type Config struct {
	Vpc                string
	Cluster            string
	SecurityGroup      string
	IamRole            string
	NodeGroupRole      string
	Min                string
	Max                string
	Type               string
	ZoneId             string
	RouteName          string
	RoleToAssumeARN    string
	RoleToAssumeRegion string
	GatewayValues      *GatewayValues
	Provider           *aws.Provider
}

type GatewayValues struct {
	Chart     string
	Crds      string
	Namespace string
	Service   struct {
		ExternalTrafficPolicy string
		Annotations           map[string]string
	}
}

func NewConfig(ctx *pulumi.Context) (*Config, error) {
	cfg := config.New(ctx, "")

	gwv := &GatewayValues{
		Chart:     "oci://ghcr.io/nginxinc/charts/nginx-gateway-fabric",
		Namespace: "nginx-gateway",
	}

	c := &Config{
		Vpc:                cfg.Require("vpcName"),
		Cluster:            cfg.Require("clusterName"),
		Min:                cfg.Require("minSize"),
		Max:                cfg.Require("maxSize"),
		Type:               cfg.Require("instanceType"),
		ZoneId:             cfg.Require("zoneID"),
		RouteName:          cfg.Require("routeName"),
		RoleToAssumeARN:    cfg.Require("roleToAssumeARN"),
		RoleToAssumeRegion: cfg.Require("roleToAssumeRegion"),
		GatewayValues:      gwv,
	}

	awsxProvider, err := aws.NewProvider(ctx, "awsx", &aws.ProviderArgs{
		AssumeRole: &aws.ProviderAssumeRoleArgs{
			RoleArn:     pulumi.StringPtr(c.RoleToAssumeARN),
			SessionName: pulumi.String("PulumiSession"),
			ExternalId:  pulumi.String("PulumiApplication"),
		},
		Region: pulumi.String(c.RoleToAssumeRegion),
	})

	if err != nil {
		return nil, err
	}

	c.Provider = awsxProvider

	return c, nil
}

func (c Config) MustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to MustInt: %v", err)
	}
	return i
}
