package exporter

import (
	"fmt"
	"time"
)

func Execute(namespace string, service string) error {
	fmt.Println(namespace)
	fmt.Println(service)
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
