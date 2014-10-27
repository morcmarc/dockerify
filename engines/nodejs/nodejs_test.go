package nodejs

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

func TestGenerateDockerfile(t *testing.T) {
	ne := &NodeJs{}
	writer := &WriterMock{}
	expected := "FROM dockerfiles/nodejs-runtime\n\n\n"

	ne.GenerateDockerfile(writer)

	outstring := fmt.Sprintf("%s", writer.output)
	if outstring != expected {
		t.Errorf("Was expecting %s, got: %s", expected, outstring)
	}
}
