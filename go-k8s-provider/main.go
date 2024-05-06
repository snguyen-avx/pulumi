package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		cfg := config.New(ctx, "")
		clusterName := cfg.Require("clusterName")

		eksClusterCurrent, err := eks.LookupCluster(ctx, &eks.LookupClusterArgs{
			Name: clusterName,
		})
		if err != nil {
			return err
		}

		kcfg := generateKubeconfig(
			pulumi.String(eksClusterCurrent.Endpoint).ToStringOutput(),
			pulumi.String(eksClusterCurrent.CertificateAuthorities[0].Data).ToStringOutput(),
			pulumi.String(eksClusterCurrent.Name).ToStringOutput(),
		)

		k8sProvider, err := kubernetes.NewProvider(ctx, "k8sprovider", &kubernetes.ProviderArgs{
			Kubeconfig: kcfg,
		})
		if err != nil {
			return err
		}

		ctx.Export("k8sProvider", k8sProvider)
		ctx.Export("kubeconfig", kcfg)

		return nil
	})
}

func generateKubeconfig(clusterEndpoint pulumi.StringOutput, certData pulumi.StringOutput, clusterName pulumi.StringOutput) pulumi.StringOutput {
	return pulumi.Sprintf(`{
        "apiVersion": "v1",
        "clusters": [{
            "cluster": {
                "server": "%s",
                "certificate-authority-data": "%s"
            },
            "name": "kubernetes",
        }],
        "contexts": [{
            "context": {
                "cluster": "kubernetes",
                "user": "aws",
            },
            "name": "aws",
        }],
        "current-context": "aws",
        "kind": "Config",
        "users": [{
            "name": "aws",
            "user": {
                "exec": {
                    "apiVersion": "client.authentication.k8s.io/v1beta1",
                    "command": "aws-iam-authenticator",
                    "args": [
                        "token",
                        "-i",
                        "%s",
                    ],
                },
            },
        }],
    }`, clusterEndpoint, certData, clusterName)
}
