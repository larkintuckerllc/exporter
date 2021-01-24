package exporter

import (
	"fmt"
	"time"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Execute(project string, location string, cluster string,
	namespace string, app string,
	start int, end int, minimum int,
	development bool) error {
	clientset, err := k8sAuth(development)
	if err != nil {
		return err
	}
	api := clientset.CoreV1().Pods(namespace)
	pods, resourceVersion, err := initialPods(api, app)
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
			pods, err = updatePods(update, pods)
			if err != nil {
				return err
			}
			fmt.Println(pods)
		case <-ticker.C:
			err = export(project, location, cluster, namespace, start, end, minimum, pods)
			if err != nil {
				return err
			}
		}
	}
}
