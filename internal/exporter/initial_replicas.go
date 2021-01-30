package exporter

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
)

func initialReplicas(api v2beta1.HorizontalPodAutoscalerInterface, hpa string) (int32, string, error) {
	getOptions := metaV1.GetOptions{}
	hpaObject, err := api.Get(hpa, getOptions)
	if err != nil {
		return 0, "", err
	}
	resourceVersion := hpaObject.ObjectMeta.ResourceVersion
	replicas := hpaObject.Status.CurrentReplicas
	return replicas, resourceVersion, nil
}
