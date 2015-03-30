package golang

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/morcmarc/dockerify/utils"
)

type WriterMock struct {
	io.Writer
	output []byte
}

func (w *WriterMock) Write(p []byte) (n int, err error) {
	w.output = append(w.output, p...)
	return len(p), nil
}

type FileUtilsMock struct {
	utils.FileUtils
}

type FileInfoMock struct {
	os.FileInfo
}

func (f FileUtilsMock) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func TestGenerateDockerfile(t *testing.T) {
	golang := &Golang{}
	writer := &WriterMock{}
	expected := "FROM dockerfile/go-runtime\n"

	golang.GenerateDockerfile(writer)

	outstring := fmt.Sprintf("%s", writer.output)
	if outstring != expected {
		t.Errorf("Was expecting %s, got: %s", expected, outstring)
	}
}

func TestDiscoverChecksMainGoFile(t *testing.T) {
	p, err := ioutil.TempDir("/tmp", "golang_test")
	if err != nil {
		t.Fatalf("Could not create temp directory: %s", err)
	}

	testData := "package main"
	if err := ioutil.WriteFile(path.Join(p, "main.go"), []byte(testData), 0777); err != nil {
		t.Fatalf("Could not create temp main.go: %s", err)
	}

	fu := new(FileUtilsMock)
	pv := utils.NewPathValidator(p)
	ne := NewEngine(p, pv, fu)

	if !ne.Discover() {
		t.Errorf("Was expecting true, got false")
	}

	if err := os.RemoveAll(p); err != nil {
		t.Fatalf("Could not remove temp files: %s", err)
	}
}
