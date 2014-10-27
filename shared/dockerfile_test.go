package shared

import (
	"testing"
)

func TestGetCommandString(t *testing.T) {
	commands := "command --args1 value1 -a v2"
	expected := "[\"command\",\"--args1\",\"value1\",\"-a\",\"v2\"]"
	cs := GetCommandString(commands)
	if cs != expected {
		t.Errorf("Was expecting: %s\n, got: %s", expected, cs)
	}
}
