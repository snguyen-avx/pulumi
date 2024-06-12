// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package helm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A `Release` is an instance of a chart running in a Kubernetes cluster. A `Chart` is a Helm package. It contains all the
// resource definitions necessary to run an application, tool, or service inside a Kubernetes cluster.
//
// This resource models a Helm Release as if it were created by the Helm CLI. The underlying implementation embeds Helm as
// a library to perform the orchestration of the resources. As a result, the full spectrum of Helm features are supported
// natively.
//
// ## Example Usage
// ### Local Chart Directory
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := helm.NewRelease(ctx, "nginx-ingress", helm.ReleaseArgs{
//				Chart: pulumi.String("./nginx-ingress"),
//			})
//			if err != nil {
//				return err
//			}
//
//			return nil
//		})
//	}
//
// ```
// ### Remote Chart
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := helm.NewRelease(ctx, "nginx-ingress", helm.ReleaseArgs{
//				Chart:   pulumi.String("nginx-ingress"),
//				Version: pulumi.String("1.24.4"),
//				RepositoryOpts: helm.RepositoryOptsArgs{
//					Repo: pulumi.String("https://charts.helm.sh/stable"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//
//			return nil
//		})
//	}
//
// ```
// ### Set Chart Values
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := helm.NewRelease(ctx, "nginx-ingress", helm.ReleaseArgs{
//				Chart:   pulumi.String("nginx-ingress"),
//				Version: pulumi.String("1.24.4"),
//				RepositoryOpts: helm.RepositoryOptsArgs{
//					Repo: pulumi.String("https://charts.helm.sh/stable"),
//				},
//				Values: pulumi.Map{
//					"controller": pulumi.Map{
//						"metrics": pulumi.Map{
//							"enabled": pulumi.Bool(true),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//
//			return nil
//		})
//	}
//
// ```
// ### Deploy Chart into Namespace
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := helm.NewRelease(ctx, "nginx-ingress", helm.ReleaseArgs{
//				Chart:     pulumi.String("nginx-ingress"),
//				Version:   pulumi.String("1.24.4"),
//				Namespace: pulumi.String("test-namespace"),
//				RepositoryOpts: helm.RepositoryOptsArgs{
//					Repo: pulumi.String("https://charts.helm.sh/stable"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//
//			return nil
//		})
//	}
//
// ```
//
// ### Depend on a Chart resource
// ```go
// package main
//
// import (
//
//	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			release, err := helm.NewRelease(ctx, "nginx-ingress", helm.ReleaseArgs{
//				Chart:     pulumi.String("nginx-ingress"),
//				Version:   pulumi.String("1.24.4"),
//				Namespace: pulumi.String("test-namespace"),
//				RepositoryOpts: helm.RepositoryOptsArgs{
//					Repo: pulumi.String("https://charts.helm.sh/stable"),
//				},
//				SkipAwait: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//
//			// Create a ConfigMap depending on the Chart. The ConfigMap will not be created until after all of the Chart
//			// resources are ready. Notice SkipAwait is set to false above. This is the default and will cause Helm
//			// to await the underlying resources to be available. Setting it to true will make the ConfigMap available right away.
//			_, err = corev1.NewConfigMap(ctx, "cm", &corev1.ConfigMapArgs{
//				Data: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//			}, pulumi.DependsOnInputs(release))
//			if err != nil {
//				return err
//			}
//
//			return nil
//		})
//	}
//
// ```
// ### Specify Helm Chart Values in File and Code
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := helm.NewRelease(ctx, "redis", helm.ReleaseArgs{
//				Chart:   pulumi.String("redis"),
//				RepositoryOpts: helm.RepositoryOptsArgs{
//					Repo: pulumi.String("https://charts.helm.sh/stable"),
//				},
//				ValueYamlFiles: pulumi.NewFileAsset("./metrics.yml"),
//				Value: pulumi.Map{
//					"cluster": pulumi.Map{
//						"enabled": pulumi.Bool(true),
//					},
//					"rbac": pulumi.Map{
//						"create": pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//
//			return nil
//		})
//	}
//
// // -- Contents of metrics.yml --
// // metrics:
// //     enabled: true
// ```
// ### Query Kubernetes Resource Installed By Helm Chart
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			rel, err := helm.NewRelease(ctx, "redis", &helm.ReleaseArgs{
//				Chart: pulumi.String("redis"),
//				RepositoryOpts: helm.RepositoryOptsArgs{
//					Repo: pulumi.String("https://raw.githubusercontent.com/bitnami/charts/eb5f9a9513d987b519f0ecd732e7031241c50328/bitnami"),
//				},
//				Values: pulumi.Map{
//					"cluster": pulumi.Map{
//						"enabled": pulumi.Bool(true),
//					},
//					"rbac": pulumi.BoolMap{
//						"create": pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//
//			// srv will only resolve after the redis chart is installed.
//			srv := pulumi.All(rel.Status.Namespace(), rel.Status.Name()).
//				ApplyT(func(r interface{}) (interface{}, error) {
//					arr := r.([]interface{})
//					namespace := arr[0].(*string)
//					name := arr[1].(*string)
//					svc, err := corev1.GetService(ctx,
//						"redis-master-svc",
//						pulumi.ID(fmt.Sprintf("%s/%s-master", *namespace, *name)),
//						nil,
//					)
//					if err != nil {
//						return "", nil
//					}
//					return svc.Spec.ClusterIP(), nil
//				})
//			ctx.Export("redisMasterClusterIP", srv)
//
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// An existing Helm Release resource can be imported using its `type token`, `name` and identifier, e.g.
//
// ```sh
// $ pulumi import kubernetes:helm.sh/v3:Release myRelease <namespace>/<releaseName>
// ```
type Release struct {
	pulumi.CustomResourceState

	// Whether to allow Null values in helm chart configs.
	AllowNullValues pulumi.BoolPtrOutput `pulumi:"allowNullValues"`
	// If set, installation process purges chart on fail. `skipAwait` will be disabled automatically if atomic is used.
	Atomic pulumi.BoolPtrOutput `pulumi:"atomic"`
	// Chart name to be installed. A path may be used.
	Chart pulumi.StringOutput `pulumi:"chart"`
	// Allow deletion of new resources created in this upgrade when upgrade fails.
	CleanupOnFail pulumi.BoolPtrOutput `pulumi:"cleanupOnFail"`
	// Create the namespace if it does not exist.
	CreateNamespace pulumi.BoolPtrOutput `pulumi:"createNamespace"`
	// Run helm dependency update before installing the chart.
	DependencyUpdate pulumi.BoolPtrOutput `pulumi:"dependencyUpdate"`
	// Add a custom description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Use chart development versions, too. Equivalent to version '>0.0.0-0'. If `version` is set, this is ignored.
	Devel pulumi.BoolPtrOutput `pulumi:"devel"`
	// Prevent CRD hooks from running, but run other hooks.  See helm install --no-crd-hook
	DisableCRDHooks pulumi.BoolPtrOutput `pulumi:"disableCRDHooks"`
	// If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema
	DisableOpenapiValidation pulumi.BoolPtrOutput `pulumi:"disableOpenapiValidation"`
	// Prevent hooks from running.
	DisableWebhooks pulumi.BoolPtrOutput `pulumi:"disableWebhooks"`
	// Force resource update through delete/recreate if needed.
	ForceUpdate pulumi.BoolPtrOutput `pulumi:"forceUpdate"`
	// Location of public keys used for verification. Used only if `verify` is true
	Keyring pulumi.StringPtrOutput `pulumi:"keyring"`
	// Run helm lint when planning.
	Lint pulumi.BoolPtrOutput `pulumi:"lint"`
	// The rendered manifests as JSON. Not yet supported.
	Manifest pulumi.MapOutput `pulumi:"manifest"`
	// Limit the maximum number of revisions saved per release. Use 0 for no limit.
	MaxHistory pulumi.IntPtrOutput `pulumi:"maxHistory"`
	// Release name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Namespace to install the release into.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Postrender command to run.
	Postrender pulumi.StringPtrOutput `pulumi:"postrender"`
	// Perform pods restart during upgrade/rollback.
	RecreatePods pulumi.BoolPtrOutput `pulumi:"recreatePods"`
	// If set, render subchart notes along with the parent.
	RenderSubchartNotes pulumi.BoolPtrOutput `pulumi:"renderSubchartNotes"`
	// Re-use the given name, even if that name is already used. This is unsafe in production
	Replace pulumi.BoolPtrOutput `pulumi:"replace"`
	// Specification defining the Helm chart repository to use.
	RepositoryOpts RepositoryOptsPtrOutput `pulumi:"repositoryOpts"`
	// When upgrading, reset the values to the ones built into the chart.
	ResetValues pulumi.BoolPtrOutput `pulumi:"resetValues"`
	// Names of resources created by the release grouped by "kind/version".
	ResourceNames pulumi.StringArrayMapOutput `pulumi:"resourceNames"`
	// When upgrading, reuse the last release's values and merge in any overrides. If 'resetValues' is specified, this is ignored
	ReuseValues pulumi.BoolPtrOutput `pulumi:"reuseValues"`
	// By default, the provider waits until all resources are in a ready state before marking the release as successful. Setting this to true will skip such await logic.
	SkipAwait pulumi.BoolPtrOutput `pulumi:"skipAwait"`
	// If set, no CRDs will be installed. By default, CRDs are installed if not already present.
	SkipCrds pulumi.BoolPtrOutput `pulumi:"skipCrds"`
	// Status of the deployed release.
	Status ReleaseStatusOutput `pulumi:"status"`
	// Time in seconds to wait for any individual kubernetes operation.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// List of assets (raw yaml files). Content is read and merged with values (with values taking precedence).
	ValueYamlFiles pulumi.AssetOrArchiveArrayOutput `pulumi:"valueYamlFiles"`
	// Custom values set for the release.
	Values pulumi.MapOutput `pulumi:"values"`
	// Verify the package before installing it.
	Verify pulumi.BoolPtrOutput `pulumi:"verify"`
	// Specify the exact chart version to install. If this is not specified, the latest version is installed.
	Version pulumi.StringPtrOutput `pulumi:"version"`
	// Will wait until all Jobs have been completed before marking the release as successful. This is ignored if `skipAwait` is enabled.
	WaitForJobs pulumi.BoolPtrOutput `pulumi:"waitForJobs"`
}

// NewRelease registers a new resource with the given unique name, arguments, and options.
func NewRelease(ctx *pulumi.Context,
	name string, args *ReleaseArgs, opts ...pulumi.ResourceOption) (*Release, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Chart == nil {
		return nil, errors.New("invalid value for required argument 'Chart'")
	}
	args.Compat = pulumi.StringPtr("true")
	var resource Release
	err := ctx.RegisterResource("kubernetes:helm.sh/v3:Release", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRelease gets an existing Release resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseState, opts ...pulumi.ResourceOption) (*Release, error) {
	var resource Release
	err := ctx.ReadResource("kubernetes:helm.sh/v3:Release", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Release resources.
type releaseState struct {
}

type ReleaseState struct {
}

func (ReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseState)(nil)).Elem()
}

type releaseArgs struct {
	// Whether to allow Null values in helm chart configs.
	AllowNullValues *bool `pulumi:"allowNullValues"`
	// If set, installation process purges chart on fail. `skipAwait` will be disabled automatically if atomic is used.
	Atomic *bool `pulumi:"atomic"`
	// Chart name to be installed. A path may be used.
	Chart string `pulumi:"chart"`
	// Allow deletion of new resources created in this upgrade when upgrade fails.
	CleanupOnFail *bool   `pulumi:"cleanupOnFail"`
	Compat        *string `pulumi:"compat"`
	// Create the namespace if it does not exist.
	CreateNamespace *bool `pulumi:"createNamespace"`
	// Run helm dependency update before installing the chart.
	DependencyUpdate *bool `pulumi:"dependencyUpdate"`
	// Add a custom description
	Description *string `pulumi:"description"`
	// Use chart development versions, too. Equivalent to version '>0.0.0-0'. If `version` is set, this is ignored.
	Devel *bool `pulumi:"devel"`
	// Prevent CRD hooks from running, but run other hooks.  See helm install --no-crd-hook
	DisableCRDHooks *bool `pulumi:"disableCRDHooks"`
	// If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema
	DisableOpenapiValidation *bool `pulumi:"disableOpenapiValidation"`
	// Prevent hooks from running.
	DisableWebhooks *bool `pulumi:"disableWebhooks"`
	// Force resource update through delete/recreate if needed.
	ForceUpdate *bool `pulumi:"forceUpdate"`
	// Location of public keys used for verification. Used only if `verify` is true
	Keyring *string `pulumi:"keyring"`
	// Run helm lint when planning.
	Lint *bool `pulumi:"lint"`
	// The rendered manifests as JSON. Not yet supported.
	Manifest map[string]interface{} `pulumi:"manifest"`
	// Limit the maximum number of revisions saved per release. Use 0 for no limit.
	MaxHistory *int `pulumi:"maxHistory"`
	// Release name.
	Name *string `pulumi:"name"`
	// Namespace to install the release into.
	Namespace *string `pulumi:"namespace"`
	// Postrender command to run.
	Postrender *string `pulumi:"postrender"`
	// Perform pods restart during upgrade/rollback.
	RecreatePods *bool `pulumi:"recreatePods"`
	// If set, render subchart notes along with the parent.
	RenderSubchartNotes *bool `pulumi:"renderSubchartNotes"`
	// Re-use the given name, even if that name is already used. This is unsafe in production
	Replace *bool `pulumi:"replace"`
	// Specification defining the Helm chart repository to use.
	RepositoryOpts *RepositoryOpts `pulumi:"repositoryOpts"`
	// When upgrading, reset the values to the ones built into the chart.
	ResetValues *bool `pulumi:"resetValues"`
	// Names of resources created by the release grouped by "kind/version".
	ResourceNames map[string][]string `pulumi:"resourceNames"`
	// When upgrading, reuse the last release's values and merge in any overrides. If 'resetValues' is specified, this is ignored
	ReuseValues *bool `pulumi:"reuseValues"`
	// By default, the provider waits until all resources are in a ready state before marking the release as successful. Setting this to true will skip such await logic.
	SkipAwait *bool `pulumi:"skipAwait"`
	// If set, no CRDs will be installed. By default, CRDs are installed if not already present.
	SkipCrds *bool `pulumi:"skipCrds"`
	// Time in seconds to wait for any individual kubernetes operation.
	Timeout *int `pulumi:"timeout"`
	// List of assets (raw yaml files). Content is read and merged with values.
	ValueYamlFiles []pulumi.AssetOrArchive `pulumi:"valueYamlFiles"`
	// Custom values set for the release.
	Values map[string]interface{} `pulumi:"values"`
	// Verify the package before installing it.
	Verify *bool `pulumi:"verify"`
	// Specify the exact chart version to install. If this is not specified, the latest version is installed.
	Version *string `pulumi:"version"`
	// Will wait until all Jobs have been completed before marking the release as successful. This is ignored if `skipAwait` is enabled.
	WaitForJobs *bool `pulumi:"waitForJobs"`
}

// The set of arguments for constructing a Release resource.
type ReleaseArgs struct {
	// Whether to allow Null values in helm chart configs.
	AllowNullValues pulumi.BoolPtrInput
	// If set, installation process purges chart on fail. `skipAwait` will be disabled automatically if atomic is used.
	Atomic pulumi.BoolPtrInput
	// Chart name to be installed. A path may be used.
	Chart pulumi.StringInput
	// Allow deletion of new resources created in this upgrade when upgrade fails.
	CleanupOnFail pulumi.BoolPtrInput
	Compat        pulumi.StringPtrInput
	// Create the namespace if it does not exist.
	CreateNamespace pulumi.BoolPtrInput
	// Run helm dependency update before installing the chart.
	DependencyUpdate pulumi.BoolPtrInput
	// Add a custom description
	Description pulumi.StringPtrInput
	// Use chart development versions, too. Equivalent to version '>0.0.0-0'. If `version` is set, this is ignored.
	Devel pulumi.BoolPtrInput
	// Prevent CRD hooks from running, but run other hooks.  See helm install --no-crd-hook
	DisableCRDHooks pulumi.BoolPtrInput
	// If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema
	DisableOpenapiValidation pulumi.BoolPtrInput
	// Prevent hooks from running.
	DisableWebhooks pulumi.BoolPtrInput
	// Force resource update through delete/recreate if needed.
	ForceUpdate pulumi.BoolPtrInput
	// Location of public keys used for verification. Used only if `verify` is true
	Keyring pulumi.StringPtrInput
	// Run helm lint when planning.
	Lint pulumi.BoolPtrInput
	// The rendered manifests as JSON. Not yet supported.
	Manifest pulumi.MapInput
	// Limit the maximum number of revisions saved per release. Use 0 for no limit.
	MaxHistory pulumi.IntPtrInput
	// Release name.
	Name pulumi.StringPtrInput
	// Namespace to install the release into.
	Namespace pulumi.StringPtrInput
	// Postrender command to run.
	Postrender pulumi.StringPtrInput
	// Perform pods restart during upgrade/rollback.
	RecreatePods pulumi.BoolPtrInput
	// If set, render subchart notes along with the parent.
	RenderSubchartNotes pulumi.BoolPtrInput
	// Re-use the given name, even if that name is already used. This is unsafe in production
	Replace pulumi.BoolPtrInput
	// Specification defining the Helm chart repository to use.
	RepositoryOpts RepositoryOptsPtrInput
	// When upgrading, reset the values to the ones built into the chart.
	ResetValues pulumi.BoolPtrInput
	// Names of resources created by the release grouped by "kind/version".
	ResourceNames pulumi.StringArrayMapInput
	// When upgrading, reuse the last release's values and merge in any overrides. If 'resetValues' is specified, this is ignored
	ReuseValues pulumi.BoolPtrInput
	// By default, the provider waits until all resources are in a ready state before marking the release as successful. Setting this to true will skip such await logic.
	SkipAwait pulumi.BoolPtrInput
	// If set, no CRDs will be installed. By default, CRDs are installed if not already present.
	SkipCrds pulumi.BoolPtrInput
	// Time in seconds to wait for any individual kubernetes operation.
	Timeout pulumi.IntPtrInput
	// List of assets (raw yaml files). Content is read and merged with values.
	ValueYamlFiles pulumi.AssetOrArchiveArrayInput
	// Custom values set for the release.
	Values pulumi.MapInput
	// Verify the package before installing it.
	Verify pulumi.BoolPtrInput
	// Specify the exact chart version to install. If this is not specified, the latest version is installed.
	Version pulumi.StringPtrInput
	// Will wait until all Jobs have been completed before marking the release as successful. This is ignored if `skipAwait` is enabled.
	WaitForJobs pulumi.BoolPtrInput
}

func (ReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseArgs)(nil)).Elem()
}

type ReleaseInput interface {
	pulumi.Input

	ToReleaseOutput() ReleaseOutput
	ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput
}

func (*Release) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (i *Release) ToReleaseOutput() ReleaseOutput {
	return i.ToReleaseOutputWithContext(context.Background())
}

func (i *Release) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseOutput)
}

// ReleaseArrayInput is an input type that accepts ReleaseArray and ReleaseArrayOutput values.
// You can construct a concrete instance of `ReleaseArrayInput` via:
//
//	ReleaseArray{ ReleaseArgs{...} }
type ReleaseArrayInput interface {
	pulumi.Input

	ToReleaseArrayOutput() ReleaseArrayOutput
	ToReleaseArrayOutputWithContext(context.Context) ReleaseArrayOutput
}

type ReleaseArray []ReleaseInput

func (ReleaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Release)(nil)).Elem()
}

func (i ReleaseArray) ToReleaseArrayOutput() ReleaseArrayOutput {
	return i.ToReleaseArrayOutputWithContext(context.Background())
}

func (i ReleaseArray) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseArrayOutput)
}

// ReleaseMapInput is an input type that accepts ReleaseMap and ReleaseMapOutput values.
// You can construct a concrete instance of `ReleaseMapInput` via:
//
//	ReleaseMap{ "key": ReleaseArgs{...} }
type ReleaseMapInput interface {
	pulumi.Input

	ToReleaseMapOutput() ReleaseMapOutput
	ToReleaseMapOutputWithContext(context.Context) ReleaseMapOutput
}

type ReleaseMap map[string]ReleaseInput

func (ReleaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Release)(nil)).Elem()
}

func (i ReleaseMap) ToReleaseMapOutput() ReleaseMapOutput {
	return i.ToReleaseMapOutputWithContext(context.Background())
}

func (i ReleaseMap) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseMapOutput)
}

type ReleaseOutput struct{ *pulumi.OutputState }

func (ReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (o ReleaseOutput) ToReleaseOutput() ReleaseOutput {
	return o
}

func (o ReleaseOutput) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return o
}

// Whether to allow Null values in helm chart configs.
func (o ReleaseOutput) AllowNullValues() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.AllowNullValues }).(pulumi.BoolPtrOutput)
}

// If set, installation process purges chart on fail. `skipAwait` will be disabled automatically if atomic is used.
func (o ReleaseOutput) Atomic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.Atomic }).(pulumi.BoolPtrOutput)
}

// Chart name to be installed. A path may be used.
func (o ReleaseOutput) Chart() pulumi.StringOutput {
	return o.ApplyT(func(v *Release) pulumi.StringOutput { return v.Chart }).(pulumi.StringOutput)
}

// Allow deletion of new resources created in this upgrade when upgrade fails.
func (o ReleaseOutput) CleanupOnFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.CleanupOnFail }).(pulumi.BoolPtrOutput)
}

// Create the namespace if it does not exist.
func (o ReleaseOutput) CreateNamespace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.CreateNamespace }).(pulumi.BoolPtrOutput)
}

// Run helm dependency update before installing the chart.
func (o ReleaseOutput) DependencyUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.DependencyUpdate }).(pulumi.BoolPtrOutput)
}

// Add a custom description
func (o ReleaseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Use chart development versions, too. Equivalent to version '>0.0.0-0'. If `version` is set, this is ignored.
func (o ReleaseOutput) Devel() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.Devel }).(pulumi.BoolPtrOutput)
}

// Prevent CRD hooks from running, but run other hooks.  See helm install --no-crd-hook
func (o ReleaseOutput) DisableCRDHooks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.DisableCRDHooks }).(pulumi.BoolPtrOutput)
}

// If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema
func (o ReleaseOutput) DisableOpenapiValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.DisableOpenapiValidation }).(pulumi.BoolPtrOutput)
}

// Prevent hooks from running.
func (o ReleaseOutput) DisableWebhooks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.DisableWebhooks }).(pulumi.BoolPtrOutput)
}

// Force resource update through delete/recreate if needed.
func (o ReleaseOutput) ForceUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.ForceUpdate }).(pulumi.BoolPtrOutput)
}

// Location of public keys used for verification. Used only if `verify` is true
func (o ReleaseOutput) Keyring() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.StringPtrOutput { return v.Keyring }).(pulumi.StringPtrOutput)
}

// Run helm lint when planning.
func (o ReleaseOutput) Lint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.Lint }).(pulumi.BoolPtrOutput)
}

// The rendered manifests as JSON. Not yet supported.
func (o ReleaseOutput) Manifest() pulumi.MapOutput {
	return o.ApplyT(func(v *Release) pulumi.MapOutput { return v.Manifest }).(pulumi.MapOutput)
}

// Limit the maximum number of revisions saved per release. Use 0 for no limit.
func (o ReleaseOutput) MaxHistory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.IntPtrOutput { return v.MaxHistory }).(pulumi.IntPtrOutput)
}

// Release name.
func (o ReleaseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Namespace to install the release into.
func (o ReleaseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Postrender command to run.
func (o ReleaseOutput) Postrender() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.StringPtrOutput { return v.Postrender }).(pulumi.StringPtrOutput)
}

// Perform pods restart during upgrade/rollback.
func (o ReleaseOutput) RecreatePods() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.RecreatePods }).(pulumi.BoolPtrOutput)
}

// If set, render subchart notes along with the parent.
func (o ReleaseOutput) RenderSubchartNotes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.RenderSubchartNotes }).(pulumi.BoolPtrOutput)
}

// Re-use the given name, even if that name is already used. This is unsafe in production
func (o ReleaseOutput) Replace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.Replace }).(pulumi.BoolPtrOutput)
}

// Specification defining the Helm chart repository to use.
func (o ReleaseOutput) RepositoryOpts() RepositoryOptsPtrOutput {
	return o.ApplyT(func(v *Release) RepositoryOptsPtrOutput { return v.RepositoryOpts }).(RepositoryOptsPtrOutput)
}

// When upgrading, reset the values to the ones built into the chart.
func (o ReleaseOutput) ResetValues() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.ResetValues }).(pulumi.BoolPtrOutput)
}

// Names of resources created by the release grouped by "kind/version".
func (o ReleaseOutput) ResourceNames() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *Release) pulumi.StringArrayMapOutput { return v.ResourceNames }).(pulumi.StringArrayMapOutput)
}

// When upgrading, reuse the last release's values and merge in any overrides. If 'resetValues' is specified, this is ignored
func (o ReleaseOutput) ReuseValues() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.ReuseValues }).(pulumi.BoolPtrOutput)
}

// By default, the provider waits until all resources are in a ready state before marking the release as successful. Setting this to true will skip such await logic.
func (o ReleaseOutput) SkipAwait() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.SkipAwait }).(pulumi.BoolPtrOutput)
}

// If set, no CRDs will be installed. By default, CRDs are installed if not already present.
func (o ReleaseOutput) SkipCrds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.SkipCrds }).(pulumi.BoolPtrOutput)
}

// Status of the deployed release.
func (o ReleaseOutput) Status() ReleaseStatusOutput {
	return o.ApplyT(func(v *Release) ReleaseStatusOutput { return v.Status }).(ReleaseStatusOutput)
}

// Time in seconds to wait for any individual kubernetes operation.
func (o ReleaseOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

// List of assets (raw yaml files). Content is read and merged with values (with values taking precedence).
func (o ReleaseOutput) ValueYamlFiles() pulumi.AssetOrArchiveArrayOutput {
	return o.ApplyT(func(v *Release) pulumi.AssetOrArchiveArrayOutput { return v.ValueYamlFiles }).(pulumi.AssetOrArchiveArrayOutput)
}

// Custom values set for the release.
func (o ReleaseOutput) Values() pulumi.MapOutput {
	return o.ApplyT(func(v *Release) pulumi.MapOutput { return v.Values }).(pulumi.MapOutput)
}

// Verify the package before installing it.
func (o ReleaseOutput) Verify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.Verify }).(pulumi.BoolPtrOutput)
}

// Specify the exact chart version to install. If this is not specified, the latest version is installed.
func (o ReleaseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

// Will wait until all Jobs have been completed before marking the release as successful. This is ignored if `skipAwait` is enabled.
func (o ReleaseOutput) WaitForJobs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Release) pulumi.BoolPtrOutput { return v.WaitForJobs }).(pulumi.BoolPtrOutput)
}

type ReleaseArrayOutput struct{ *pulumi.OutputState }

func (ReleaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Release)(nil)).Elem()
}

func (o ReleaseArrayOutput) ToReleaseArrayOutput() ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) Index(i pulumi.IntInput) ReleaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Release {
		return vs[0].([]*Release)[vs[1].(int)]
	}).(ReleaseOutput)
}

type ReleaseMapOutput struct{ *pulumi.OutputState }

func (ReleaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Release)(nil)).Elem()
}

func (o ReleaseMapOutput) ToReleaseMapOutput() ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) MapIndex(k pulumi.StringInput) ReleaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Release {
		return vs[0].(map[string]*Release)[vs[1].(string)]
	}).(ReleaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseInput)(nil)).Elem(), &Release{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseArrayInput)(nil)).Elem(), ReleaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseMapInput)(nil)).Elem(), ReleaseMap{})
	pulumi.RegisterOutputType(ReleaseOutput{})
	pulumi.RegisterOutputType(ReleaseArrayOutput{})
	pulumi.RegisterOutputType(ReleaseMapOutput{})
}