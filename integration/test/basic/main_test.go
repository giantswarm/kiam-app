//go:build k8srequired
// +build k8srequired

package basic

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/giantswarm/apptest"
	"github.com/giantswarm/micrologger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/kiam-app/integration/env"
)

const (
	app                = "kiam"
	appName            = "kiam-app"
	certManagerName    = "cert-manager-app"
	certManagerVersion = "2.3.3"
	defaultCatalog     = "giantswarm"
	defaultTestCatalog = "giantswarm-test"
	kiamAgent          = "kiam-agent"
	kiamServer         = "kiam-server"
)

var (
	appTest apptest.Interface
	l       micrologger.Logger
)

func init() {
	var err error

	{
		c := micrologger.Config{}
		l, err = micrologger.New(c)
		if err != nil {
			panic(err.Error())
		}
	}

	{
		c := apptest.Config{
			Logger: l,

			KubeConfigPath: env.KubeConfig(),
		}
		appTest, err = apptest.New(c)
		if err != nil {
			panic(err.Error())
		}
	}
}

// TestMain allows us to have common setup and teardown steps that are run
// once for all the tests https://golang.org/pkg/testing/#hdr-Main.
func TestMain(m *testing.M) {
	ctx := context.Background()

	var err error

	{
		apps := []apptest.App{
			{
				CatalogName:   defaultCatalog,
				Name:          certManagerName,
				Namespace:     metav1.NamespaceSystem,
				Version:       certManagerVersion,
				WaitForDeploy: true,
			},
			{
				CatalogName:   defaultTestCatalog,
				Name:          appName,
				Namespace:     metav1.NamespaceSystem,
				SHA:           env.CircleSHA(),
				WaitForDeploy: true,
			},
		}
		err = appTest.InstallApps(ctx, apps)
		if err != nil {
			l.LogCtx(ctx, "level", "error", "message", "install apps failed", "stack", fmt.Sprintf("%#v\n", err))
			os.Exit(-1)
		}
	}

	os.Exit(m.Run())
}
