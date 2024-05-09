package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := NewConfig(ctx)

		vpc := &Vpc{Name: cfg.Vpc}
		if err := vpc.New(ctx); err != nil {
			return err
		}
		ctx.Export("vpc", vpc.ID)

		eksCluster := &Cluster{
			Name:             pulumi.String(cfg.Cluster),
			VpcID:            vpc.ID,
			PublicSubnetIds:  vpc.PublicSubnets,
			PrivateSubnetIds: vpc.PrivateSubnets,
			InstanceType:     cfg.Type,
			Min:              cfg.MustInt(cfg.Min),
			Max:              cfg.MustInt(cfg.Max),
		}
		if err := eksCluster.New(ctx); err != nil {
			return err
		}
		ctx.Export("kubeconfig", eksCluster.Output.KubeconfigJson)

		gatewayRelease := &GatewayRelease{
			Provider: eksCluster.Provider,
			Values:   cfg.GatewayValues,
		}
		if err := gatewayRelease.New(ctx); err != nil {
			return err
		}

		ctx.Export("gatewayChart", gatewayRelease.Output.Chart)
		ctx.Export("gatewayValues", gatewayRelease.Output.Values)

		return nil
	})
}
