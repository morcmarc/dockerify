package shared

import (
	"io"
)

type Engine interface {
	NewEngine(path string) *Engine
	Discover() bool
	GenerateDockerfile(out io.Writer) error
}
