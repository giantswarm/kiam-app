// +build k8srequired

package basic

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/giantswarm/appcatalog"
	e2esetup "github.com/giantswarm/e2esetup/chart"
	"github.com/giantswarm/e2esetup/chart/env"
	"github.com/giantswarm/e2etests/basicapp"
	"github.com/giantswarm/helmclient"
	"github.com/giantswarm/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	agentName          = "kiam-agent"
	certManagerAppName = "cert-manager-app"
	catalogURL         = "https://giantswarm.github.io/default-catalog"
	testCatalogURL     = "https://giantswarm.github.io/default-test-catalog"
	name               = "kiam"
	serverName         = "kiam-server"
)

var (
	ba         *basicapp.BasicApp
	helmClient *helmclient.Client
	k8sSetup   *k8sclient.Setup
	l          micrologger.Logger
	tarballURL string
)

func init() {
	ctx := context.Background()
	var err error

	var latestRelease string
	{
		latestRelease, err = appcatalog.GetLatestVersion(ctx, catalogURL, fmt.Sprintf("%s-app", name), "")
		if err != nil {
			panic(err.Error())
		}
	}

	{
		version := fmt.Sprintf("%s-%s", latestRelease, env.CircleSHA())
		tarballURL, err = appcatalog.NewTarballURL(testCatalogURL, fmt.Sprintf("%s-app", name), version)
		if err != nil {
			panic(err.Error())
		}
	}

	{
		c := micrologger.Config{}
		l, err = micrologger.New(c)
		if err != nil {
			panic(err.Error())
		}
	}

	var k8sClients *k8sclient.Clients
	{
		c := k8sclient.ClientsConfig{
			Logger: l,

			KubeConfigPath: env.KubeConfigPath(),
		}
		k8sClients, err = k8sclient.NewClients(c)
		if err != nil {
			panic(err.Error())
		}
	}

	{
		c := k8sclient.SetupConfig{
			Logger: l,

			Clients: k8sClients,
		}
		k8sSetup, err = k8sclient.NewSetup(c)
		if err != nil {
			panic(err.Error())
		}
	}

	{
		c := helmclient.Config{
			Logger:    l,
			K8sClient: k8sClients,
		}
		helmClient, err = helmclient.New(c)
		if err != nil {
			panic(err.Error())
		}
	}

	{
		c := basicapp.Config{
			Clients:    k8sClients,
			HelmClient: helmClient,
			Logger:     l,

			App: basicapp.Chart{
				Name:      name,
				Namespace: metav1.NamespaceSystem,
				URL:       tarballURL,
			},
			ChartResources: basicapp.ChartResources{
				DaemonSets: []basicapp.DaemonSet{
					{
						Name:      agentName,
						Namespace: metav1.NamespaceSystem,
						Labels: map[string]string{
							"app":                        name,
							"component":                  agentName,
							"giantswarm.io/service-type": "managed",
						},
						MatchLabels: map[string]string{
							"app":       name,
							"component": agentName,
						},
					},
					{
						Name:      serverName,
						Namespace: metav1.NamespaceSystem,
						Labels: map[string]string{
							"app":                        name,
							"component":                  serverName,
							"giantswarm.io/service-type": "managed",
						},
						MatchLabels: map[string]string{
							"app":       name,
							"component": serverName,
						},
					},
				},
			},
		}
		ba, err = basicapp.New(c)
		if err != nil {
			panic(err.Error())
		}
	}
}

// TestMain allows us to have common setup and teardown steps that are run
// once for all the tests https://golang.org/pkg/testing/#hdr-Main.
func TestMain(m *testing.M) {
	ctx := context.Background()

	{
		c := e2esetup.Config{
			HelmClient: helmClient,
			Setup:      k8sSetup,
		}

		v, err := e2esetup.Setup(ctx, m, c)
		if err != nil {
			l.LogCtx(ctx, "level", "error", "message", "e2e test failed", "stack", microerror.JSON(err))
		}

		os.Exit(v)
	}
}
