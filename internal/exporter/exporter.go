package exporter

import "fmt"

func Execute(namespace string, service string) error {
	fmt.Println(namespace)
	fmt.Println(service)
	return nil
}
