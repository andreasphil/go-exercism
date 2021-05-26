package twofer

import (
	"fmt"
	"strings"
)

func ShareWith(name string) string {
	if len(strings.TrimSpace(name)) == 0 {
		name = "you"
	}

	return fmt.Sprintf("One for %v, one for me.", name)
}
