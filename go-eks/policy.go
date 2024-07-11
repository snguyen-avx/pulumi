package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Iam struct {
	RoleName             string
	NgRoleName           string
	Role                 *iam.Role
	NodeGroupRole        *iam.Role
	PolicyAttachments    []*iam.RolePolicyAttachment
	NgrPolicyAttachments []*iam.RolePolicyAttachment
}

type SecurityGroup struct {
	Name   string
	VpcID  pulumi.IDOutput
	Output *ec2.SecurityGroup
}

type Policy struct {
	List   []string
	Role   pulumi.StringOutput
	Prefix string
}

func (im *Iam) New(ctx *pulumi.Context) error {

	role, err := iam.NewRole(ctx, im.RoleName, &iam.RoleArgs{
		AssumeRolePolicy: pulumi.String(`{
        "Version": "2008-10-17",
        "Statement": [{
            "Sid": "",
            "Effect": "Allow",
            "Principal": {
                "Service": "eks.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        }]
    }`),
	})

	if err != nil {
		return err
	}

	pas, err := im.NewPolicyAttachments(ctx, &Policy{
		List: []string{
			"arn:aws:iam::aws:policy/AmazonEKSServicePolicy",
			"arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
		},
		Role:   role.Name,
		Prefix: "rpa",
	})

	if err != nil {
		return err
	}

	// Create the EC2 NodeGroup Role
	nodeGroupRole, err := iam.NewRole(ctx, im.NgRoleName, &iam.RoleArgs{
		AssumeRolePolicy: pulumi.String(`{
		    "Version": "2012-10-17",
		    "Statement": [{
		        "Sid": "",
		        "Effect": "Allow",
		        "Principal": {
		            "Service": "ec2.amazonaws.com"
		        },
		        "Action": "sts:AssumeRole"
		    }]
		}`),
	})

	if err != nil {
		return err
	}

	ngPas, err := im.NewPolicyAttachments(ctx, &Policy{
		List: []string{
			"arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
			"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
			"arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
		},
		Role:   nodeGroupRole.Name,
		Prefix: "ngpa",
	})

	if err != nil {
		return err
	}

	im.Role = role
	im.NodeGroupRole = nodeGroupRole

	im.PolicyAttachments = pas
	im.NgrPolicyAttachments = ngPas

	return nil
}

func (im *Iam) NewPolicyAttachments(ctx *pulumi.Context, policy *Policy) ([]*iam.RolePolicyAttachment, error) {
	pas := []*iam.RolePolicyAttachment{}
	for i, p := range policy.List {
		pa, err := iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("%s-%d", policy.Prefix, i), &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String(p),
			Role:      policy.Role,
		})
		if err != nil {
			return nil, err
		}
		pas = append(pas, pa)
	}
	return pas, nil
}

func (sg *SecurityGroup) New(ctx *pulumi.Context) error {
	// Create a Security Group that we can use to actually connect to our cluster
	output, err := ec2.NewSecurityGroup(ctx, sg.Name, &ec2.SecurityGroupArgs{
		VpcId: sg.VpcID,
		Egress: ec2.SecurityGroupEgressArray{
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("-1"),
				FromPort:   pulumi.Int(0),
				ToPort:     pulumi.Int(0),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
		},
		Ingress: ec2.SecurityGroupIngressArray{
			ec2.SecurityGroupIngressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(80),
				ToPort:     pulumi.Int(80),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
		},
	})
	sg.Output = output
	return err
}
