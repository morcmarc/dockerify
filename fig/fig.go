/*
Fig package allows easy handling of fig configs.
*/
package fig

import (
	"errors"
	"io"

	"gopkg.in/yaml.v2"
)

// FigApplication represents a single "unit" in the fig config. Supports all
// the standard fig / docker attributes.
//
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

// FigFile is just a wrapper for a collection of FigApplications
type FigFile struct {
	Applications map[string]*FigApplication
}

// Create a new FigFile
func NewFigFile() *FigFile {
	f := &FigFile{
		Applications: make(map[string]*FigApplication),
	}
	return f
}

// Create a new FigApplication. Must have either an image or build attribute,
// otherwise an error will be returned
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

// Append given FigApplication to the config.
func (f *FigFile) AddApplication(name string, a *FigApplication) {
	f.Applications[name] = a
}

// Returns the marshalled config file.
func (f *FigFile) GetYaml() ([]byte, error) {
	y, err := yaml.Marshal(&f.Applications)
	if err != nil {
		return nil, err
	}
	return y, nil
}

// Write config onto the given io.Writer (i.e.: file, stdout etc)
func (f *FigFile) WriteConfig(out io.Writer) error {
	b, err := f.GetYaml()
	if err != nil {
		return err
	}
	n, err := out.Write(b)
	if err != nil || n == 0 {
		return err
	}
	return nil
}
