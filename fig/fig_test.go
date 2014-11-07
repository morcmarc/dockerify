package fig

import (
	"testing"
)

func TestThatNewFigApplicationValidatesFields(t *testing.T) {
	_, err := NewFigApplication("", "")
	if err == nil {
		t.Errorf("Was expecting error")
	}
}

func TestGetYamlReturnsAllApplications(t *testing.T) {
	ff := NewFigFile()

	fa1, err := NewFigApplication("", ".")
	if err != nil {
		t.Errorf("Wasn't expecting error: %s", err)
	}

	fa1.MemLimit = 1024
	fa1.VolumesFrom = []string{"a:a", "b"}
	fa1.WorkingDir = "/data"
	fa1.Links = []string{"db"}

	fa2, err := NewFigApplication("mongo", "")
	if err != nil {
		t.Errorf("Wasn't expecting error: %s", err)
	}

	expected := `app:
  build: .
  links:
  - db
  volumes_from:
  - a:a
  - b
  working_dir: /data
  mem_limit: 1024
db:
  image: mongo
`

	ff.AddApplication("app", fa1)
	ff.AddApplication("db", fa2)

	y, err := ff.GetYaml()
	if err != nil {
		t.Errorf("Wasn't expecting error: %s", err)
	}
	s := string(y)
	if s != expected {
		t.Errorf("Was expecting: %s, got: %s", expected, s)
	}
}
