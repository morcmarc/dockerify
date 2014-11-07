package utils

import (
	"io/ioutil"
)

// Interface for file operations
type FileUtils interface {
	ReadFile(filename string) ([]byte, error)
}

// Wrapper for built-in io/ioutil package
type OSFileUtils struct {
	FileUtils
}

// Same as ioutil.ReadFile(filename string)
func (o *OSFileUtils) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
