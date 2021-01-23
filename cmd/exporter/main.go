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
			&cli.BoolFlag{
				Name: "development",
			},
		},
		Action: func(c *cli.Context) error {
			namespace := c.String("namespace")
			service := c.String("service")
			development := c.Bool("development")
			err := exporter.Execute(namespace, service, development)
			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
