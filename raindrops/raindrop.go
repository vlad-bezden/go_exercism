/*Package raindrops convert a number to a string,
the contents of which depend on the number's factors

If the number has 3 as a factor, output 'Pling'.
If the number has 5 as a factor, output 'Plang'.
If the number has 7 as a factor, output 'Plong'.
If the number does not have 3, 5, or 7 as a factor,
just pass the number's digits straight through.*/
package raindrops

import (
	"sort"
	"strconv"
)

//Convert function converts a number to a string
func Convert(n int) string {
	var drop string
	var keys []int

	dic := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	for k := range dic {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if n%k == 0 {
			drop += dic[k]
		}
	}
	if 0 < len(drop) {
		return drop
	}
	return strconv.Itoa(n)
}
