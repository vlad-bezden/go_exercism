package space

//Planet type
type Planet string
type year float64

//earthYear number of second in a year on Earth
const earthYear year = 31557600

//planet is a dict of planets and their equivalent earth year
var planets = map[Planet]year{
	"Earth":   earthYear,
	"Mercury": 0.2408467 * earthYear,
	"Venus":   0.61519726 * earthYear,
	"Mars":    1.8808158 * earthYear,
	"Jupiter": 11.862615 * earthYear,
	"Saturn":  29.447498 * earthYear,
	"Uranus":  84.016846 * earthYear,
	"Neptune": 164.79132 * earthYear,
}

//Age calculate how old someone would be on planet
func Age(secs float64, planet Planet) float64 {
	if p, ok := planets[planet]; ok {
		return secs / float64(p)
	}
	return secs
}
