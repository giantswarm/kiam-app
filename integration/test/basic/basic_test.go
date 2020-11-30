// +build k8srequired

package basic

import (
	"context"
	"testing"
	"time"

	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestDaemonSets ensures that kiam-server and kiam-agent daemonset are created
func TestDaemonSets(t *testing.T) {
	ctx := context.Background()
	var err error

	err = checkDaemonsetExists(ctx)
	if err != nil {
		t.Fatalf("could not get kiam-agent & kiam-server: %v", err)
	}
}

func checkDaemonsetExists(ctx context.Context) error {
	var err error

	l.LogCtx(ctx, "level", "debug", "message", "waiting for daemonset create")

	o := func() error {
		_, err := appTest.K8sClient().AppsV1().DaemonSets(metav1.NamespaceSystem).Get(ctx, kiamServer, metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			return microerror.Maskf(executionFailedError, "daemonset %#q in %#q not found", kiamServer, metav1.NamespaceSystem)
		}

		_, err = appTest.K8sClient().AppsV1().DaemonSets(metav1.NamespaceSystem).Get(ctx, kiamAgent, metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			return microerror.Maskf(executionFailedError, "daemonset %#q in %#q not found", kiamAgent, metav1.NamespaceSystem)
		}

		return nil
	}
	b := backoff.NewConstant(2*time.Minute, 5*time.Second)
	n := backoff.NewNotifier(l, ctx)

	err = backoff.RetryNotify(o, b, n)
	if err != nil {
		return microerror.Mask(err)
	}

	l.LogCtx(ctx, "level", "debug", "message", "daemonsets are created")

	return nil
}
