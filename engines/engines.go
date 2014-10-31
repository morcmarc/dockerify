/*
Engines are basically project descriptors and can do certain checks on a
given path to determine the type of the application. If a type has been found
a sample Dockerfile will be generated to specified output (e.g.: stdout or file)
*/
package engines

import (
	"errors"
	"fmt"
	"os"

	"github.com/morcmarc/dockerify/engines/golang"
	"github.com/morcmarc/dockerify/engines/nodejs"
	"github.com/morcmarc/dockerify/shared"
	"github.com/morcmarc/dockerify/utils"
)

// Will attempt to determine project type at given path and create a Dockerfile
func GetDockerTemplate(path string) error {
	engines := createEngines(path)

	for i, engine := range engines {
		if engine.Discover() {
			fmt.Printf("-->> Found project type: %s\n", utils.Colorize(i, utils.C_YELLOW))
			fmt.Printf("-->> %s:\n\n", utils.Colorize("Dockerfile", utils.C_GREEN))
			engine.GenerateDockerfile(os.Stdout)
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
