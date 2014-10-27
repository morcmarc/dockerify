package utils

import (
	"errors"
	"fmt"
	"os"
)

type PathUtils interface {
	Stat(name string) (os.FileInfo, error)
}

type OSPathUtils struct {
	PathUtils
}

func (o *OSPathUtils) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

type PathValidator struct {
	path      string
	pathUtils PathUtils
}

func NewPathValidator(path string) *PathValidator {
	pu := &OSPathUtils{}
	pv := &PathValidator{
		path:      path,
		pathUtils: pu,
	}
	return pv
}

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
