package provider

import (
	"github.com/imroc/req/v3"

	"context"

	"github.com/pulumi/pulumi-go-provider/infer"
)

func CreateClient(ctx context.Context) *req.Client {
	config := infer.GetConfig[Config](ctx)
	client := req.C().SetCommonHeader("x-api-key", config.ApiKey)

	return client
}
