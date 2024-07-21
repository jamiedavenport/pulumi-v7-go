package provider

import (
	"context"
	"fmt"
)

type Property struct{}

type PropertyInput struct {
	PropertyId string `pulumi:"propertyId"`
}

type PropertyArgs struct {
	WorkspaceId string `pulumi:"workspaceId" json:"-"`
	ProjectId   string `pulumi:"projectId" json:"-"`

	Name        string          `pulumi:"name"`
	Description string          `pulumi:"description"`
	Tool        string          `pulumi:"tool"` // todo enum
	Type        string          `pulumi:"type"`
	Inputs      []PropertyInput `pulumi:"inputs"`
}

type PropertyState struct {
	PropertyArgs

	PropertyId string `pulumi:"propertyId"`
}

type CreatePropertyInput struct {
	PropertyId string `json:"property_id"`
}

type CreatePropertyBody struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Tool        string                `json:"tool"` // todo enum
	Type        string                `json:"type"`
	Inputs      []CreatePropertyInput `json:"inputs"`
}

type CreatePropertyResult struct {
	Id string `json:"id"`
}

func (Property) Create(ctx context.Context, name string, args PropertyArgs, preview bool) (string, PropertyState, error) {
	if preview {
		return name, PropertyState{}, nil
	}

	var res CreatePropertyResult
	client := CreateClient(ctx)

	inputs := make([]CreatePropertyInput, len(args.Inputs))

	for i, input := range args.Inputs {
		inputs[i] = CreatePropertyInput(input)
	}

	resp, err := client.R().
		SetBody(&CreatePropertyBody{
			Name:        args.Name,
			Description: args.Description,
			Tool:        args.Tool,
			Type:        args.Type,
			Inputs:      inputs,
		}).
		SetSuccessResult(&res).
		Post(fmt.Sprintf("https://go.v7labs.com/api/workspaces/%s/projects/%s/properties", args.WorkspaceId, args.ProjectId))

	if err != nil {
		return name, PropertyState{}, err
	}

	if !resp.IsSuccessState() {
		return name, PropertyState{}, fmt.Errorf("failed to create property: %s", resp.String())
	}

	return name, PropertyState{
		PropertyArgs: args,
		PropertyId:   res.Id,
	}, nil
}

func (Property) Delete(ctx context.Context, id string, state PropertyState) error {
	client := CreateClient(ctx)

	resp, err := client.R().
		Delete(fmt.Sprintf("https://go.v7labs.com/api/workspaces/%s/projects/%s/properties/%s", state.WorkspaceId, state.ProjectId, state.PropertyId))
	if err != nil {
		return err
	}

	if !resp.IsSuccessState() {
		return fmt.Errorf("failed to delete project: %s", resp.String())
	}

	return nil
}
