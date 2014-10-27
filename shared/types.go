package shared

import (
	"io"
)

type Engine interface {
	Discover(path string) bool
	GenerateDockerfile(out io.Writer) error
}
