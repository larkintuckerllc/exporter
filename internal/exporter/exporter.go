package exporter

import (
	"fmt"
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
	fmt.Println(count)
	fmt.Println(resourceVersion)
	/*
		done := make(chan bool)
		ticker := time.NewTicker(1 * time.Second)
		go func() {
			for {
				select {
				case <-done:
					ticker.Stop()
					return
				case <-ticker.C:
					fmt.Println("Hello !!")
				}
			}
		}()
		time.Sleep(10 * time.Second)
		done <- true
	*/
	return nil
}
