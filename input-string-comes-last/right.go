package to

import (
	"fmt"
	"strings"
)

// Indented returns a string with each line indented by the specified
// number of spaces. Carriage returns are stripped (if found) as
// a side-effect. Per [pkg/text/template] rules, the string input
// argument is last so that this function can be used as-is within
// template pipe constructs:
//
//	{{ .Some | indent 4 }}
func Indented(indent int, in string) string {
	var buf string
	for _, line := range Lines(in) {
		buf += fmt.Sprintln(strings.Repeat(" ", indent) + line)
	}
	return buf
}
