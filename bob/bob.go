package bob

import (
	"regexp"
	"strings"
)

// Returns a response based on a remark.
//
// Asking a question -> "Sure."
// YELL AT HIM -> "Whoa, chill out!"
// YELL A QUESTION? -> "Calm down, I don't know what I'm doing!"
// Empty remark -> "Fine. Be that way!"
// Anything else -> "Whatever."
func Hey(remark string) (message string) {
	message = "Whatever."

	remark = strings.TrimSpace(remark)

	if len(remark) == 0 {
		message = "Fine. Be that way!"
	} else {
		isQuestion := strings.HasSuffix(remark, "?")
		isYelling, _ := regexp.MatchString(`^[A-Z0-9\W]*[A-Z]+[A-Z0-9\W]*$`, remark)

		switch {
		case isQuestion && isYelling:
			message = "Calm down, I know what I'm doing!"
		case isQuestion && !isYelling:
			message = "Sure."
		case isYelling:
			message = "Whoa, chill out!"
		}
	}

	return
}
