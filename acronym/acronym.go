//Package acronym convert a phrase to its acronym.
package acronym

import (
	"regexp"
	"strings"
)

//preProcess removes all noisy chars
func preProcess(s string) []string {
	// find all chars that are not alphabet
	reg := regexp.MustCompile("[^a-zA-Z']+")
	// replace those chars with spaces
	s = reg.ReplaceAllString(s, " ")
	s = strings.ToUpper(s)
	return strings.Fields(s)
}

// Abbreviate Convert a phrase to its acronym.
func Abbreviate(s string) string {
	tokens := preProcess(s)
	fl := make([]rune, len(tokens))
	for i, t := range tokens {
		fl[i] = rune(t[0])
	}
	return string(fl)
}
