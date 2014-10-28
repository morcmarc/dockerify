package utils

import (
	"errors"
	"os"
	"testing"
)

type PathUtilsMock struct {
	PathUtils
}

type DirInfoMock struct {
	os.FileInfo
}

type FileInfoMock struct {
	os.FileInfo
}

func (f FileInfoMock) IsDir() bool {
	return false
}

func (f DirInfoMock) IsDir() bool {
	return true
}

func (p PathUtilsMock) Stat(path string) (os.FileInfo, error) {
	if path == "filepath" {
		fi := new(FileInfoMock)
		return fi, nil
	}
	if path == "invalidpath" {
		return nil, errors.New("invalid path")
	}
	if path == "dirpath" {
		di := new(DirInfoMock)
		return di, nil
	}
	if path == "dirpath/a.ext" {
		fi := new(FileInfoMock)
		return fi, nil
	}
	return nil, errors.New("Does not exists")
}

func TestValidatePathWithValidPath(t *testing.T) {
	validator := &PathValidator{
		path:      "dirpath",
		pathUtils: new(PathUtilsMock),
	}
	if err := validator.ValidatePath(); err != nil {
		t.Errorf("Was not expecting error")
	}
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

func TestValidateFilesWithSingleValidFile(t *testing.T) {
	validator := &PathValidator{
		path:      "dirpath",
		pathUtils: new(PathUtilsMock),
	}
	files := []string{"a.ext"}
	l := len(validator.ValidateFiles(files))
	if l != 1 {
		t.Errorf("Was expecting 1, got %d", l)
	}
}

func TestValidateFilesWithInvalidFile(t *testing.T) {
	validator := &PathValidator{
		path:      "dirpath",
		pathUtils: new(PathUtilsMock),
	}
	files := []string{"i.vld"}
	l := len(validator.ValidateFiles(files))
	if l != 0 {
		t.Errorf("Was expecting 0, got: %d", l)
	}
}

func TestValidateFilesWithMultipleFiles(t *testing.T) {
	validator := &PathValidator{
		path:      "dirpath",
		pathUtils: new(PathUtilsMock),
	}
	files := []string{"i.vld", "a.ext"}
	l := len(validator.ValidateFiles(files))
	if l != 1 {
		t.Errorf("Was expecting true, got false")
	}
}
