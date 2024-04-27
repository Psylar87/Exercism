// Package weather tells current conditions for a current location.
package weather

// CurrentCondition is a variable that holds a string.
var CurrentCondition string
//CurrentLocation is a variable that holds a strings.
var CurrentLocation string

// Forecast provides the current weather for the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
