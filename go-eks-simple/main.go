package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := NewConfig(ctx)

		iamRole := &Iam{
			RoleName:   cfg.IamRole,
			NgRoleName: cfg.NodeGroupRole,
		}
		if err := iamRole.New(ctx); err != nil {
			return err
		}
		ctx.Export("iamRole", iamRole.Role.ToRoleOutput())
		ctx.Export("nodeGroupRole", iamRole.NodeGroupRole.ToRoleOutput())

		vpc := &Vpc{Name: cfg.Vpc}
		if err := vpc.New(ctx); err != nil {
			return err
		}
		ctx.Export("vpc", vpc.ID)

		sg := &SecurityGroup{
			Name:  cfg.SecurityGroup,
			VpcID: vpc.ID,
		}
		if err := sg.New(ctx); err != nil {
			return err
		}
		ctx.Export("securityGroup", sg.Output)

		eksCluster := &Cluster{
			Name:             pulumi.String(cfg.Cluster),
			RoleArn:          iamRole.Role.Arn,
			NgRoleArn:        iamRole.NodeGroupRole.Arn,
			SgID:             sg.Output.ID().ToStringOutput(),
			VpcID:            vpc.ID,
			PublicSubnetIds:  vpc.PublicSubnets,
			PrivateSubnetIds: vpc.PrivateSubnets,
			Min:              cfg.MustInt(cfg.Min),
			Max:              cfg.MustInt(cfg.Max),
		}
		if err := eksCluster.New(ctx); err != nil {
			return err
		}
		ctx.Export("kubeconfig", eksCluster.Output.KubeconfigJson)

		return nil
	})
}
