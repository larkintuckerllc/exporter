package exporter

import (
	"errors"

	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func updateCount(update watch.Event, service string, count int) (int, error) {
	endpoints, ok := update.Object.(*coreV1.Endpoints)
	if !ok {
		err := errors.New("failed to cast update object to endpoints")
		return count, err
	}
	name := endpoints.ObjectMeta.Name
	if name != service {
		return count, nil
	}
	count = 0
	for _, endpoint := range endpoints.Subsets {
		count += len(endpoint.Addresses)
	}
	return count, nil
}
