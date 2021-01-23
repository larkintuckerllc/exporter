package main

import (
	"log"
	"os"

	"github.com/larkintuckerllc/exporter/internal/exporter"
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
			err := exporter.Execute(namespace, service)
			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
