package main

import (
	"fmt"
	"log"
	"strconv"

	ec2x "github.com/pulumi/pulumi-awsx/sdk/go/awsx/ec2"
	eksx "github.com/pulumi/pulumi-eks/sdk/go/eks"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
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
}

func NewConfig(ctx *pulumi.Context) *Config {
	cfg := config.New(ctx, "")
	return &Config{
		Vpc:           cfg.Require("vpcName"),
		Cluster:       cfg.Require("clusterName"),
		SecurityGroup: cfg.Require("clusterSg"),
		IamRole:       cfg.Require("iamRole"),
		NodeGroupRole: cfg.Require("ngrRole"),
		Min:           cfg.Require("minSize"),
		Max:           cfg.Require("maxSize"),
		Type:          cfg.Require("instanceType"),
	}
}

func (c Config) MustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to MustInt: %v", err)
	}
	return i
}

type Vpc struct {
	Name           string
	ID             pulumi.IDOutput
	Output         ec2.VpcOutput
	PublicSubnets  pulumi.StringArrayOutput
	PrivateSubnets pulumi.StringArrayOutput
}

type Subnet struct {
	ID     pulumi.IDOutput
	Output ec2.SubnetOutput
}

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

type Cluster struct {
	Name             pulumi.String
	RoleArn          pulumi.StringOutput
	NgRoleArn        pulumi.StringOutput
	SgID             pulumi.StringOutput
	VpcID            pulumi.IDOutput
	PublicSubnetIds  pulumi.StringArrayOutput
	PrivateSubnetIds pulumi.StringArrayOutput
	Output           *eksx.Cluster
	NodeGroup        *eks.NodeGroup
	Min              int
	Max              int
	KubeConfig       string
	Provider         *kubernetes.Provider
}

type Policy struct {
	List   []string
	Role   pulumi.StringOutput
	Prefix string
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

	return &Subnet{
		// LookupVpcOutput: vpc,
		// GetOutput:       subnet,
	}, nil
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

func (c *Cluster) New(ctx *pulumi.Context) error {
	// Create EKS Cluster
	output, err := eksx.NewCluster(ctx, string(c.Name), &eksx.ClusterArgs{
		VpcId:                        c.VpcID,
		PublicSubnetIds:              c.PublicSubnetIds,
		PrivateSubnetIds:             c.PrivateSubnetIds,
		NodeAssociatePublicIpAddress: pulumi.BoolRef(false),
	})
	if err != nil {
		return err
	}

	c.Output = output

	k8sProvider, err := kubernetes.NewProvider(ctx, "k8sprovider", &kubernetes.ProviderArgs{
		Kubeconfig: output.KubeconfigJson.ToStringPtrOutput(),
	}, pulumi.DependsOn([]pulumi.Resource{
		output,
	}))
	if err != nil {
		return err
	}
	c.Provider = k8sProvider

	_, err = corev1.NewNamespace(ctx, "sre", &corev1.NamespaceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String("sre"),
		},
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return err
	}

	_, err = helm.NewRelease(ctx, "nginx-ingress", &helm.ReleaseArgs{
		Chart: pulumi.String("ingress-nginx"),
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://kubernetes.github.io/ingress-nginx"),
		},
		Namespace:       pulumi.String("nginx"),
		CreateNamespace: pulumi.Bool(true),
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return err
	}

	return nil
}
