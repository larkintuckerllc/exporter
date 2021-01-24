package exporter

import (
	"errors"

	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func updatePods(update watch.Event, prevPods []string) ([]string, error) {
	nextPods := append([]string{}, prevPods...)
	pod, ok := update.Object.(*coreV1.Pod)
	if !ok {
		err := errors.New("failed to cast update object to pod")
		return []string{}, err
	}
	name := pod.ObjectMeta.Name
	if update.Type == "ADDED" {
		nextPods = append(nextPods, name)
	}
	if update.Type == "DELETED" {
		i := index(nextPods, name)
		if i != -1 {
			nextPods[i] = nextPods[len(nextPods)-1]
			nextPods[len(nextPods)-1] = ""
			nextPods = nextPods[:len(nextPods)-1]
		}
	}
	return nextPods, nil
}

func index(slice []string, item string) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}
