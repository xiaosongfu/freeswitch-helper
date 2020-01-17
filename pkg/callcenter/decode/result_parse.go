package decode

import (
	"strings"
)

const empty = "+OK"

// str = ""
// str = "+OK"
func parse(str string) string {
	if str == "" {
		return ""
	}
	if str == empty {
		return ""
	}

	lines := strings.Split(str, "\n")

	return lines[0]
}
