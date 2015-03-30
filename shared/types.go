package shared

import (
	"io"
)

// Engine represents a project type and encapsulates logic about how to generate
// descriptors.
type Engine interface {
	// Returs true if the current path matches the project criteria
	Discover() bool
	// Write dockerfile
	GenerateDockerfile(out io.Writer) error
	// Write fig config
	GenerateFigConfig(out io.Writer) error
	// Print usage instructions
	Instructions()
}
