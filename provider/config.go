package provider

import (
	"context"

	p "github.com/pulumi/pulumi-go-provider"
)

type Config struct {
	ApiKey string `pulumi:"apiKey" provider:"secret"`
}

func (c *Config) Configure(ctx context.Context) error {
	p.GetLogger(ctx).Info("Hello!!")
	p.GetLogger(ctx).Infof("API Key: %s", c.ApiKey)

	return nil
}
