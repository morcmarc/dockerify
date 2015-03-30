/*
Provides utility functions for generating dockerfiles.
*/
package docker

import (
	"errors"
	"io"
	"strings"
	"text/template"
)

// Dockerfile descriptor
type DockerfileParams struct {
	Image   string
	Command string
	Expose  string
}

// Raw Dockerfile template
const dockerfileTemplate = "FROM {{.Image}}\n{{if .Command}}CMD {{.Command}}\n{{end}}{{if .Expose}}EXPOSE {{.Expose}}\n{{end}}"

/*
Wrap each word in the given string in doublequotes, then the whole sentence
into square brackets. Example:

	"docker -t -i test/image ./run.sh" =>
	["docker","-t","-i","test/image","./run.sh"]
*/
func GetCommandString(commands string) string {
	if len(commands) == 0 {
		return ""
	}
	// Explode string into an array
	commandArr := strings.Split(commands, " ")
	// Init empty array for command partials
	commandStr := []string{}
	// Surround each command part with double quotes
	for _, c := range commandArr {
		commandStr = append(commandStr, "\""+c+"\"")
	}
	// Surround the whole thing with [ and ]
	return "[" + strings.Join(commandStr, ",") + "]"
}

// Compile template with the given DockerfileParams and write the result onto
// the output.
func ParseTemplate(params *DockerfileParams, out io.Writer) error {
	if err := ValidateTemplateParams(params); err != nil {
		return err
	}

	params.Command = GetCommandString(params.Command)

	t := template.Must(template.New("Dockerfile").Parse(dockerfileTemplate))
	if t == nil {
		return errors.New("Could not parse template")
	}

	if err := t.Execute(out, params); err != nil {
		return err
	}

	return nil
}

// Validate DockerfileParams object
func ValidateTemplateParams(params *DockerfileParams) error {
	if params.Image == "" {
		return errors.New("Missing image parameter")
	}
	return nil
}
