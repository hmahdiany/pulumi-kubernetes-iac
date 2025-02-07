package helm

import (
	"fmt"
	"os"
	"path/filepath"

	"pulumi-kubernetes-iac/pkg/merge"

	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func IngressNginx(ctx *pulumi.Context) error {
	env := os.Getenv("PULUMI_ENV")

	// Build the full paths dynamically
	baseFile := filepath.Join(chartsDirectory, "ingress-nginx", "values.yaml")
	overridesFile := filepath.Join(chartsDirectory, "ingress-nginx", "overrides", fmt.Sprintf("%s.yaml", env))

	valuesFile, err := merge.MergeValues(baseFile, overridesFile)
	if err != nil {
		return err
	}

	_, err = helm.NewRelease(ctx, "ingress-nginx", &helm.ReleaseArgs{
		Name:      pulumi.String("ingress-nginx"),
		Namespace: pulumi.String("ingress-nginx"),
		CreateNamespace: pulumi.BoolPtr(true),
		Chart:     pulumi.String("ingress-nginx"),
		Version:   pulumi.String("4.12.0"),
		Values:    pulumi.ToMap(valuesFile),
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://kubernetes.github.io/ingress-nginx"),
		},
	})
	return err
}
