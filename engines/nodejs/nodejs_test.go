package nodejs

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
	ne := &NodeJs{}
	writer := &WriterMock{}
	expected := "FROM dockerfile/nodejs-runtime\n"

	ne.GenerateDockerfile(writer)

	outstring := fmt.Sprintf("%s", writer.output)
	if outstring != expected {
		t.Errorf("Was expecting %s, got: %s", expected, outstring)
	}
}

func TestDiscoverChecksServerJsFile(t *testing.T) {
	p, err := ioutil.TempDir("/tmp", "nodejs_test")
	if err != nil {
		t.Fatalf("Could not create temp directory: %s", err)
	}

	testData := `"use strict";`
	if err := ioutil.WriteFile(path.Join(p, "server.js"), []byte(testData), 0777); err != nil {
		t.Fatalf("Could not create temp server.js: %s", err)
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

func TestDiscoverChecksPackageFileStartScript(t *testing.T) {
	p, err := ioutil.TempDir("/tmp", "nodejs_test")
	if err != nil {
		t.Fatalf("Could not create temp directory: %s", err)
	}

	testData := `{
		"scripts": {
			"start": "node ./app.js"
		},
		"dependencies": { },
		"devDependencies": { }
	}`
	if err := ioutil.WriteFile(path.Join(p, "package.json"), []byte(testData), 0777); err != nil {
		t.Fatalf("Could not create temp package.json: %s", err)
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
