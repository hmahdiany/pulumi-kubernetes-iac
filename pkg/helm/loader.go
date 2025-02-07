package helm

import (
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List of deployment functions
var deployments = []func(ctx *pulumi.Context) error{
	IngressNginx,
	// Add more deployment functions here
}

// Package-level variable for the charts directory
var chartsDirectory string

func init() {
	// Initialize the chartsDirectory variable
	chartsDirectory = filepath.Join(os.Getenv("PROJECT_ROOT"), "charts")
}

// DeployAllCharts runs all Helm chart deployments
func DeployAllCharts(ctx *pulumi.Context) error {
	for _, deploy := range deployments {
		if err := deploy(ctx); err != nil {
			return err
		}
	}
	return nil
}
