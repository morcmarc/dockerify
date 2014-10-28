package nodejs

import (
	"io"

	"github.com/morcmarc/dockerify/shared"
	"github.com/morcmarc/dockerify/utils"
)

type NodeJs struct {
	shared.Engine
	path          string
	checkFiles    []string
	pathValidator *utils.PathValidator
}

func NewEngine(path string) *NodeJs {
	njs := &NodeJs{
		path:          path,
		checkFiles:    []string{"package.json", "server.js"},
		pathValidator: utils.NewPathValidator(path),
	}
	return njs
}

func (n *NodeJs) Discover() bool {
	return n.pathValidator.ValidateFiles(n.checkFiles)
}

func (n *NodeJs) GenerateDockerfile(out io.Writer) error {
	params := &utils.DockerfileParams{
		Image: "dockerfiles/nodejs-runtime",
	}
	if err := utils.ParseTemplate(params, out); err != nil {
		return err
	}
	return nil
}
