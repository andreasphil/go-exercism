package space

import (
	"log"
	"strings"
)

type Planet string

var orbPeriods = map[string]float64{
	"mercury": 0.2408467,
	"venus":   0.61519726,
	"earth":   1.0,
	"mars":    1.8808158,
	"jupiter": 11.862615,
	"saturn":  29.447498,
	"uranus":  84.016846,
	"neptune": 164.79132,
}

var earthYearSeconds = float64(31557600)

// For an age in seconds, calculate how many years that would be on the given
// planet.
func Age(seconds float64, planet Planet) (age float64) {
	planetName := strings.ToLower(string(planet))

	orbPeriod, ok := orbPeriods[planetName]
	if !ok {
		log.Fatalf("%v is not a planet I've heard of", planet)
	}

	earthYearOnPlanet := 1 / orbPeriod
	ageOnEarth := seconds / earthYearSeconds
	age = ageOnEarth * earthYearOnPlanet

	return
}
