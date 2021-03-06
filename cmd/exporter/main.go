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
		Usage: "exports pod metrics",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "project",
				Usage:    "project ID",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "location",
				Usage:    "cluster location",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "cluster",
				Usage:    "cluster name",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "namespace",
				Value: "default",
				Usage: "namespace name",
			},
			&cli.StringFlag{
				Name:     "hpa",
				Usage:    "hpa name",
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
			&cli.BoolFlag{
				Name: "development",
			},
		},
		Action: func(c *cli.Context) error {
			project := c.String("project")
			location := c.String("location")
			cluster := c.String("cluster")
			namespace := c.String("namespace")
			hpa := c.String("hpa")
			start := c.Int("start")
			end := c.Int("end")
			minimum := c.Int("minimum")
			development := c.Bool("development")
			if start < 0 || start > 23 || end < 0 || end > 23 || start == end {
				err := errors.New("start and end must be between 0 to 23 and unequal")
				return err
			}
			if minimum < 2 {
				err := errors.New("minimum must be greater than 2")
				return err
			}
			err := exporter.Execute(project, location, cluster, namespace, hpa, start, end, minimum, development)
			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
