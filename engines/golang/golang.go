package golang

import (
	"fmt"
	"io"

	"github.com/morcmarc/dockerify/shared"
	"github.com/morcmarc/dockerify/utils"
)

type Golang struct {
	shared.Engine
	path          string
	checkFiles    []string
	pathValidator *utils.PathValidator
	fileUtils     utils.FileUtils
}

func NewEngine(path string, pValidator *utils.PathValidator, fUtils utils.FileUtils) *Golang {
	golang := &Golang{
		path:          path,
		checkFiles:    []string{"main.go"},
		pathValidator: pValidator,
		fileUtils:     fUtils,
	}
	return golang
}

func (g *Golang) Discover() bool {
	valid := g.pathValidator.ValidateFiles(g.checkFiles)
	if len(valid) > 0 {
		return true
	}
	return false
}

func (g *Golang) GenerateDockerfile(out io.Writer) error {
	params := &utils.DockerfileParams{
		Image: "google/golang-runtime",
	}
	if err := utils.ParseTemplate(params, out); err != nil {
		return err
	}
	return nil
}

func (g *Golang) GenerateFigConfig(out io.Writer) error {
	ff := utils.NewFigFile()
	fa, err := utils.NewFigApplication("", ".")
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

func (g *Golang) Instructions() {
	usage := `The image assumes that your application:
- listens on port 8080
- may have a .godir file containing the import path for your application if it vendors its dependencies`
	fmt.Println(utils.Colorize(usage, utils.C_RED))
}
