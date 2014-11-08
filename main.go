package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/morcmarc/dockerify/engines"
	"github.com/morcmarc/dockerify/utils"
)

const (
	Version = "0.2.0"
)

var (
	versionFlag      bool
	createDockerfile bool
	useFig           bool
	path             string
)

func init() {
	flag.BoolVar(&versionFlag, "version", false, "Print version and exit")
	flag.BoolVar(&useFig, "fig", true, "Create fig file")
	flag.BoolVar(&createDockerfile, "w", true, "Create and write Dockerfile")
}

func main() {
	flag.Parse()

	if versionFlag {
		fmt.Printf("Dockerify v%s\n", Version)
		os.Exit(0)
	}

	path = flag.Arg(0)
	pathValidator := utils.NewPathValidator(path)

	if err := pathValidator.ValidatePath(); err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(10)
	}

	if err := engines.GetDockerTemplate(path, createDockerfile, useFig); err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(20)
	}

	os.Exit(0)
}
