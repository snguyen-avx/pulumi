// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IAM access key. This is a set of credentials that allow API requests to be made as an IAM user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			lbUser, err := iam.NewUser(ctx, "lb", &iam.UserArgs{
//				Name: pulumi.String("loadbalancer"),
//				Path: pulumi.String("/system/"),
//			})
//			if err != nil {
//				return err
//			}
//			lb, err := iam.NewAccessKey(ctx, "lb", &iam.AccessKeyArgs{
//				User:   lbUser.Name,
//				PgpKey: pulumi.String("keybase:some_person_that_exists"),
//			})
//			if err != nil {
//				return err
//			}
//			lbRo, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"ec2:Describe*",
//						},
//						Resources: []string{
//							"*",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewUserPolicy(ctx, "lb_ro", &iam.UserPolicyArgs{
//				Name:   pulumi.String("test"),
//				User:   lbUser.Name,
//				Policy: pulumi.String(lbRo.Json),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("secret", lb.EncryptedSecret)
//			return nil
//		})
//	}
//
// ```
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := iam.NewUser(ctx, "test", &iam.UserArgs{
//				Name: pulumi.String("test"),
//				Path: pulumi.String("/test/"),
//			})
//			if err != nil {
//				return err
//			}
//			testAccessKey, err := iam.NewAccessKey(ctx, "test", &iam.AccessKeyArgs{
//				User: test.Name,
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("awsIamSmtpPasswordV4", testAccessKey.SesSmtpPasswordV4)
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import IAM Access Keys using the identifier. For example:
//
// ```sh
// $ pulumi import aws:iam/accessKey:AccessKey example AKIA1234567890
// ```
// Resource attributes such as `encrypted_secret`, `key_fingerprint`, `pgp_key`, `secret`, `ses_smtp_password_v4`, and `encrypted_ses_smtp_password_v4` are not available for imported resources as this information cannot be read from the IAM API.
type AccessKey struct {
	pulumi.CustomResourceState

	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the access key was created.
	CreateDate pulumi.StringOutput `pulumi:"createDate"`
	// Encrypted secret, base64 encoded, if `pgpKey` was specified. This attribute is not available for imported resources. The encrypted secret may be decrypted using the command line.
	EncryptedSecret pulumi.StringOutput `pulumi:"encryptedSecret"`
	// Encrypted SES SMTP password, base64 encoded, if `pgpKey` was specified. This attribute is not available for imported resources. The encrypted password may be decrypted using the command line.
	EncryptedSesSmtpPasswordV4 pulumi.StringOutput `pulumi:"encryptedSesSmtpPasswordV4"`
	// Fingerprint of the PGP key used to encrypt the secret. This attribute is not available for imported resources.
	KeyFingerprint pulumi.StringOutput `pulumi:"keyFingerprint"`
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`, for use in the `encryptedSecret` output attribute. If providing a base-64 encoded PGP public key, make sure to provide the "raw" version and not the "armored" one (e.g. avoid passing the `-a` option to `gpg --export`).
	PgpKey pulumi.StringPtrOutput `pulumi:"pgpKey"`
	// Secret access key. This attribute is not available for imported resources. Note that this will be written to the state file. If you use this, please protect your backend state file judiciously. Alternatively, you may supply a `pgpKey` instead, which will prevent the secret from being stored in plaintext, at the cost of preventing the use of the secret key in automation.
	Secret pulumi.StringOutput `pulumi:"secret"`
	// Secret access key converted into an SES SMTP password by applying [AWS's documented Sigv4 conversion algorithm](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/smtp-credentials.html#smtp-credentials-convert). This attribute is not available for imported resources. As SigV4 is region specific, valid Provider regions are `ap-south-1`, `ap-southeast-2`, `eu-central-1`, `eu-west-1`, `us-east-1` and `us-west-2`. See current [AWS SES regions](https://docs.aws.amazon.com/general/latest/gr/rande.html#ses_region).
	SesSmtpPasswordV4 pulumi.StringOutput `pulumi:"sesSmtpPasswordV4"`
	// Access key status to apply. Defaults to `Active`. Valid values are `Active` and `Inactive`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// IAM user to associate with this access key.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewAccessKey registers a new resource with the given unique name, arguments, and options.
func NewAccessKey(ctx *pulumi.Context,
	name string, args *AccessKeyArgs, opts ...pulumi.ResourceOption) (*AccessKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
		"sesSmtpPasswordV4",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessKey
	err := ctx.RegisterResource("aws:iam/accessKey:AccessKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessKey gets an existing AccessKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessKeyState, opts ...pulumi.ResourceOption) (*AccessKey, error) {
	var resource AccessKey
	err := ctx.ReadResource("aws:iam/accessKey:AccessKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessKey resources.
type accessKeyState struct {
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the access key was created.
	CreateDate *string `pulumi:"createDate"`
	// Encrypted secret, base64 encoded, if `pgpKey` was specified. This attribute is not available for imported resources. The encrypted secret may be decrypted using the command line.
	EncryptedSecret *string `pulumi:"encryptedSecret"`
	// Encrypted SES SMTP password, base64 encoded, if `pgpKey` was specified. This attribute is not available for imported resources. The encrypted password may be decrypted using the command line.
	EncryptedSesSmtpPasswordV4 *string `pulumi:"encryptedSesSmtpPasswordV4"`
	// Fingerprint of the PGP key used to encrypt the secret. This attribute is not available for imported resources.
	KeyFingerprint *string `pulumi:"keyFingerprint"`
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`, for use in the `encryptedSecret` output attribute. If providing a base-64 encoded PGP public key, make sure to provide the "raw" version and not the "armored" one (e.g. avoid passing the `-a` option to `gpg --export`).
	PgpKey *string `pulumi:"pgpKey"`
	// Secret access key. This attribute is not available for imported resources. Note that this will be written to the state file. If you use this, please protect your backend state file judiciously. Alternatively, you may supply a `pgpKey` instead, which will prevent the secret from being stored in plaintext, at the cost of preventing the use of the secret key in automation.
	Secret *string `pulumi:"secret"`
	// Secret access key converted into an SES SMTP password by applying [AWS's documented Sigv4 conversion algorithm](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/smtp-credentials.html#smtp-credentials-convert). This attribute is not available for imported resources. As SigV4 is region specific, valid Provider regions are `ap-south-1`, `ap-southeast-2`, `eu-central-1`, `eu-west-1`, `us-east-1` and `us-west-2`. See current [AWS SES regions](https://docs.aws.amazon.com/general/latest/gr/rande.html#ses_region).
	SesSmtpPasswordV4 *string `pulumi:"sesSmtpPasswordV4"`
	// Access key status to apply. Defaults to `Active`. Valid values are `Active` and `Inactive`.
	Status *string `pulumi:"status"`
	// IAM user to associate with this access key.
	User *string `pulumi:"user"`
}

type AccessKeyState struct {
	// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the access key was created.
	CreateDate pulumi.StringPtrInput
	// Encrypted secret, base64 encoded, if `pgpKey` was specified. This attribute is not available for imported resources. The encrypted secret may be decrypted using the command line.
	EncryptedSecret pulumi.StringPtrInput
	// Encrypted SES SMTP password, base64 encoded, if `pgpKey` was specified. This attribute is not available for imported resources. The encrypted password may be decrypted using the command line.
	EncryptedSesSmtpPasswordV4 pulumi.StringPtrInput
	// Fingerprint of the PGP key used to encrypt the secret. This attribute is not available for imported resources.
	KeyFingerprint pulumi.StringPtrInput
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`, for use in the `encryptedSecret` output attribute. If providing a base-64 encoded PGP public key, make sure to provide the "raw" version and not the "armored" one (e.g. avoid passing the `-a` option to `gpg --export`).
	PgpKey pulumi.StringPtrInput
	// Secret access key. This attribute is not available for imported resources. Note that this will be written to the state file. If you use this, please protect your backend state file judiciously. Alternatively, you may supply a `pgpKey` instead, which will prevent the secret from being stored in plaintext, at the cost of preventing the use of the secret key in automation.
	Secret pulumi.StringPtrInput
	// Secret access key converted into an SES SMTP password by applying [AWS's documented Sigv4 conversion algorithm](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/smtp-credentials.html#smtp-credentials-convert). This attribute is not available for imported resources. As SigV4 is region specific, valid Provider regions are `ap-south-1`, `ap-southeast-2`, `eu-central-1`, `eu-west-1`, `us-east-1` and `us-west-2`. See current [AWS SES regions](https://docs.aws.amazon.com/general/latest/gr/rande.html#ses_region).
	SesSmtpPasswordV4 pulumi.StringPtrInput
	// Access key status to apply. Defaults to `Active`. Valid values are `Active` and `Inactive`.
	Status pulumi.StringPtrInput
	// IAM user to associate with this access key.
	User pulumi.StringPtrInput
}

func (AccessKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessKeyState)(nil)).Elem()
}

type accessKeyArgs struct {
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`, for use in the `encryptedSecret` output attribute. If providing a base-64 encoded PGP public key, make sure to provide the "raw" version and not the "armored" one (e.g. avoid passing the `-a` option to `gpg --export`).
	PgpKey *string `pulumi:"pgpKey"`
	// Access key status to apply. Defaults to `Active`. Valid values are `Active` and `Inactive`.
	Status *string `pulumi:"status"`
	// IAM user to associate with this access key.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a AccessKey resource.
type AccessKeyArgs struct {
	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`, for use in the `encryptedSecret` output attribute. If providing a base-64 encoded PGP public key, make sure to provide the "raw" version and not the "armored" one (e.g. avoid passing the `-a` option to `gpg --export`).
	PgpKey pulumi.StringPtrInput
	// Access key status to apply. Defaults to `Active`. Valid values are `Active` and `Inactive`.
	Status pulumi.StringPtrInput
	// IAM user to associate with this access key.
	User pulumi.StringInput
}

func (AccessKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessKeyArgs)(nil)).Elem()
}

type AccessKeyInput interface {
	pulumi.Input

	ToAccessKeyOutput() AccessKeyOutput
	ToAccessKeyOutputWithContext(ctx context.Context) AccessKeyOutput
}

func (*AccessKey) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessKey)(nil)).Elem()
}

func (i *AccessKey) ToAccessKeyOutput() AccessKeyOutput {
	return i.ToAccessKeyOutputWithContext(context.Background())
}

func (i *AccessKey) ToAccessKeyOutputWithContext(ctx context.Context) AccessKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessKeyOutput)
}

// AccessKeyArrayInput is an input type that accepts AccessKeyArray and AccessKeyArrayOutput values.
// You can construct a concrete instance of `AccessKeyArrayInput` via:
//
//	AccessKeyArray{ AccessKeyArgs{...} }
type AccessKeyArrayInput interface {
	pulumi.Input

	ToAccessKeyArrayOutput() AccessKeyArrayOutput
	ToAccessKeyArrayOutputWithContext(context.Context) AccessKeyArrayOutput
}

type AccessKeyArray []AccessKeyInput

func (AccessKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessKey)(nil)).Elem()
}

func (i AccessKeyArray) ToAccessKeyArrayOutput() AccessKeyArrayOutput {
	return i.ToAccessKeyArrayOutputWithContext(context.Background())
}

func (i AccessKeyArray) ToAccessKeyArrayOutputWithContext(ctx context.Context) AccessKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessKeyArrayOutput)
}

// AccessKeyMapInput is an input type that accepts AccessKeyMap and AccessKeyMapOutput values.
// You can construct a concrete instance of `AccessKeyMapInput` via:
//
//	AccessKeyMap{ "key": AccessKeyArgs{...} }
type AccessKeyMapInput interface {
	pulumi.Input

	ToAccessKeyMapOutput() AccessKeyMapOutput
	ToAccessKeyMapOutputWithContext(context.Context) AccessKeyMapOutput
}

type AccessKeyMap map[string]AccessKeyInput

func (AccessKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessKey)(nil)).Elem()
}

func (i AccessKeyMap) ToAccessKeyMapOutput() AccessKeyMapOutput {
	return i.ToAccessKeyMapOutputWithContext(context.Background())
}

func (i AccessKeyMap) ToAccessKeyMapOutputWithContext(ctx context.Context) AccessKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessKeyMapOutput)
}

type AccessKeyOutput struct{ *pulumi.OutputState }

func (AccessKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessKey)(nil)).Elem()
}

func (o AccessKeyOutput) ToAccessKeyOutput() AccessKeyOutput {
	return o
}

func (o AccessKeyOutput) ToAccessKeyOutputWithContext(ctx context.Context) AccessKeyOutput {
	return o
}

// Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the access key was created.
func (o AccessKeyOutput) CreateDate() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessKey) pulumi.StringOutput { return v.CreateDate }).(pulumi.StringOutput)
}

// Encrypted secret, base64 encoded, if `pgpKey` was specified. This attribute is not available for imported resources. The encrypted secret may be decrypted using the command line.
func (o AccessKeyOutput) EncryptedSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessKey) pulumi.StringOutput { return v.EncryptedSecret }).(pulumi.StringOutput)
}

// Encrypted SES SMTP password, base64 encoded, if `pgpKey` was specified. This attribute is not available for imported resources. The encrypted password may be decrypted using the command line.
func (o AccessKeyOutput) EncryptedSesSmtpPasswordV4() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessKey) pulumi.StringOutput { return v.EncryptedSesSmtpPasswordV4 }).(pulumi.StringOutput)
}

// Fingerprint of the PGP key used to encrypt the secret. This attribute is not available for imported resources.
func (o AccessKeyOutput) KeyFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessKey) pulumi.StringOutput { return v.KeyFingerprint }).(pulumi.StringOutput)
}

// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`, for use in the `encryptedSecret` output attribute. If providing a base-64 encoded PGP public key, make sure to provide the "raw" version and not the "armored" one (e.g. avoid passing the `-a` option to `gpg --export`).
func (o AccessKeyOutput) PgpKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessKey) pulumi.StringPtrOutput { return v.PgpKey }).(pulumi.StringPtrOutput)
}

// Secret access key. This attribute is not available for imported resources. Note that this will be written to the state file. If you use this, please protect your backend state file judiciously. Alternatively, you may supply a `pgpKey` instead, which will prevent the secret from being stored in plaintext, at the cost of preventing the use of the secret key in automation.
func (o AccessKeyOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessKey) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

// Secret access key converted into an SES SMTP password by applying [AWS's documented Sigv4 conversion algorithm](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/smtp-credentials.html#smtp-credentials-convert). This attribute is not available for imported resources. As SigV4 is region specific, valid Provider regions are `ap-south-1`, `ap-southeast-2`, `eu-central-1`, `eu-west-1`, `us-east-1` and `us-west-2`. See current [AWS SES regions](https://docs.aws.amazon.com/general/latest/gr/rande.html#ses_region).
func (o AccessKeyOutput) SesSmtpPasswordV4() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessKey) pulumi.StringOutput { return v.SesSmtpPasswordV4 }).(pulumi.StringOutput)
}

// Access key status to apply. Defaults to `Active`. Valid values are `Active` and `Inactive`.
func (o AccessKeyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessKey) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// IAM user to associate with this access key.
func (o AccessKeyOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessKey) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

type AccessKeyArrayOutput struct{ *pulumi.OutputState }

func (AccessKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessKey)(nil)).Elem()
}

func (o AccessKeyArrayOutput) ToAccessKeyArrayOutput() AccessKeyArrayOutput {
	return o
}

func (o AccessKeyArrayOutput) ToAccessKeyArrayOutputWithContext(ctx context.Context) AccessKeyArrayOutput {
	return o
}

func (o AccessKeyArrayOutput) Index(i pulumi.IntInput) AccessKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessKey {
		return vs[0].([]*AccessKey)[vs[1].(int)]
	}).(AccessKeyOutput)
}

type AccessKeyMapOutput struct{ *pulumi.OutputState }

func (AccessKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessKey)(nil)).Elem()
}

func (o AccessKeyMapOutput) ToAccessKeyMapOutput() AccessKeyMapOutput {
	return o
}

func (o AccessKeyMapOutput) ToAccessKeyMapOutputWithContext(ctx context.Context) AccessKeyMapOutput {
	return o
}

func (o AccessKeyMapOutput) MapIndex(k pulumi.StringInput) AccessKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessKey {
		return vs[0].(map[string]*AccessKey)[vs[1].(string)]
	}).(AccessKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessKeyInput)(nil)).Elem(), &AccessKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessKeyArrayInput)(nil)).Elem(), AccessKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessKeyMapInput)(nil)).Elem(), AccessKeyMap{})
	pulumi.RegisterOutputType(AccessKeyOutput{})
	pulumi.RegisterOutputType(AccessKeyArrayOutput{})
	pulumi.RegisterOutputType(AccessKeyMapOutput{})
}
