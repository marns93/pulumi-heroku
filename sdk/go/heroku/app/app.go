// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type App struct {
	pulumi.CustomResourceState

	Acm                 pulumi.BoolOutput        `pulumi:"acm"`
	AllConfigVars       pulumi.MapOutput         `pulumi:"allConfigVars"`
	Buildpacks          pulumi.StringArrayOutput `pulumi:"buildpacks"`
	ConfigVars          pulumi.MapOutput         `pulumi:"configVars"`
	GitUrl              pulumi.StringOutput      `pulumi:"gitUrl"`
	HerokuHostname      pulumi.StringOutput      `pulumi:"herokuHostname"`
	InternalRouting     pulumi.BoolOutput        `pulumi:"internalRouting"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	Organization        AppOrganizationPtrOutput `pulumi:"organization"`
	Region              pulumi.StringOutput      `pulumi:"region"`
	SensitiveConfigVars pulumi.StringMapOutput   `pulumi:"sensitiveConfigVars"`
	Space               pulumi.StringPtrOutput   `pulumi:"space"`
	Stack               pulumi.StringOutput      `pulumi:"stack"`
	Uuid                pulumi.StringOutput      `pulumi:"uuid"`
	WebUrl              pulumi.StringOutput      `pulumi:"webUrl"`
}

// NewApp registers a new resource with the given unique name, arguments, and options.
func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.SensitiveConfigVars != nil {
		args.SensitiveConfigVars = pulumi.ToSecret(args.SensitiveConfigVars).(pulumi.StringMapInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"allConfigVars",
		"sensitiveConfigVars",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource App
	err := ctx.RegisterResource("heroku:app/app:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApp gets an existing App resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("heroku:app/app:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering App resources.
type appState struct {
	Acm                 *bool                  `pulumi:"acm"`
	AllConfigVars       map[string]interface{} `pulumi:"allConfigVars"`
	Buildpacks          []string               `pulumi:"buildpacks"`
	ConfigVars          map[string]interface{} `pulumi:"configVars"`
	GitUrl              *string                `pulumi:"gitUrl"`
	HerokuHostname      *string                `pulumi:"herokuHostname"`
	InternalRouting     *bool                  `pulumi:"internalRouting"`
	Name                *string                `pulumi:"name"`
	Organization        *AppOrganization       `pulumi:"organization"`
	Region              *string                `pulumi:"region"`
	SensitiveConfigVars map[string]string      `pulumi:"sensitiveConfigVars"`
	Space               *string                `pulumi:"space"`
	Stack               *string                `pulumi:"stack"`
	Uuid                *string                `pulumi:"uuid"`
	WebUrl              *string                `pulumi:"webUrl"`
}

type AppState struct {
	Acm                 pulumi.BoolPtrInput
	AllConfigVars       pulumi.MapInput
	Buildpacks          pulumi.StringArrayInput
	ConfigVars          pulumi.MapInput
	GitUrl              pulumi.StringPtrInput
	HerokuHostname      pulumi.StringPtrInput
	InternalRouting     pulumi.BoolPtrInput
	Name                pulumi.StringPtrInput
	Organization        AppOrganizationPtrInput
	Region              pulumi.StringPtrInput
	SensitiveConfigVars pulumi.StringMapInput
	Space               pulumi.StringPtrInput
	Stack               pulumi.StringPtrInput
	Uuid                pulumi.StringPtrInput
	WebUrl              pulumi.StringPtrInput
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	Acm                 *bool                  `pulumi:"acm"`
	Buildpacks          []string               `pulumi:"buildpacks"`
	ConfigVars          map[string]interface{} `pulumi:"configVars"`
	InternalRouting     *bool                  `pulumi:"internalRouting"`
	Name                *string                `pulumi:"name"`
	Organization        *AppOrganization       `pulumi:"organization"`
	Region              string                 `pulumi:"region"`
	SensitiveConfigVars map[string]string      `pulumi:"sensitiveConfigVars"`
	Space               *string                `pulumi:"space"`
	Stack               *string                `pulumi:"stack"`
}

// The set of arguments for constructing a App resource.
type AppArgs struct {
	Acm                 pulumi.BoolPtrInput
	Buildpacks          pulumi.StringArrayInput
	ConfigVars          pulumi.MapInput
	InternalRouting     pulumi.BoolPtrInput
	Name                pulumi.StringPtrInput
	Organization        AppOrganizationPtrInput
	Region              pulumi.StringInput
	SensitiveConfigVars pulumi.StringMapInput
	Space               pulumi.StringPtrInput
	Stack               pulumi.StringPtrInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (*App) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (i *App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i *App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

// AppArrayInput is an input type that accepts AppArray and AppArrayOutput values.
// You can construct a concrete instance of `AppArrayInput` via:
//
//          AppArray{ AppArgs{...} }
type AppArrayInput interface {
	pulumi.Input

	ToAppArrayOutput() AppArrayOutput
	ToAppArrayOutputWithContext(context.Context) AppArrayOutput
}

type AppArray []AppInput

func (AppArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*App)(nil)).Elem()
}

func (i AppArray) ToAppArrayOutput() AppArrayOutput {
	return i.ToAppArrayOutputWithContext(context.Background())
}

func (i AppArray) ToAppArrayOutputWithContext(ctx context.Context) AppArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppArrayOutput)
}

// AppMapInput is an input type that accepts AppMap and AppMapOutput values.
// You can construct a concrete instance of `AppMapInput` via:
//
//          AppMap{ "key": AppArgs{...} }
type AppMapInput interface {
	pulumi.Input

	ToAppMapOutput() AppMapOutput
	ToAppMapOutputWithContext(context.Context) AppMapOutput
}

type AppMap map[string]AppInput

func (AppMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*App)(nil)).Elem()
}

func (i AppMap) ToAppMapOutput() AppMapOutput {
	return i.ToAppMapOutputWithContext(context.Background())
}

func (i AppMap) ToAppMapOutputWithContext(ctx context.Context) AppMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppMapOutput)
}

type AppOutput struct{ *pulumi.OutputState }

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**App)(nil)).Elem()
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

func (o AppOutput) Acm() pulumi.BoolOutput {
	return o.ApplyT(func(v *App) pulumi.BoolOutput { return v.Acm }).(pulumi.BoolOutput)
}

func (o AppOutput) AllConfigVars() pulumi.MapOutput {
	return o.ApplyT(func(v *App) pulumi.MapOutput { return v.AllConfigVars }).(pulumi.MapOutput)
}

func (o AppOutput) Buildpacks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *App) pulumi.StringArrayOutput { return v.Buildpacks }).(pulumi.StringArrayOutput)
}

func (o AppOutput) ConfigVars() pulumi.MapOutput {
	return o.ApplyT(func(v *App) pulumi.MapOutput { return v.ConfigVars }).(pulumi.MapOutput)
}

func (o AppOutput) GitUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.GitUrl }).(pulumi.StringOutput)
}

func (o AppOutput) HerokuHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.HerokuHostname }).(pulumi.StringOutput)
}

func (o AppOutput) InternalRouting() pulumi.BoolOutput {
	return o.ApplyT(func(v *App) pulumi.BoolOutput { return v.InternalRouting }).(pulumi.BoolOutput)
}

func (o AppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppOutput) Organization() AppOrganizationPtrOutput {
	return o.ApplyT(func(v *App) AppOrganizationPtrOutput { return v.Organization }).(AppOrganizationPtrOutput)
}

func (o AppOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o AppOutput) SensitiveConfigVars() pulumi.StringMapOutput {
	return o.ApplyT(func(v *App) pulumi.StringMapOutput { return v.SensitiveConfigVars }).(pulumi.StringMapOutput)
}

func (o AppOutput) Space() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *App) pulumi.StringPtrOutput { return v.Space }).(pulumi.StringPtrOutput)
}

func (o AppOutput) Stack() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Stack }).(pulumi.StringOutput)
}

func (o AppOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o AppOutput) WebUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *App) pulumi.StringOutput { return v.WebUrl }).(pulumi.StringOutput)
}

type AppArrayOutput struct{ *pulumi.OutputState }

func (AppArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*App)(nil)).Elem()
}

func (o AppArrayOutput) ToAppArrayOutput() AppArrayOutput {
	return o
}

func (o AppArrayOutput) ToAppArrayOutputWithContext(ctx context.Context) AppArrayOutput {
	return o
}

func (o AppArrayOutput) Index(i pulumi.IntInput) AppOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *App {
		return vs[0].([]*App)[vs[1].(int)]
	}).(AppOutput)
}

type AppMapOutput struct{ *pulumi.OutputState }

func (AppMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*App)(nil)).Elem()
}

func (o AppMapOutput) ToAppMapOutput() AppMapOutput {
	return o
}

func (o AppMapOutput) ToAppMapOutputWithContext(ctx context.Context) AppMapOutput {
	return o
}

func (o AppMapOutput) MapIndex(k pulumi.StringInput) AppOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *App {
		return vs[0].(map[string]*App)[vs[1].(string)]
	}).(AppOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppInput)(nil)).Elem(), &App{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppArrayInput)(nil)).Elem(), AppArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppMapInput)(nil)).Elem(), AppMap{})
	pulumi.RegisterOutputType(AppOutput{})
	pulumi.RegisterOutputType(AppArrayOutput{})
	pulumi.RegisterOutputType(AppMapOutput{})
}
