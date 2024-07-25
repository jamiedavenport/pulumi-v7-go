package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

var Version string

const Name string = "v7-go"

func Provider() p.Provider {
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[Project, ProjectArgs, ProjectState](),
			infer.Resource[Property, PropertyArgs, PropertyState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
		Config: infer.Config[*Config](),
		Metadata: schema.Metadata{
			LanguageMap: map[string]any{
				"nodejs": nodejs.NodePackageInfo{
					PackageName: "@jamiedavenport/pulumi-v7-go",
				},
			},
			PluginDownloadURL: "github://api.github.com/jamiedavenport/pulumi-v7-go",
		},
	})
}
