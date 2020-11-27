// +build k8srequired

package basic

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestSecrets ensures that kiam-server and kiam-agent created
func TestSecrets(t *testing.T) {
	ctx := context.Background()
	var err error

	err = checkReadyDaemonset(ctx)
	if err != nil {
		t.Fatalf("could not get kiam-agent & kiam-server: %v", err)
	}
}

func checkReadyDaemonset(ctx context.Context) error {
	var err error

	l.LogCtx(ctx, "level", "debug", "message", "waiting for ready daemonset")

	o := func() error {
		lo := metav1.ListOptions{
			LabelSelector: fmt.Sprintf("%s=%s", "app.kubernetes.io/name", app),
		}
		dms, err := appTest.K8sClient().AppsV1().DaemonSets(metav1.NamespaceSystem).List(ctx, lo)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(dms.Items) != 2 {
			return microerror.Maskf(executionFailedError, "the number of kiam daemonsets in %#q should be equal to 2 but got %d", metav1.NamespaceSystem, len(dms.Items))
		}

		for _, ds := range dms.Items {
			if ds.Status.NumberReady != ds.Status.DesiredNumberScheduled {
				return microerror.Maskf(executionFailedError, "daemonset %#q want %d replicas %d ready", ds.Name, ds.Status.DesiredNumberScheduled, ds.Status.NumberReady)
			}
		}

		return nil
	}
	b := backoff.NewConstant(2*time.Minute, 5*time.Second)
	n := backoff.NewNotifier(l, ctx)

	err = backoff.RetryNotify(o, b, n)
	if err != nil {
		return microerror.Mask(err)
	}

	l.LogCtx(ctx, "level", "debug", "message", "deployment is ready")

	return nil
}
