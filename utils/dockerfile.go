package utils

import (
	"errors"
	"io"
	"strings"
	"text/template"
)

type DockerfileParams struct {
	Image   string
	Command string
	Expose  string
}

const dockerfileTemplate = `FROM {{.Image}}
{{if .Command}}CMD {{.Command}}{{end}}
{{if .Expose}}EXPOSE {{.Expose}}{{end}}
`

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

func ValidateTemplateParams(params *DockerfileParams) error {
	if params.Image == "" {
		return errors.New("Missing image parameter")
	}
	return nil
}
