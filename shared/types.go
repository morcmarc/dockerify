package shared

import (
	"io"
)

// Engine represents a certain project type and encapsulates the logic on how to
// generate the descriptors. A project type is not necessarily a single
// programming language, but rather a collection of languages and
// configurations. The Discover method describes these custom relationships and
// tells whether the given project satisfies them.
type Engine interface {
	Discover() bool
	GenerateDockerfile(out io.Writer) error
	GenerateFigConfig(out io.Writer) error
	Instructions()
}
