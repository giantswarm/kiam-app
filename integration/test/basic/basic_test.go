// +build k8srequired

package basic

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/giantswarm/appcatalog"
	"github.com/giantswarm/helmclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/helm/pkg/helm"
)

func TestHelm(t *testing.T) {
	ctx := context.Background()

	// Install cert-manager so kiam certs can be issued.
	err := installCertManager(ctx, helmClient, l)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	err = ba.Test(context.Background())
	if err != nil {
		t.Fatalf("%#v", err)
	}
}

func installCertManager(ctx context.Context, helmClient helmclient.Interface, logger micrologger.Logger) error {
	tarballURL, err := appcatalog.GetLatestChart(ctx, defaultCatalogURL, certManagerAppName)
	if err != nil {
		return microerror.Mask(err)
	}

	tarballPath, err := helmClient.PullChartTarball(ctx, tarballURL)
	if err != nil {
		return microerror.Mask(err)
	}

	defer func() {
		err := os.Remove(tarballPath)
		if err != nil {
			logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("deletion of %#q failed", tarballPath), "stack", microerror.Stack(err))
		}
	}()

	err = helmClient.InstallReleaseFromTarball(ctx, tarballPath, metav1.NamespaceSystem, helm.ReleaseName(certManagerAppName))
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
