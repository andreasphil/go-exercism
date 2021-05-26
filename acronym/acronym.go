package acronym

import (
	"regexp"
	"strings"
)

// For an input, creates an acronym consisting of the first letter in each word
// in the input.
func Abbreviate(s string) (acronym string) {
	acronymCandidates := regexp.MustCompile(`[^a-zA-Z0-9']+`)
	words := acronymCandidates.Split(s, -1)

	for _, word := range words {
		acronym += strings.Split(word, "")[0]
	}

	acronym = strings.ToUpper(acronym)

	return
}
