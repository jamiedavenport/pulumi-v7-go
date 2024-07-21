package main

import (
	v7Go "github.com/pulumi/pulumi-v7-go/sdk/go/v7-go"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myProject, err := v7Go.NewProject(ctx, "myProject", &v7Go.ProjectArgs{
			WorkspaceId: pulumi.String("01905686-d113-79ff-9b33-0fa7c3ff6664"),
			Name:        pulumi.String("My New Pulumi Project Updated"),
		})
		if err != nil {
			return err
		}
		myInput, err := v7Go.NewProperty(ctx, "myInput", &v7Go.PropertyArgs{
			WorkspaceId: myProject.WorkspaceId,
			ProjectId:   myProject.ProjectId,
			Name:        pulumi.String("My Input"),
			Tool:        pulumi.String("manual"),
			Type:        pulumi.String("text"),
			Inputs:      v7go.PropertyInputTypeArray{},
		})
		if err != nil {
			return err
		}
		_, err = v7Go.NewProperty(ctx, "myOutput", &v7Go.PropertyArgs{
			WorkspaceId: myProject.WorkspaceId,
			ProjectId:   myProject.ProjectId,
			Name:        pulumi.String("My Output"),
			Description: pulumi.String("Return the text in uppercase"),
			Tool:        pulumi.String("gpt_4o"),
			Type:        pulumi.String("text"),
			Inputs: v7go.PropertyInputTypeArray{
				&v7go.PropertyInputTypeArgs{
					PropertyId: myInput.PropertyId,
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("output", []interface{}{})
		return nil
	})
}
