package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	ec2x "github.com/pulumi/pulumi-awsx/sdk/go/awsx/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Vpc struct {
	Name           string
	ID             pulumi.IDOutput
	Output         ec2.VpcOutput
	PublicSubnets  pulumi.StringArrayOutput
	PrivateSubnets pulumi.StringArrayOutput
	Provider       pulumi.ProviderResource
}

type Subnet struct {
	ID     pulumi.IDOutput
	Output ec2.SubnetOutput
}

func GetSubnet(ctx *pulumi.Context) (*Subnet, error) {
	// Read back the default VPC and public subnets, which we will use.
	t := true

	vpc, err := ec2.LookupVpc(ctx, &ec2.LookupVpcArgs{Default: &t})

	if err != nil {
		return nil, err
	}

	_, err = ec2.GetSubnets(ctx, &ec2.GetSubnetsArgs{
		Filters: []ec2.GetSubnetsFilter{
			{Name: "vpc-id", Values: []string{vpc.Id}},
		},
	})

	if err != nil {
		return nil, err
	}

	return &Subnet{}, nil
}

func (v *Vpc) New(ctx *pulumi.Context) error {

	vpc, err := ec2x.NewVpc(ctx, v.Name, nil)
	if err != nil {
		return err
	}

	v.ID = pulumi.IDOutput(vpc.VpcId)
	v.Output = ec2.VpcOutput(vpc.ToVpcOutput())
	v.PublicSubnets = vpc.PublicSubnetIds
	v.PrivateSubnets = vpc.PrivateSubnetIds

	return nil
}
