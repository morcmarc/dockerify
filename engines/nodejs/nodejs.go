package nodejs

import (
	"encoding/json"
	"fmt"
	"io"
	"path"

	"github.com/morcmarc/dockerify/shared"
	"github.com/morcmarc/dockerify/utils"
)

type NodeJs struct {
	shared.Engine
	path          string
	checkFiles    []string
	pathValidator *utils.PathValidator
	fileUtils     utils.FileUtils
}

func NewEngine(path string, pValidator *utils.PathValidator, fUtils utils.FileUtils) *NodeJs {
	njs := &NodeJs{
		path:          path,
		checkFiles:    []string{"package.json", "server.js"},
		pathValidator: pValidator,
		fileUtils:     fUtils,
	}
	return njs
}

func (n *NodeJs) Discover() bool {
	valid := n.pathValidator.ValidateFiles(n.checkFiles)
	result := false

	for _, f := range valid {
		if f == "package.json" {
			result = n.validatePackageJson()
		}
	}

	return result
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

func (n *NodeJs) validatePackageJson() bool {
	filename := path.Join(n.path, "package.json")

	b, err := n.fileUtils.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Failed reading package file: %s\n", err)
		return false
	}

	var content map[string]interface{}
	if err := json.Unmarshal(b, &content); err != nil {
		fmt.Errorf("Failed unmarshaling package file contents: %s\n", err)
		return false
	}

	deps := content["dependencies"].(map[string]interface{})
	if _, ok := deps["express"]; ok {
		return true
	}

	return false
}
