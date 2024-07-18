package main

import (
	"github.com/jamiedavenport/pulumi-v7-go/sdk/go/v7-go"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := v7-go.NewRandom(ctx, "myRandomResource", &v7-go.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", map[string]interface{}{
			"value": myRandomResource.Result,
		})
		return nil
	})
}
