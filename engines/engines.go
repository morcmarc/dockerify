package engines

import (
	"errors"
	"fmt"
	"os"

	"github.com/morcmarc/dockerify/engines/nodejs"
	"github.com/morcmarc/dockerify/shared"
)

func GetEngines() map[string]shared.Engine {
	engines := make(map[string]shared.Engine)
	engines["nodejs"] = &nodejs.NodeJs{}
	return engines
}

func GetDockerTemplate(path string) error {
	engines := GetEngines()
	for i, engine := range engines {
		if engine.Discover(path) {
			fmt.Printf("Found project type: %s\n", i)
			engine.GenerateDockerfile(os.Stdout)
			return nil
		}
	}
	return errors.New("Could not determine project type")
}
