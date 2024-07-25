# `pulumi-v7-go`

This Pulumi provider enables you to manage [V7 Go](https://www.v7labs.com/go) resources as code.

# Installation

## Node ([example](./examples/nodejs))

```sh
npm install @jamiedavenport/pulumi-v7-go
```

## SST

```sh
sst add @jamiedavenport/pulumi-v7-go
```

## Go ([example](./examples/go))

```sh
go get github.com/jamiedavenport/pulumi-v7-go/sdk/go/v7go
```

## YAML ([example](./examples/yaml/Pulumi.yaml))

```sh
pulumi plugin install resource v7-go --server github://api.github.com/jamiedavenport/pulumi-v7-go
```

## Python ([example](./examples/python))

TODO

## .NET ([example](./examples/dotnet))

```sh
dotnet add package JamieDavenport.V7Go
```

# Configuration

- `v7-go:apiKey` - (Required) The API key to use for authentication.

# Background

We've built this provider for use at Habitable. It gives us a number of benefits:

1. Manage V7 Go resources as code
1. Create multiple identical environments (e.g. staging, production)
1. Review changes to V7 Go resources in pull requests
1. Full type safety when using SST