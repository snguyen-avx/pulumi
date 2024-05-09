package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewayRelease struct {
	Values   *GatewayValues
	Provider *kubernetes.Provider
	Output   *helm.Release
}

func (gr *GatewayRelease) New(ctx *pulumi.Context) error {
	_, err := yaml.NewConfigFile(ctx,
		"./config/gw-api-crd.yaml",
		&yaml.ConfigFileArgs{File: gr.Values.Crds},
		pulumi.Provider(gr.Provider),
	)
	if err != nil {
		return nil
	}

	hr, err := helm.NewRelease(ctx, "nginx-gateway", &helm.ReleaseArgs{
		Chart:           pulumi.String(gr.Values.Chart),
		Namespace:       pulumi.String(gr.Values.Namespace),
		CreateNamespace: pulumi.Bool(true),
	},
		pulumi.Provider(gr.Provider),
	)
	if err != nil {
		return err
	}

	gr.Output = hr

	return nil
}
