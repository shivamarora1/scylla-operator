package scyllacluster

import (
	"context"
	"fmt"

	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
	"github.com/scylladb/scylla-operator/pkg/controllerhelpers"
	"github.com/scylladb/scylla-operator/pkg/resourceapply"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

func (scc *Controller) syncIngresses(
	ctx context.Context,
	sc *scyllav1.ScyllaCluster,
	ingresses map[string]*networkingv1.Ingress,
	services map[string]*corev1.Service,
) ([]metav1.Condition, error) {
	var err error
	var progressingConditions []metav1.Condition

	requiredIngresses := MakeIngresses(sc, services)

	// Delete any excessive Ingresses.
	// Delete has to be the fist action to avoid getting stuck on quota.
	var deletionErrors []error
	for _, ingress := range ingresses {
		if ingress.DeletionTimestamp != nil {
			continue
		}

		isRequired := false
		for _, req := range requiredIngresses {
			if ingress.Name == req.Name {
				isRequired = true
			}
		}
		if isRequired {
			continue
		}

		propagationPolicy := metav1.DeletePropagationBackground
		controllerhelpers.AddGenericProgressingStatusCondition(&progressingConditions, ingressControllerProgressingCondition, ingress, "delete", sc.Generation)
		err = scc.kubeClient.NetworkingV1().Ingresses(ingress.Namespace).Delete(ctx, ingress.Name, metav1.DeleteOptions{
			Preconditions: &metav1.Preconditions{
				UID: &ingress.UID,
			},
			PropagationPolicy: &propagationPolicy,
		})
		deletionErrors = append(deletionErrors, err)
	}
	err = utilerrors.NewAggregate(deletionErrors)
	if err != nil {
		return progressingConditions, fmt.Errorf("can't delete ingress(s): %w", err)
	}

	for _, requiredIngress := range requiredIngresses {
		_, changed, err := resourceapply.ApplyIngress(ctx, scc.kubeClient.NetworkingV1(), scc.ingressLister, scc.eventRecorder, requiredIngress)
		if changed {
			controllerhelpers.AddGenericProgressingStatusCondition(&progressingConditions, ingressControllerProgressingCondition, requiredIngress, "apply", sc.Generation)
		}
		if err != nil {
			return progressingConditions, fmt.Errorf("can't apply ingress: %w", err)
		}
	}

	return progressingConditions, nil
}
