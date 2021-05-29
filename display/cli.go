package display

import (
	"fmt"
	"strings"
)

func Header(fields ...string) {
	for _, field := range fields {
		fmt.Printf("%s ", field)
	}
	fmt.Printf("\n")
	for _, field := range fields {
		dashes := []string{}
		for i := 0; i < len(field); i++ {
			dashes = append(dashes, "-")
		}
		fmt.Printf("%s ", strings.Join(dashes, ""))
	}
	fmt.Printf("\n")
}
func Row(sizes []int, items ...interface{}) {
	for i, item := range items {
		fmt.Printf("%s ", LeftAligned(item, sizes[i]))
	}
	fmt.Printf("\n")
}

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
