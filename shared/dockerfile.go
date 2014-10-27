package shared

import (
	"strings"
)

const DockerfileTemplate = `
# Pull base image
FROM %s

%s

# Define default command
CMD %s

# Expose port
EXPOSE %s
`

func GetCommandString(commands string) string {
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
