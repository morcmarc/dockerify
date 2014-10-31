package utils

import (
	"fmt"
	"io"
	"testing"
)

type WriterMock struct {
	io.Writer
	output []byte
}

func (w *WriterMock) Write(p []byte) (n int, err error) {
	w.output = append(w.output, p...)
	return len(p), nil
}

func TestGetCommandStringReturnEmptyStringIfCommandIsEmpty(t *testing.T) {
	commands := ""
	expected := ""
	cs := GetCommandString(commands)
	if cs != expected {
		t.Errorf("Was expecting: %s\n, got: %s", expected, cs)
	}
}

func TestGetCommandString(t *testing.T) {
	commands := "command --args1 value1 -a v2"
	expected := "[\"command\",\"--args1\",\"value1\",\"-a\",\"v2\"]"
	cs := GetCommandString(commands)
	if cs != expected {
		t.Errorf("Was expecting: %s\n, got: %s", expected, cs)
	}
}

func TestParseTemplateShouldFailWithoutImage(t *testing.T) {
	params := &DockerfileParams{}
	writer := &WriterMock{}

	if err := ParseTemplate(params, writer); err == nil {
		t.Errorf("Was expecting to fail")
	}
}

func TestParseTemplateWorksWithImageOnly(t *testing.T) {
	params := &DockerfileParams{
		Image: "test/image",
	}
	writer := &WriterMock{}
	expected := "FROM test/image\n"

	ParseTemplate(params, writer)

	outstring := fmt.Sprintf("%s", writer.output)
	if outstring != expected {
		t.Errorf("Was expecting %s, got: %s", expected, outstring)
	}
}

func TestParseTemplateRendersAllParams(t *testing.T) {
	params := &DockerfileParams{
		Image:   "test/image",
		Command: "run test.js -p 8080",
		Expose:  "8080",
	}
	writer := &WriterMock{}
	expected := "FROM test/image\nCMD [\"run\",\"test.js\",\"-p\",\"8080\"]\nEXPOSE 8080\n"

	ParseTemplate(params, writer)

	outstring := fmt.Sprintf("%s", writer.output)
	if outstring != expected {
		t.Errorf("Was expecting %s, got: %s", expected, outstring)
	}
}
