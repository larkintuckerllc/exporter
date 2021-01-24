package exporter

import (
	"time"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Execute(namespace string, service string, start int, end int, minimum int,
	project string, development bool) error {
	clientset, err := k8sAuth(development)
	if err != nil {
		return err
	}
	api := clientset.CoreV1().Endpoints(namespace)
	count, resourceVersion, err := initialCount(api, service)
	if err != nil {
		return err
	}
	listOptions := metaV1.ListOptions{
		ResourceVersion: resourceVersion,
	}
	watcher, err := api.Watch(listOptions)
	if err != nil {
		return err
	}
	defer watcher.Stop()
	updater := watcher.ResultChan()
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case update := <-updater:
			count, err = updateCount(update, service, count)
			if err != nil {
				return err
			}
		case <-ticker.C:
			err = export(start, end, minimum, count, project, namespace, service)
			if err != nil {
				return err
			}
		}
	}
}
