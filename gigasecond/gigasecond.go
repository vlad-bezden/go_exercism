// Package gigasecond calculates the time at which someone has lived for a gigasecond.
package gigasecond

// import path for the time package from the standard library
import "time"

const gigasecond = 1e9 * time.Second

// AddGigasecond adds 10^9 seconds to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
