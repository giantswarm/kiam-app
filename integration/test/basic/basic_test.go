// +build k8srequired

package basic

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/giantswarm/appcatalog"
	"github.com/giantswarm/helmclient/v2/pkg/helmclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	tarballURL, err := appcatalog.GetLatestChart(ctx, catalogURL, certManagerAppName, "")
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
			logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("deletion of %#q failed", tarballPath), "stack", microerror.JSON(err))
		}
	}()

	options := helmclient.InstallOptions{
		ReleaseName: certManagerAppName,
	}

	err = helmClient.InstallReleaseFromTarball(ctx, tarballPath, metav1.NamespaceSystem, map[string]interface{}{}, options)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
