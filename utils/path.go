package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// PathUtils is wrapper interface for path operations
type PathUtils interface {
	Stat(name string) (os.FileInfo, error)
}

// OSPathUtils wraps the OS package and implements the PathUtils interface
type OSPathUtils struct {
	PathUtils
}

// Same as os.Stat(string) (os.FileInfo, error)
func (o *OSPathUtils) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

type PathValidator struct {
	path      string
	pathUtils PathUtils
}

// Create new PathValidator
func NewPathValidator(path string) *PathValidator {
	pu := &OSPathUtils{}
	pv := &PathValidator{
		path:      path,
		pathUtils: pu,
	}
	return pv
}

// Check whether the path exists and is a directory. Required because dockerify
// cannot be run on files.
func (p *PathValidator) ValidatePath() error {
	if p.path == "" {
		return errors.New("Invalid path")
	}
	fi, err := p.pathUtils.Stat(p.path)
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid path: %s", err))
	}
	if !fi.IsDir() {
		return errors.New("Path is not a directory")
	}
	return nil
}

// Check that files can be found under the basepath.
func (p *PathValidator) ValidateFiles(files []string) []string {
	found := []string{}
	for _, f := range files {
		fp := filepath.Join(p.path, f)
		if _, err := p.pathUtils.Stat(fp); err == nil {
			found = append(found, f)
		}
	}
	return found
}
