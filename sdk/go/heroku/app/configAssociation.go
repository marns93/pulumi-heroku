// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigAssociation struct {
	pulumi.CustomResourceState

	AppId         pulumi.StringOutput    `pulumi:"appId"`
	SensitiveVars pulumi.StringMapOutput `pulumi:"sensitiveVars"`
	Vars          pulumi.StringMapOutput `pulumi:"vars"`
}

// NewConfigAssociation registers a new resource with the given unique name, arguments, and options.
func NewConfigAssociation(ctx *pulumi.Context,
	name string, args *ConfigAssociationArgs, opts ...pulumi.ResourceOption) (*ConfigAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.SensitiveVars != nil {
		args.SensitiveVars = pulumi.ToSecret(args.SensitiveVars).(pulumi.StringMapInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"sensitiveVars",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource ConfigAssociation
	err := ctx.RegisterResource("heroku:app/configAssociation:ConfigAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigAssociation gets an existing ConfigAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigAssociationState, opts ...pulumi.ResourceOption) (*ConfigAssociation, error) {
	var resource ConfigAssociation
	err := ctx.ReadResource("heroku:app/configAssociation:ConfigAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigAssociation resources.
type configAssociationState struct {
	AppId         *string           `pulumi:"appId"`
	SensitiveVars map[string]string `pulumi:"sensitiveVars"`
	Vars          map[string]string `pulumi:"vars"`
}

type ConfigAssociationState struct {
	AppId         pulumi.StringPtrInput
	SensitiveVars pulumi.StringMapInput
	Vars          pulumi.StringMapInput
}

func (ConfigAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configAssociationState)(nil)).Elem()
}

type configAssociationArgs struct {
	AppId         string            `pulumi:"appId"`
	SensitiveVars map[string]string `pulumi:"sensitiveVars"`
	Vars          map[string]string `pulumi:"vars"`
}

// The set of arguments for constructing a ConfigAssociation resource.
type ConfigAssociationArgs struct {
	AppId         pulumi.StringInput
	SensitiveVars pulumi.StringMapInput
	Vars          pulumi.StringMapInput
}

func (ConfigAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configAssociationArgs)(nil)).Elem()
}

type ConfigAssociationInput interface {
	pulumi.Input

	ToConfigAssociationOutput() ConfigAssociationOutput
	ToConfigAssociationOutputWithContext(ctx context.Context) ConfigAssociationOutput
}

func (*ConfigAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigAssociation)(nil)).Elem()
}

func (i *ConfigAssociation) ToConfigAssociationOutput() ConfigAssociationOutput {
	return i.ToConfigAssociationOutputWithContext(context.Background())
}

func (i *ConfigAssociation) ToConfigAssociationOutputWithContext(ctx context.Context) ConfigAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigAssociationOutput)
}

// ConfigAssociationArrayInput is an input type that accepts ConfigAssociationArray and ConfigAssociationArrayOutput values.
// You can construct a concrete instance of `ConfigAssociationArrayInput` via:
//
//          ConfigAssociationArray{ ConfigAssociationArgs{...} }
type ConfigAssociationArrayInput interface {
	pulumi.Input

	ToConfigAssociationArrayOutput() ConfigAssociationArrayOutput
	ToConfigAssociationArrayOutputWithContext(context.Context) ConfigAssociationArrayOutput
}

type ConfigAssociationArray []ConfigAssociationInput

func (ConfigAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigAssociation)(nil)).Elem()
}

func (i ConfigAssociationArray) ToConfigAssociationArrayOutput() ConfigAssociationArrayOutput {
	return i.ToConfigAssociationArrayOutputWithContext(context.Background())
}

func (i ConfigAssociationArray) ToConfigAssociationArrayOutputWithContext(ctx context.Context) ConfigAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigAssociationArrayOutput)
}

// ConfigAssociationMapInput is an input type that accepts ConfigAssociationMap and ConfigAssociationMapOutput values.
// You can construct a concrete instance of `ConfigAssociationMapInput` via:
//
//          ConfigAssociationMap{ "key": ConfigAssociationArgs{...} }
type ConfigAssociationMapInput interface {
	pulumi.Input

	ToConfigAssociationMapOutput() ConfigAssociationMapOutput
	ToConfigAssociationMapOutputWithContext(context.Context) ConfigAssociationMapOutput
}

type ConfigAssociationMap map[string]ConfigAssociationInput

func (ConfigAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigAssociation)(nil)).Elem()
}

func (i ConfigAssociationMap) ToConfigAssociationMapOutput() ConfigAssociationMapOutput {
	return i.ToConfigAssociationMapOutputWithContext(context.Background())
}

func (i ConfigAssociationMap) ToConfigAssociationMapOutputWithContext(ctx context.Context) ConfigAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigAssociationMapOutput)
}

type ConfigAssociationOutput struct{ *pulumi.OutputState }

func (ConfigAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigAssociation)(nil)).Elem()
}

func (o ConfigAssociationOutput) ToConfigAssociationOutput() ConfigAssociationOutput {
	return o
}

func (o ConfigAssociationOutput) ToConfigAssociationOutputWithContext(ctx context.Context) ConfigAssociationOutput {
	return o
}

func (o ConfigAssociationOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigAssociation) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o ConfigAssociationOutput) SensitiveVars() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigAssociation) pulumi.StringMapOutput { return v.SensitiveVars }).(pulumi.StringMapOutput)
}

func (o ConfigAssociationOutput) Vars() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigAssociation) pulumi.StringMapOutput { return v.Vars }).(pulumi.StringMapOutput)
}

type ConfigAssociationArrayOutput struct{ *pulumi.OutputState }

func (ConfigAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigAssociation)(nil)).Elem()
}

func (o ConfigAssociationArrayOutput) ToConfigAssociationArrayOutput() ConfigAssociationArrayOutput {
	return o
}

func (o ConfigAssociationArrayOutput) ToConfigAssociationArrayOutputWithContext(ctx context.Context) ConfigAssociationArrayOutput {
	return o
}

func (o ConfigAssociationArrayOutput) Index(i pulumi.IntInput) ConfigAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigAssociation {
		return vs[0].([]*ConfigAssociation)[vs[1].(int)]
	}).(ConfigAssociationOutput)
}

type ConfigAssociationMapOutput struct{ *pulumi.OutputState }

func (ConfigAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigAssociation)(nil)).Elem()
}

func (o ConfigAssociationMapOutput) ToConfigAssociationMapOutput() ConfigAssociationMapOutput {
	return o
}

func (o ConfigAssociationMapOutput) ToConfigAssociationMapOutputWithContext(ctx context.Context) ConfigAssociationMapOutput {
	return o
}

func (o ConfigAssociationMapOutput) MapIndex(k pulumi.StringInput) ConfigAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigAssociation {
		return vs[0].(map[string]*ConfigAssociation)[vs[1].(string)]
	}).(ConfigAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigAssociationInput)(nil)).Elem(), &ConfigAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigAssociationArrayInput)(nil)).Elem(), ConfigAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigAssociationMapInput)(nil)).Elem(), ConfigAssociationMap{})
	pulumi.RegisterOutputType(ConfigAssociationOutput{})
	pulumi.RegisterOutputType(ConfigAssociationArrayOutput{})
	pulumi.RegisterOutputType(ConfigAssociationMapOutput{})
}
