// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v7go

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-v7-go/sdk/go/v7-go/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Property struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput       `pulumi:"description"`
	Inputs      PropertyInputTypeArrayOutput `pulumi:"inputs"`
	Name        pulumi.StringOutput          `pulumi:"name"`
	ProjectId   pulumi.StringOutput          `pulumi:"projectId"`
	PropertyId  pulumi.StringOutput          `pulumi:"propertyId"`
	Tool        pulumi.StringOutput          `pulumi:"tool"`
	Type        pulumi.StringOutput          `pulumi:"type"`
	WorkspaceId pulumi.StringOutput          `pulumi:"workspaceId"`
}

// NewProperty registers a new resource with the given unique name, arguments, and options.
func NewProperty(ctx *pulumi.Context,
	name string, args *PropertyArgs, opts ...pulumi.ResourceOption) (*Property, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Inputs == nil {
		return nil, errors.New("invalid value for required argument 'Inputs'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Tool == nil {
		return nil, errors.New("invalid value for required argument 'Tool'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Property
	err := ctx.RegisterResource("v7-go:index:Property", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProperty gets an existing Property resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProperty(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertyState, opts ...pulumi.ResourceOption) (*Property, error) {
	var resource Property
	err := ctx.ReadResource("v7-go:index:Property", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Property resources.
type propertyState struct {
}

type PropertyState struct {
}

func (PropertyState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyState)(nil)).Elem()
}

type propertyArgs struct {
	Description *string             `pulumi:"description"`
	Inputs      []PropertyInputType `pulumi:"inputs"`
	Name        string              `pulumi:"name"`
	ProjectId   string              `pulumi:"projectId"`
	Tool        string              `pulumi:"tool"`
	Type        string              `pulumi:"type"`
	WorkspaceId string              `pulumi:"workspaceId"`
}

// The set of arguments for constructing a Property resource.
type PropertyArgs struct {
	Description pulumi.StringPtrInput
	Inputs      PropertyInputTypeArrayInput
	Name        pulumi.StringInput
	ProjectId   pulumi.StringInput
	Tool        pulumi.StringInput
	Type        pulumi.StringInput
	WorkspaceId pulumi.StringInput
}

func (PropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyArgs)(nil)).Elem()
}

type PropertyInput interface {
	pulumi.Input

	ToPropertyOutput() PropertyOutput
	ToPropertyOutputWithContext(ctx context.Context) PropertyOutput
}

func (*Property) ElementType() reflect.Type {
	return reflect.TypeOf((**Property)(nil)).Elem()
}

func (i *Property) ToPropertyOutput() PropertyOutput {
	return i.ToPropertyOutputWithContext(context.Background())
}

func (i *Property) ToPropertyOutputWithContext(ctx context.Context) PropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyOutput)
}

type PropertyOutput struct{ *pulumi.OutputState }

func (PropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Property)(nil)).Elem()
}

func (o PropertyOutput) ToPropertyOutput() PropertyOutput {
	return o
}

func (o PropertyOutput) ToPropertyOutputWithContext(ctx context.Context) PropertyOutput {
	return o
}

func (o PropertyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Property) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PropertyOutput) Inputs() PropertyInputTypeArrayOutput {
	return o.ApplyT(func(v *Property) PropertyInputTypeArrayOutput { return v.Inputs }).(PropertyInputTypeArrayOutput)
}

func (o PropertyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Property) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PropertyOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Property) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

func (o PropertyOutput) PropertyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Property) pulumi.StringOutput { return v.PropertyId }).(pulumi.StringOutput)
}

func (o PropertyOutput) Tool() pulumi.StringOutput {
	return o.ApplyT(func(v *Property) pulumi.StringOutput { return v.Tool }).(pulumi.StringOutput)
}

func (o PropertyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Property) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PropertyOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Property) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyInput)(nil)).Elem(), &Property{})
	pulumi.RegisterOutputType(PropertyOutput{})
}
