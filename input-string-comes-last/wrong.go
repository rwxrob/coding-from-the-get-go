package to

import (
	"fmt"
	"strings"
)

// WRONG! DO NOT DO THIS. This unnecessarily prevents a perfectly good
// function from being used within text/template pipelines.
func Indented(in string, indent int) string {
	var buf string
	for _, line := range Lines(in) {
		buf += fmt.Sprintln(strings.Repeat(" ", indent) + line)
	}
	return buf
}
