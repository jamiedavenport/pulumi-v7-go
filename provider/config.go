package provider

import (
	"context"
)

type Config struct {
	ApiKey string `pulumi:"apiKey" provider:"secret"`
}

func (c *Config) Configure(ctx context.Context) error {
	return nil
}
