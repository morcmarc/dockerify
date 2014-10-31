package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/morcmarc/dockerify/engines"
	"github.com/morcmarc/dockerify/utils"
)

const (
	Version = "1.0.0"
)

var (
	versionFlag      bool
	createDockerfile bool
)

func init() {
	flag.BoolVar(&versionFlag, "version", false, "Print version and exit")
	flag.BoolVar(&createDockerfile, "w", true, "Create and write Dockerfile")
}

func main() {
	flag.Parse()

	if versionFlag {
		fmt.Printf("Dockerify v%s\n", Version)
		os.Exit(0)
	}

	path := flag.Arg(0)
	pathValidator := utils.NewPathValidator(path)

	if err := pathValidator.ValidatePath(); err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(10)
	}

	err := engines.GetDockerTemplate(path, createDockerfile)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(20)
	}

	os.Exit(0)
}
