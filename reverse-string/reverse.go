//Package reverse provides solution for 'Reverse' exercise on Exercism
package reverse

import (
	"strings"
)

//Reverse reverses string using strings.Builder. It's about 3 times faster
//than the one with using a string
// BenchmarkReverse-8	3000000	499 ns/op	56 B/op	6 allocs/op
func Reverse(in string) string {
	var sb strings.Builder
	runes := []rune(in)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

//Reverse reverses string using string
// BenchmarkReverse-8	1000000	1571 ns/op	176 B/op	29 allocs/op
// func Reverse(in string) (out string) {
// 	for _, r := range in {
// 		out = string(r) + out
// 	}
// 	return
// }
