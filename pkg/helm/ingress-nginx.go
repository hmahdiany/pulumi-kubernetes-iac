package helm

import (
	"os"

	"pulumi-kubernetes-iac/pkg/merge"

	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func IngressNginx(ctx *pulumi.Context) error {
	env := os.Getenv("PULUMI_ENV")
	valuesFile, err := merge.MergeValues("/home/hmahdiany/Documents/github/hmahdiany/pulumi-kubernetes-iac/charts/ingress-nginx/values.yaml", "/home/hmahdiany/Documents/github/hmahdiany/pulumi-kubernetes-iac/charts/ingress-nginx/overrides/"+env+".yaml")
	if err != nil {
		return err
	}

	_, err = helm.NewRelease(ctx, "ingress-nginx", &helm.ReleaseArgs{
		Name:      pulumi.String("ingress-nginx"),
		Namespace: pulumi.String("ingress-nginx"),
		Chart:     pulumi.String("https://kubernetes.github.io/ingress-nginx"),
		Version:   pulumi.String("3.15.2"),
		Values:    pulumi.ToMap(valuesFile),
	})
	return err
}
