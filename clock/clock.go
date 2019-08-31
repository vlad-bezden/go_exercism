//Package clock provides solution for "Clock" exercise on Exercism
package clock

import "fmt"

const (
	hPerDay  = 24
	mPerHour = 60
	mPerDay  = hPerDay * mPerHour
)

//Clock structure for collecting hours and minutes
type Clock struct {
	h, m int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02v:%02v", c.h, c.m)
}

//Add adds minutes to the clock time
func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

//Subtract subtracts minutes from the clock time
func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

//New converts a phrase to its acronym.
func New(h, m int) Clock {
	// convert everything to minutes
	totMins := (h%hPerDay)*mPerHour + m
	fh := (totMins%mPerDay + mPerDay) / mPerHour % hPerDay
	fm := (totMins%mPerHour + mPerHour) % mPerHour
	return Clock{fh, fm}
}
