package exporter

import (
	coreV1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func initialCount(api coreV1.EndpointsInterface, service string) (int, string, error) {
	return 0, "", nil
}
