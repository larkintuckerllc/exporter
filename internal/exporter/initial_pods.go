package exporter

import (
	"fmt"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	coreV1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func initialPods(api coreV1.PodInterface, app string) ([]string, string, error) {
	labelSelector := fmt.Sprintf("app=%s", app)
	listOptions := metaV1.ListOptions{
		LabelSelector: labelSelector,
	}
	podList, err := api.List(listOptions)
	if err != nil {
		return []string{}, "", err
	}
	pods := []string{}
	resourceVersion := podList.ListMeta.ResourceVersion
	for _, pod := range podList.Items {
		name := pod.ObjectMeta.Name
		pods = append(pods, name)
	}
	return pods, resourceVersion, nil
}
