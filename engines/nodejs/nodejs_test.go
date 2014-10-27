package nodejs

import (
	"testing"

	// "github.com/morcmarc/dockerify/shared"
)

var template = `
# Pull base image
FROM dockerfile/nodejs-runtime

# Any custom commands

# Define default command
CMD ["npm","start"]

# Expose port
EXPOSE 8080
`

func TestGetDockerfileTemplateNotEmpty(t *testing.T) {
	e := &NodeJs{}
	if e.GetDockerfileTemplate() == "" {
		t.Errorf("Was expecting dockerfile, got empty string")
	}
}

func TestGetDockerfileTemplateCorrectTemplate(t *testing.T) {
	e := &NodeJs{}
	dft := e.GetDockerfileTemplate()
	if dft != template {
		t.Errorf("Was expecting: %s\n, got: %s", template, dft)
	}
}
