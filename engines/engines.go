/*
Engines are basically project descriptors and can do certain checks on a
given path to determine the type of the application. If a type has been found
a sample Dockerfile will be generated to specified output (e.g.: stdout or file)
*/
package engines

import (
	"errors"
	"fmt"
	"io"
	"os"
	p "path"

	"github.com/morcmarc/dockerify/engines/golang"
	"github.com/morcmarc/dockerify/engines/nodejs"
	"github.com/morcmarc/dockerify/shared"
	"github.com/morcmarc/dockerify/utils"
)

// Will attempt to determine project type at given path and create a Dockerfile
func GetDockerTemplate(path string, createDockerfile, useFig bool) error {
	engines := createEngines(path)

	for i, engine := range engines {
		if engine.Discover() {
			fmt.Printf("-->> %s\n", utils.Colorize("Found project type: "+i, utils.C_YELLOW))

			dfw := getDockerfileWriter(path, createDockerfile)
			fmt.Printf("-->> %s\n", utils.Colorize("Writing Dockerfile", utils.C_GREEN))
			engine.GenerateDockerfile(dfw)

			if useFig {
				ffw := getFigfileWriter(path)
				fmt.Printf("-->> %s\n", utils.Colorize("Writing Fig config", utils.C_GREEN))
				engine.GenerateFigConfig(ffw)
			}

			fmt.Printf("-->> %s\n", utils.Colorize("Instructions:", utils.C_YELLOW))
			engine.Instructions()

			return nil
		}
	}

	return errors.New("Could not determine project type\n")
}

// Initiate engines
func createEngines(path string) map[string]shared.Engine {
	pathValidator := utils.NewPathValidator(path)
	fileUtils := new(utils.OSFileUtils)
	engines := make(map[string]shared.Engine)

	engines["golang"] = golang.NewEngine(path, pathValidator, fileUtils)
	engines["nodejs"] = nodejs.NewEngine(path, pathValidator, fileUtils)

	return engines
}

func getDockerfileWriter(path string, createDockerfile bool) io.Writer {
	output := os.Stdout
	if createDockerfile {
		f, err := os.Create(p.Join(path, "Dockerfile"))
		if err != nil {
			panic(err)
		}
		output = f
	}

	return output
}

func getFigfileWriter(path string) *os.File {
	figFileWriter, err := os.Create(p.Join(path, "fig.yml"))
	if err != nil {
		panic(err)
	}

	return figFileWriter
}
