package main

import (
	"errors"
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
			&cli.IntFlag{
				Name:     "start",
				Usage:    "start hour - between 0 to 23",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "end",
				Usage:    "end hour - between 0 to 23",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "minimum",
				Usage:    "minimum number of pods - greater than 2",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "project",
				Usage:    "GCP project ID",
				Required: true,
			},
			&cli.BoolFlag{
				Name: "development",
			},
		},
		Action: func(c *cli.Context) error {
			namespace := c.String("namespace")
			service := c.String("service")
			start := c.Int("start")
			end := c.Int("end")
			minimum := c.Int("minimum")
			project := c.String("project")
			development := c.Bool("development")
			if start < 0 || start > 23 || end < 0 || end > 23 || start == end {
				err := errors.New("start and end must be between 0 to 23 and unequal")
				return err
			}
			if minimum < 2 {
				err := errors.New("minimum must be greater than 2")
				return err
			}
			err := exporter.Execute(namespace, service, start, end, minimum, project, development)
			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
