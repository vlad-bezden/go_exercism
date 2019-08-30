//Package luhn provides solution for "Luhn" exercise on Exercism
package luhn

import (
	"strconv"
	"strings"
)

func removeSpaces(s string) string {
	return strings.Join(strings.Fields(s), "")
}

func validations(s string) bool {
	// check if string is only one char length
	if len(s) <= 1 {
		return false
	}
	// string must only contain numeric chars
	if _, err := strconv.ParseInt(s, 10, 64); err != nil {
		return false
	}
	return true
}

//Valid determine for a given number whether or not it is valid per Luhn formula
func Valid(s string) bool {
	// remove all empty spaces
	s = removeSpaces(s)
	var counter int

	if !validations(s) {
		return false
	}

	for i := range s {
		iv := int(s[len(s)-1-i] - '0')
		if i%2 == 1 {
			if v := iv * 2; v < 10 {
				counter += v
			} else {
				counter += (v - 9)
			}
		} else {
			counter += iv
		}
	}
	return counter%10 == 0
}
