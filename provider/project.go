package provider

import (
	"context"
	"fmt"

	"github.com/imroc/req/v3"

	"github.com/pulumi/pulumi-go-provider/infer"
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

	ProjectId string `pulumi:"projectId"`
}

type CreateProjectRequest struct {
	Name string `json:"name"`
}

type CreateProjectResponse struct {
	Id string `json:"id"`
}

func (Project) Create(ctx context.Context, name string, args ProjectArgs, preview bool) (string, ProjectState, error) {
	state := ProjectState{}
	if preview {
		return name, state, nil
	}

	config := infer.GetConfig[Config](ctx)
	client := req.C().DevMode()
	var response CreateProjectResponse

	resp, err := client.R().
		SetBody(&CreateProjectRequest{Name: args.Name}).
		SetSuccessResult(&response).
		SetHeader("x-api-key", config.ApiKey).
		Post(fmt.Sprintf("https://go.v7labs.com/api/workspaces/%s/projects", args.WorkspaceId))
	if err != nil {
		return Name, state, err
	}

	if !resp.IsSuccessState() {
		return name, state, fmt.Errorf("failed to create project: %s", resp.String())
	}

	state.ProjectId = response.Id

	return name, state, nil
}
