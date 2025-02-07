package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"pulumi-kubernetes-iac/pkg/helm"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return helm.DeployAllCharts(ctx)
	})
}
