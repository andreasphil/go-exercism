// Package weather can store the current weather conditions for 
// a location and can return a forecast.
package weather

// CurrentCondition stores the weather conditions for the current
// location.
var CurrentCondition string

// CurrentLocation stores the location of the forecast.
var CurrentLocation string

// Forecast returns the forecast for the location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
