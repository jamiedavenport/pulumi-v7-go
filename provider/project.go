package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
)

type Project struct{}

type ProjectArgs struct{}

type ProjectState struct {
	ProjectArgs

	Result string `pulumi:"result"`
}

func (Project) Create(ctx p.Context, name string, args ProjectArgs, preview bool) (string, ProjectState, error) {
	state := ProjectState{}
	if preview {
		return name, state, nil
	}

	state.Result = "Hello, World!"

	return name, state, nil
}
