package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg, err := NewConfig(ctx)

		if err != nil {
			return err
		}

		vpc := &Vpc{
			Name:     cfg.Vpc,
			Provider: cfg.Provider,
		}
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
			// AwsProvider:      cfg.Provider,
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

		if gatewayRelease.Output != nil {
			ctx.Export("gatewayChart", gatewayRelease.Output.Chart)
			ctx.Export("gatewayValues", gatewayRelease.Output.Values)
			ctx.Export("nlb", gatewayRelease.Service)
		}

		nr, err := route53.NewRecord(ctx, "wild-route",
			&route53.RecordArgs{
				ZoneId: pulumi.String(cfg.ZoneId),
				Name:   pulumi.String(cfg.RouteName),
				Type:   pulumi.String(route53.RecordTypeCNAME),
				Ttl:    pulumi.Int(60),
				Records: pulumi.StringArray{
					gatewayRelease.Service.Status.LoadBalancer().
						Elem().Ingress().Index(pulumi.Int(0)).
						Hostname().Elem().ToStringOutput(),
				},
			},
			pulumi.DependsOn([]pulumi.Resource{gatewayRelease.Service}),
			pulumi.Provider(cfg.Provider),
		)

		if err != nil {
			return err
		}

		ctx.Export("route53Record", nr)

		return nil
	})
}
