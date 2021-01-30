package exporter

import (
	"errors"

	"k8s.io/api/autoscaling/v2beta2"
	"k8s.io/apimachinery/pkg/watch"
)

func updateReplicas(update watch.Event, hpa string, prevReplicas int32) (int32, error) {
	hpaObject, ok := update.Object.(*v2beta2.HorizontalPodAutoscaler)
	if !ok {
		err := errors.New("failed to cast update object to hpa")
		return 0, err
	}
	if hpaObject.ObjectMeta.Name != hpa {
		return prevReplicas, nil
	}
	replicas := hpaObject.Status.CurrentReplicas
	return replicas, nil
}
