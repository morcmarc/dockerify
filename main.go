package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/morcmarc/dockerify/engines"
	"github.com/morcmarc/dockerify/utils"
)

func main() {
	app := cli.NewApp()
	app.Name = "dockerify"
	app.Version = "0.2.0"
	app.Usage = "\"dockerize\" and \"figify\" your app"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "fig, f",
			Usage: "generate fig.yml",
		},
		cli.BoolFlag{
			Name:  "docker, d",
			Usage: "generate Dockerfile",
		},
		cli.StringFlag{
			Name:  "env, e",
			Usage: "environment type: [dev, prod]",
			Value: "prod",
		},
	}
	app.Action = func(c *cli.Context) {
		runApp(c)
	}
	app.Run(os.Args)
}

func runApp(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Printf("Missing path")
		os.Exit(10)
	}

	path := c.Args()[0]
	pathValidator := utils.NewPathValidator(path)

	if err := pathValidator.ValidatePath(); err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(11)
	}

	useDocker := c.Bool("docker")
	useFig := c.Bool("fig")
	env := c.String("env")

	err := engines.GetDockerTemplate(path, useDocker, useFig, env)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(20)
	}
}
