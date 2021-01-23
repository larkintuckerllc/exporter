package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "exporter",
		Usage: "TODO: usage",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Value: "default",
				Usage: "namespace name",
			},
			&cli.StringFlag{
				Name:     "service",
				Usage:    "service name",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			namespace := c.String("namespace")
			service := c.String("service")
			fmt.Println(namespace)
			fmt.Println(service)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
