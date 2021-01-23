package exporter

import (
	"fmt"
	"time"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Execute(namespace string, service string, development bool) error {
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
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-updater:
			fmt.Println("update")
		case <-ticker.C:
			fmt.Println(count)
		}
	}
}
