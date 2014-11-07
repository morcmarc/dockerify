package shared

import (
	"io"
)

type Engine interface {
	Discover() bool
	GenerateDockerfile(out io.Writer) error
	GenerateFigConfig(out io.Writer) error
	Instructions()
}
