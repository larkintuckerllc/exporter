package exporter

import (
	"fmt"
	"time"
)

func Execute(namespace string, service string, development bool) error {
	var count = 0
	clientset, err := k8sAuth(development)
	if err != nil {
		return err
	}
	api := clientset.CoreV1().Endpoints(namespace)
	// TODO: INITIAL COUNT
	fmt.Println(api)   // TODO: REMOVE
	fmt.Println(count) // TODO: REMOVE
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
	return nil
}
