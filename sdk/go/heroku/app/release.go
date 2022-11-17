// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Release struct {
	pulumi.CustomResourceState

	AppId       pulumi.StringOutput `pulumi:"appId"`
	Description pulumi.StringOutput `pulumi:"description"`
	SlugId      pulumi.StringOutput `pulumi:"slugId"`
}

// NewRelease registers a new resource with the given unique name, arguments, and options.
func NewRelease(ctx *pulumi.Context,
	name string, args *ReleaseArgs, opts ...pulumi.ResourceOption) (*Release, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.SlugId == nil {
		return nil, errors.New("invalid value for required argument 'SlugId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Release
	err := ctx.RegisterResource("heroku:app/release:Release", name, args, &resource, opts...)
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
	err := ctx.ReadResource("heroku:app/release:Release", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Release resources.
type releaseState struct {
	AppId       *string `pulumi:"appId"`
	Description *string `pulumi:"description"`
	SlugId      *string `pulumi:"slugId"`
}

type ReleaseState struct {
	AppId       pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	SlugId      pulumi.StringPtrInput
}

func (ReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseState)(nil)).Elem()
}

type releaseArgs struct {
	AppId       string  `pulumi:"appId"`
	Description *string `pulumi:"description"`
	SlugId      string  `pulumi:"slugId"`
}

// The set of arguments for constructing a Release resource.
type ReleaseArgs struct {
	AppId       pulumi.StringInput
	Description pulumi.StringPtrInput
	SlugId      pulumi.StringInput
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
//          ReleaseArray{ ReleaseArgs{...} }
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
//          ReleaseMap{ "key": ReleaseArgs{...} }
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

func (o ReleaseOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *Release) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o ReleaseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Release) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ReleaseOutput) SlugId() pulumi.StringOutput {
	return o.ApplyT(func(v *Release) pulumi.StringOutput { return v.SlugId }).(pulumi.StringOutput)
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