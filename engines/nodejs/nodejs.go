package nodejs

import (
	"fmt"

	"github.com/morcmarc/dockerify/shared"
)

var (
	Image   = "dockerfile/nodejs-runtime"
	Custom  = "# Any custom commands"
	Expose  = "8080"
	Command = "npm start"
)

type NodeJs struct {
	shared.Engine
}

func (n *NodeJs) Discover(path string) bool {
	return true
}

func (n *NodeJs) GetDockerfileTemplate() string {
	commands := shared.GetCommandString(Command)
	return fmt.Sprintf(shared.DockerfileTemplate, Image, Custom, commands, Expose)
}
