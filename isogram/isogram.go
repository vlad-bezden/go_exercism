//Package isogram provides solution for 'Isogram' exercise on Exercism
package isogram

import (
	"regexp"
	"strings"
)

//IsIsogram determines if a word or phrase is an isogram
func IsIsogram(w string) bool {
	w = strings.ToLower(w)
	// find all non alphanumeric (same as [^a-zA-Z0-9_])
	reg := regexp.MustCompile(`\W`)
	// replace found in template with empty string
	w = reg.ReplaceAllString(w, "")
	var set = make(map[rune]bool)
	for _, c := range w {
		if _, ok := set[c]; ok {
			return false
		}
		set[c] = true
	}
	return true
}
