package utils

import (
	"testing"
)

func TestColorize(t *testing.T) {
	s := "Hello world"
	e := "\x1b[34mHello world\x1b[0m"
	c := Colorize(s, C_BLUE)
	if c != e {
		t.Errorf("Was expecting: %s, got: %s", e, c)
	}
}
