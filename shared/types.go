package shared

import (
	"io"
)

type Engine interface {
	Discover() bool
	Instructions()
	GenerateDockerfile(out io.Writer) error
}
