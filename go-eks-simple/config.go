package main

import (
	"log"
	"strconv"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type Config struct {
	Vpc           string
	Cluster       string
	SecurityGroup string
	IamRole       string
	NodeGroupRole string
	Min           string
	Max           string
	Type          string
	GatewayValues *GatewayValues
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

func NewConfig(ctx *pulumi.Context) *Config {
	cfg := config.New(ctx, "")

	gwv := &GatewayValues{
		Chart:     cfg.Require("helmChart"),
		Namespace: cfg.Require("helmNamespace"),
	}

	return &Config{
		Vpc:           cfg.Require("vpcName"),
		Cluster:       cfg.Require("clusterName"),
		Min:           cfg.Require("minSize"),
		Max:           cfg.Require("maxSize"),
		Type:          cfg.Require("instanceType"),
		GatewayValues: gwv,
	}
}

func (c Config) MustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to MustInt: %v", err)
	}
	return i
}
