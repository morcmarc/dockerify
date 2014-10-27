package nodejs

import (
	"io"

	"github.com/morcmarc/dockerify/shared"
	"github.com/morcmarc/dockerify/utils"
)

type NodeJs struct {
	shared.Engine
}

func (n *NodeJs) Discover(path string) bool {
	return true
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
