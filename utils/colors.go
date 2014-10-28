package utils

import (
	"fmt"
)

const (
	C_GREEN  = 32
	C_RED    = 31
	C_YELLOW = 33
	C_BLUE   = 34
)

// Wrap the given string in the selected terminal color code
func Colorize(s string, c int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}
