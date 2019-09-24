package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency count of each run in a given text
func ConcurrentFrequency(texts []string) FreqMap {
	var m = FreqMap{}
	var c = make(chan FreqMap)

	for _, s := range texts {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	for range texts {
		val := <-c
		for k, v := range val {
			m[k] += v
		}
	}

	return m
}
