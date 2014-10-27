package utils

import (
	"errors"
	"os"
	"testing"
)

type PathUtilsMock struct {
	PathUtils
}

type FileInfoMock struct {
	os.FileInfo
}

func (f FileInfoMock) IsDir() bool {
	return false
}

func (p PathUtilsMock) Stat(path string) (os.FileInfo, error) {
	if path == "filepath" {
		fi := new(FileInfoMock)
		return fi, nil
	}
	if path == "invalidpath" {
		return nil, errors.New("invalid path")
	}
	return nil, nil
}

func TestValidatePathWithEmptyPath(t *testing.T) {
	validator := &PathValidator{
		path:      "",
		pathUtils: new(PathUtilsMock),
	}
	if err := validator.ValidatePath(); err == nil {
		t.Errorf("Was expecting error")
	}
}

func TestValidatePathWithFilePath(t *testing.T) {
	validator := &PathValidator{
		path:      "filepath",
		pathUtils: new(PathUtilsMock),
	}
	if err := validator.ValidatePath(); err == nil {
		t.Errorf("Was expecting directory path error")
	}
}

func TestValidatePathWithInvalidPath(t *testing.T) {
	validator := &PathValidator{
		path:      "invalidpath",
		pathUtils: new(PathUtilsMock),
	}
	if err := validator.ValidatePath(); err == nil {
		t.Errorf("Was expecting invalid path error")
	}
}
