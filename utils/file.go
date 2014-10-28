package utils

import (
	"io/ioutil"
)

type FileUtils interface {
	ReadFile(filename string) ([]byte, error)
}

type OSFileUtils struct {
	FileUtils
}

func (o *OSFileUtils) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
