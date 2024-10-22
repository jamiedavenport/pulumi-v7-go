package provider

import (
	"context"
	"fmt"

	p "github.com/pulumi/pulumi-go-provider"
)

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type Project struct{}

type ProjectArgs struct {
	WorkspaceId string `pulumi:"workspaceId"`
	Name        string `pulumi:"name"`
}

type ProjectState struct {
	ProjectArgs

	ProjectId  string `pulumi:"projectId"`
	MainViewId string `pulumi:"mainViewId"`
}

type CreateProjectRequest struct {
	Name string `json:"name"`
}

type CreateProjectResponse struct {
	Id         string `json:"id"`
	MainViewId string `json:"main_view_id"`
}

func (Project) Create(ctx context.Context, name string, args ProjectArgs, preview bool) (string, ProjectState, error) {
	logger := p.GetLogger(ctx)
	logger.Info("creating project")

	state := ProjectState{}
	if preview {
		return name, state, nil
	}

	client := CreateClient(ctx)
	var response CreateProjectResponse

	resp, err := client.R().
		SetBody(&CreateProjectRequest{Name: args.Name}).
		SetSuccessResult(&response).
		Post(fmt.Sprintf("https://go.v7labs.com/api/workspaces/%s/projects", args.WorkspaceId))
	if err != nil {
		return Name, state, err
	}

	if !resp.IsSuccessState() {
		return name, state, fmt.Errorf("failed to create project: %s", resp.String())
	}

	return name, ProjectState{
		ProjectArgs: args,
		ProjectId:   response.Id,
		MainViewId:  response.MainViewId,
	}, nil
}

func (Project) Delete(ctx context.Context, id string, state ProjectState) error {
	logger := p.GetLogger(ctx)
	logger.Info("deleting project")

	client := CreateClient(ctx)
	resp, err := client.R().
		Delete(fmt.Sprintf("https://go.v7labs.com/api/workspaces/%s/projects/%s", state.WorkspaceId, state.ProjectId))
	if err != nil {
		return err
	}

	if !resp.IsSuccessState() {
		return fmt.Errorf("failed to delete project: %s", resp.String())
	}

	return nil
}

func (Project) Update(ctx context.Context, id string, oldState ProjectState, newArgs ProjectArgs, preview bool) (ProjectState, error) {
	logger := p.GetLogger(ctx)
	logger.Info("updating project")

	if preview {
		return ProjectState{}, nil
	}

	client := CreateClient(ctx)
	var response CreateProjectResponse

	resp, err := client.R().
		SetBody(&CreateProjectRequest{Name: newArgs.Name}).
		SetSuccessResult(&response).
		Put(fmt.Sprintf("https://go.v7labs.com/api/workspaces/%s/projects/%s", newArgs.WorkspaceId, oldState.ProjectId))
	if err != nil {
		return ProjectState{}, err
	}

	if !resp.IsSuccessState() {
		return ProjectState{}, fmt.Errorf("failed to update project: %s", resp.String())
	}

	return ProjectState{
		ProjectArgs: newArgs,
		ProjectId:   response.Id,
		MainViewId:  response.MainViewId,
	}, nil
}

func (Project) Diff(ctx context.Context, id string, oldState ProjectState, newArgs ProjectArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	logger := p.GetLogger(ctx)
	logger.Info("diffing project")

	if newArgs.Name != oldState.Name {
		diff["name"] = p.PropertyDiff{Kind: p.Update}
		logger.Info("name changed")
	}

	if newArgs.WorkspaceId != oldState.WorkspaceId {
		diff["workspaceId"] = p.PropertyDiff{Kind: p.UpdateReplace}
		logger.Info("workspaceId changed")
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}
