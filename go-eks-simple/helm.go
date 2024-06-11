package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewayRelease struct {
	Values   *GatewayValues
	Provider *kubernetes.Provider
	Output   *helm.Release
	Service  *corev1.Service
}

func (gr *GatewayRelease) New(ctx *pulumi.Context) error {
	crds, err := yaml.NewConfigFile(ctx,
		"gw-api-crd",
		&yaml.ConfigFileArgs{File: "./config/gw-api-crd.yaml"},
		pulumi.Provider(gr.Provider),
	)
	if err != nil {
		return nil
	}

	hr, err := helm.NewRelease(ctx, "nginx-gateway", &helm.ReleaseArgs{
		Name:            pulumi.String("nginx-gateway"),
		Chart:           pulumi.String(gr.Values.Chart),
		Namespace:       pulumi.String(gr.Values.Namespace),
		CreateNamespace: pulumi.Bool(true),
		Values: pulumi.Map{
			"service": pulumi.Map{
				"create": pulumi.Bool(false),
			},
		},
	},
		pulumi.Provider(gr.Provider),
		pulumi.DependsOn([]pulumi.Resource{crds}),
	)
	if err != nil {
		return err
	}

	svc, err := corev1.NewService(
		ctx,
		"nginx-gateway-fabric",
		&corev1.ServiceArgs{
			ApiVersion: pulumi.String("v1"),
			Metadata: metav1.ObjectMetaArgs{
				Annotations: pulumi.StringMap{
					"meta.helm.sh/release-name":                         pulumi.String("nginx-gateway"),
					"meta.helm.sh/release-namespace":                    pulumi.String("nginx-gateway"),
					"service.beta.kubernetes.io/aws-load-balancer-type": pulumi.String("nlb"),
				},
				Labels: pulumi.StringMap{
					"app.kubernetes.io/instance":   pulumi.String("nginx-gateway"),
					"app.kubernetes.io/managed-by": pulumi.String("Helm"),
					"app.kubernetes.io/name":       pulumi.String("nginx-gateway-fabric"),
					"app.kubernetes.io/version":    pulumi.String("1.2.0"),
					"helm.sh/chart":                pulumi.String("nginx-gateway-fabric-1.2.0"),
				},
				Name:      pulumi.String("nginx-gateway-nginx-gateway-fabric"),
				Namespace: pulumi.String("nginx-gateway"),
			},
			Spec: corev1.ServiceSpecArgs{
				ExternalTrafficPolicy: pulumi.String("Local"),
				Ports: corev1.ServicePortArray{
					corev1.ServicePortArgs{
						Name:       pulumi.String("http"),
						NodePort:   pulumi.Int(32093),
						Port:       pulumi.Int(80),
						Protocol:   pulumi.String("TCP"),
						TargetPort: pulumi.Int(80),
					},
					corev1.ServicePortArgs{
						Name:       pulumi.String("https"),
						NodePort:   pulumi.Int(31827),
						Port:       pulumi.Int(443),
						Protocol:   pulumi.String("TCP"),
						TargetPort: pulumi.Int(443),
					},
				},
				Selector: pulumi.StringMap{
					"app.kubernetes.io/instance": pulumi.String("nginx-gateway"),
					"app.kubernetes.io/name":     pulumi.String("nginx-gateway-fabric"),
				},
				Type: pulumi.String("LoadBalancer"),
			},
		},
		pulumi.Provider(gr.Provider),
		pulumi.DependsOn([]pulumi.Resource{hr}),
	)

	if err != nil {
		return err
	}

	gr.Output = hr
	gr.Service = svc

	return nil
}
