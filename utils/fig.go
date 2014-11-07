package utils

import (
	"errors"

	"gopkg.in/yaml.v2"
)

// Reference: http://www.fig.sh/yml.html
type FigApplication struct {
	Image       string   `image,omitempty`
	Build       string   `build,omitempty`
	Links       []string `links,omitempty`
	Ports       []string `ports,omitempty`
	Expose      []string `expose,omitempty`
	Volumes     []string `volumes,omitempty`
	VolumesFrom []string `volumes_from,omitempty`
	Environment []string `environment,omitempty`
	Net         string   `net,omitempty`
	Dns         []string `dns,omitempty`
	WorkingDir  string   `working_dir,omitempty`
	Entrypoint  string   `entrypoint,omitempty`
	User        string   `user,omitempty`
	Hostname    string   `hostname,omitempty`
	Domainname  string   `domainname,omitempty`
	MemLimit    int      `mem_limit,omitempty`
	Privileged  bool     `privileged,omitempty`
}

type FigFile struct {
	Applications map[string]*FigApplication
}

func NewFigFile() *FigFile {
	f := &FigFile{
		Applications: make(map[string]*FigApplication),
	}
	return f
}

func NewFigApplication(image, build string) (*FigApplication, error) {
	if image == "" && build == "" {
		return nil, errors.New("You must provide either an image or a build attribute")
	}
	fa := &FigApplication{
		Image: image,
		Build: build,
	}
	return fa, nil
}

func (f *FigFile) AddApplication(name string, a *FigApplication) {
	f.Applications[name] = a
}

func (f *FigFile) GetYaml() ([]byte, error) {
	y, err := yaml.Marshal(&f.Applications)
	if err != nil {
		return nil, err
	}
	return y, nil
}
