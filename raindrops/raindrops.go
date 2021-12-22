package raindrops

import (
	"strconv"
)

// Convert a number into a string that contains raindrop sounds corresponding
// to certain potential factors
func Convert(input int) (result string) {
	if (input % 3) == 0 {
		result += "Pling"
	}
	if (input % 5) == 0 {
		result += "Plang"
	}
	if (input % 7) == 0 {
		result += "Plong"
	}

	if len(result) == 0 {
		result = strconv.Itoa(input)
	}

	return
}
