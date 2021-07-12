package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Distance is not defined for different lengths")
	}

	distance := 0
	for i := range a {
		if (a[i] != b[i]) {
			distance++
		}
	}

	return distance, nil
}
