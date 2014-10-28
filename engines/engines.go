package engines

import (
	"errors"
	"fmt"
	"os"

	"github.com/morcmarc/dockerify/engines/nodejs"
	"github.com/morcmarc/dockerify/shared"
)

// Will attempt to determine project type and parses a Dockerfile template
func GetDockerTemplate(path string) error {
	engines := createEngines(path)

	for i, engine := range engines {
		if engine.Discover() {
			fmt.Printf("Found project type: %s\n", i)
			engine.GenerateDockerfile(os.Stdout)
			return nil
		}
	}

	return errors.New("Could not determine project type\n")
}

// Initiate engines
func createEngines(path string) map[string]shared.Engine {
	engines := make(map[string]shared.Engine)

	engines["nodejs"] = nodejs.NewEngine(path)

	return engines
}
