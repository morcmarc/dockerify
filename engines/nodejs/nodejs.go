/*
Node project are identified by the following rules:

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

	"github.com/morcmarc/dockerify/docker"
	"github.com/morcmarc/dockerify/fig"
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

type packageJson struct {
	Scripts         map[string]interface{} `json:"scripts"`
	Dependencies    map[string]interface{} `json:"dependencies"`
	DevDependencies map[string]interface{} `json:"devDependencies"`
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
	params := &docker.DockerfileParams{
		Image: "dockerfile/nodejs-runtime",
	}
	if err := docker.ParseTemplate(params, out); err != nil {
		return err
	}
	return nil
}

func (n *NodeJs) GenerateFigConfig(out io.Writer) error {
	ff := fig.NewFigFile()
	fa, err := fig.NewFigApplication("", ".")
	if err != nil {
		return err
	}
	fa.Ports = []string{"8080:8080"}
	ff.AddApplication("app", fa)
	if err := ff.WriteConfig(out); err != nil {
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

	var content packageJson
	if err := json.Unmarshal(b, &content); err != nil {
		fmt.Errorf("Failed unmarshaling package file contents: %s\n", err)
		return false
	}

	if _, ok := content.Scripts["start"]; !ok {
		fmt.Errorf("No start script attribute")
		return false
	}

	return true
}
