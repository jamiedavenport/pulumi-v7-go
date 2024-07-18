package provider

import (
	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Project struct{}

type ProjectArgs struct{}

type ProjectState struct {
	ProjectArgs

	Result string `pulumi:"result"`
}

func (Project) Create(ctx context.Context, name string, args ProjectArgs, preview bool) (string, ProjectState, error) {
	state := ProjectState{}
	if preview {
		return name, state, nil
	}

	config := infer.GetConfig[Config](ctx)

	state.Result = config.ApiKey

	return name, state, nil
}
