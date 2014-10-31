/*
This package allows discovering NodeJS-type projects. Will try to see if there
is a "package.json" file in the root path, then it'll do the following checks:

	package.json (mandatory):
		- there is a "server.js" in the root path
		OR
		- "scripts" object has got a "start" command
		- file given to "start" command does exist
		- either "express" or "koa" is present in "dependecies" or "devDependencies" list
*/
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
		checkFiles:    []string{"server.js", "package.json"},
		pathValidator: pValidator,
		fileUtils:     fUtils,
	}
	return njs
}

func (n *NodeJs) Discover() bool {
	valid := n.pathValidator.ValidateFiles(n.checkFiles)
	result := false

	for _, f := range valid {
		if f == "server.js" {
			fmt.Println("-->> Validating server.js")
			result = true
		}
		if f == "package.json" {
			fmt.Println("-->> Validating package.json")
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

func (n *NodeJs) Instructions() {
	usage := `The image assumes that your application:
- listens on port 8080`
	fmt.Println(utils.Colorize(usage, utils.C_RED))
}

func (n *NodeJs) validatePackageJson() bool {
	filename := path.Join(n.path, "package.json")

	b, err := n.fileUtils.ReadFile(filename)
	if err != nil || b == nil {
		fmt.Errorf("Failed reading package file: %s\n", err)
		return false
	}

	var content map[string]interface{}
	if err := json.Unmarshal(b, &content); err != nil {
		fmt.Errorf("Failed unmarshaling package file contents: %s\n", err)
		return false
	}

	scripts := content["scripts"].(map[string]interface{})
	if _, ok := scripts["start"]; !ok {
		fmt.Errorf("No start script attribute")
		return false
	}

	return true
}
