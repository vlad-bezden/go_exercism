//Package hamming provides solution for
// https://exercism.io/my/solutions/d9fe9aea01ea42ba9647302ef360fb29 exercise.
package hamming

import "errors"

//Distance calculates difference between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must be the same size")
	}
	counter := 0
	for i := range a {
		if a[i] != b[i] {
			counter++
		}
	}

	return counter, nil
}
