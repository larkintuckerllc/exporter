package exporter

import (
	"fmt"
	"time"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Execute(namespace string, app string, start int, end int, minimum int,
	project string, development bool) error {
	clientset, err := k8sAuth(development)
	if err != nil {
		return err
	}
	api := clientset.CoreV1().Pods(namespace)
	pods, resourceVersion, err := initialPods(api, app)
	fmt.Println(pods)
	if err != nil {
		return err
	}
	labelSelector := fmt.Sprintf("app=%s", app)
	listOptions := metaV1.ListOptions{
		LabelSelector:   labelSelector,
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
			fmt.Println(update)
			/*
				count, err = updateCount(update, service, count)
				if err != nil {
					return err
				}
			*/
		case <-ticker.C:
			/*
				if count == 0 {
					break
				}
				// err = export(start, end, minimum, count, project, namespace, service)
				if err != nil {
					return err
				}
			*/
		}
	}
}
