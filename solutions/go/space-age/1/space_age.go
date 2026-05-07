package spaceage

type Planet string

const (
	mercury = Planet("Mercury")
	venus   = Planet("Venus")
	earth   = Planet("Earth")
	mars    = Planet("Mars")
	jupiter = Planet("Jupiter")
	saturn  = Planet("Saturn")
	uranus  = Planet("Uranus")
	neptune = Planet("Neptune")
)

var orbitalPeriods = map[Planet]float64{
	mercury: 0.2408467,
	venus:   0.61519726,
	earth:   1.0,
	mars:    1.8808158,
	jupiter: 11.862615,
	saturn:  29.447498,
	uranus:  84.016846,
	neptune: 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	orbitalPeriod, found := orbitalPeriods[planet]
	if !found {
		return -1
	}
	return (seconds / orbitalPeriod) / 31557600
}
