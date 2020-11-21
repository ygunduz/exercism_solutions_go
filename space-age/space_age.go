package space

type Planet string

const EarthSecondsInOneYear = 31557600.0

// planet earth seconds map
var timeMap = map[Planet]float64{
	"Earth":   EarthSecondsInOneYear,
	"Mercury": 0.2408467 * EarthSecondsInOneYear,
	"Venus":   0.61519726 * EarthSecondsInOneYear,
	"Mars":    1.8808158 * EarthSecondsInOneYear,
	"Jupiter": 11.862615 * EarthSecondsInOneYear,
	"Saturn":  29.447498 * EarthSecondsInOneYear,
	"Uranus":  84.016846 * EarthSecondsInOneYear,
	"Neptune": 164.79132 * EarthSecondsInOneYear,
}

func Age(seconds float64, planet Planet) (result float64) {
	return seconds / timeMap[planet]
}
