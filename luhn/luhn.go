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

func reverse(in string) string {
	var sb strings.Builder
	runes := []rune(in)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

// charToInt convert rune/char to integer/digit
func charToInt(c rune) int {
	return int(c - '0')
}

//Valid determine for a given number whether or not it is valid per Luhn formula
func Valid(s string) bool {
	// remove all empty spaces
	s = removeSpaces(s)
	var counter int

	if !validations(s) {
		return false
	}

	for i, r := range reverse(s) {
		iv := charToInt(r)
		if (i+1)%2 == 0 {
			if v := iv * 2; v < 10 {
				counter += v
			} else {
				for _, d := range strconv.Itoa(v) {
					counter += charToInt(d)
				}
			}
		} else {
			counter += iv
		}
	}
	return counter%10 == 0
}
