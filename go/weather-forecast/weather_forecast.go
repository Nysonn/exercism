// Package weather provides information on a city and its current weather condition.
package weather

// CurrentCondition represents a the current condition of a given place.
var CurrentCondition string

// CurrentLocation represents the current place.
var CurrentLocation string

// Forecast is a function reponsible for returning the current weather condition of a given place.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
