package exporter

import (
	"fmt"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Execute(project string, location string, cluster string,
	namespace string, hpa string,
	start int, end int, minimum int,
	development bool) error {
	clientset, err := k8sAuth(development)
	if err != nil {
		return err
	}
	api := clientset.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace)
	replicas, resourceVersion, err := initialReplicas(api, hpa)
	if err != nil {
		return err
	}
	fmt.Println(replicas)
	fmt.Println(resourceVersion)
	listOptions := metaV1.ListOptions{
		ResourceVersion: resourceVersion,
	}
	watcher, err := api.Watch(listOptions)
	if err != nil {
		return err
	}
	defer watcher.Stop()
	updater := watcher.ResultChan()
	for {
		select {
		case update := <-updater:
			replicas, err = updateReplicas(update, hpa, replicas)
			if err != nil {
				return err
			}
			fmt.Println(replicas)
		}
	}
	/*
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
			case <-ticker.C:
				count := len(pods)
				if count == 0 {
					break
				}
				err = export(project, location, cluster, namespace, start, end, minimum, pods)
				if err != nil {
					return err
				}
			}
		}
	*/
}
