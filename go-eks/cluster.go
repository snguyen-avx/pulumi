package main

import (
	eksx "github.com/pulumi/pulumi-eks/sdk/go/eks"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
)

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
	InstanceType     string
	KubeConfig       string
	Provider         *kubernetes.Provider
	AwsProvider      *aws.Provider
}

func (c *Cluster) New(ctx *pulumi.Context) error {
	output, err := eksx.NewCluster(ctx, string(c.Name),
		&eksx.ClusterArgs{
			Name:             c.Name,
			VpcId:            c.VpcID,
			PublicSubnetIds:  c.PublicSubnetIds,
			PrivateSubnetIds: c.PrivateSubnetIds,
			MinSize:          pulumi.Int(c.Min),
			MaxSize:          pulumi.Int(c.Max),
			InstanceType:     pulumi.String(c.InstanceType),
		},
		pulumi.Provider(c.AwsProvider),
	)
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

	_, err = corev1.NewNamespace(ctx, "sre",
		&corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("sre"),
			},
		},
		pulumi.Provider(k8sProvider),
	)

	return err
}
