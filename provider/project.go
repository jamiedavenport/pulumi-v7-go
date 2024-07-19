package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

type CreateProjectResponse struct {
	Id string `json:"id"`
}

func (Project) Create(ctx context.Context, name string, args ProjectArgs, preview bool) (string, ProjectState, error) {
	state := ProjectState{}
	if preview {
		return name, state, nil
	}

	config := infer.GetConfig[Config](ctx)

	body := map[string]string{
		"name": args.Name,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return name, state, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://go.v7labs.com/api/workspaces/%s/projects", args.WorkspaceId), bytes.NewBuffer(jsonBody))
	if err != nil {
		return name, state, err
	}
	req.Header.Set("x-api-key", config.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return name, state, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return name, state, fmt.Errorf("failed to create project: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return name, state, err
	}

	var createProjectResponse CreateProjectResponse
	if err := json.Unmarshal(data, &createProjectResponse); err != nil {
		return name, state, err
	}

	state.ProjectId = createProjectResponse.Id

	return name, state, nil
}
