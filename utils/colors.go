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

// Wrap string with terminal color codes
func Colorize(s string, c int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}
