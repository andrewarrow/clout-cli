package display

import (
	"fmt"
	"strings"
)

func LeftAligned(thing interface{}, size int) string {
	s := fmt.Sprintf("%v", thing)

	if len(s) > size {
		return s[0:size]
	}
	fill := size - len(s)
	spaces := []string{}
	for {
		spaces = append(spaces, " ")
		if len(spaces) >= fill {
			break
		}
	}
	return s + strings.Join(spaces, "")
}
