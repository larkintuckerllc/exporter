package exporter

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	coreV1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func initialCount(api coreV1.EndpointsInterface, service string) (int, string, error) {
	listOptions := metaV1.ListOptions{}
	endpointsList, err := api.List(listOptions)
	if err != nil {
		return 0, "", err
	}
	count := 0
	resourceVersion := endpointsList.ListMeta.ResourceVersion
	for _, endpoints := range endpointsList.Items {
		name := endpoints.ObjectMeta.Name
		if name != service {
			continue
		}
		for _, endpoint := range endpoints.Subsets {
			count += len(endpoint.Addresses)
		}
	}
	return count, resourceVersion, nil
}
