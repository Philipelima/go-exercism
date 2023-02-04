// Package weather provides tools to 
// takes weater forecast.
package weather


// CurrentCondition represents the current waether condition.
var CurrentCondition string

// CurrentLocation represents the location that takes its weather forecast. 
var CurrentLocation string

// Forecast function gets waether forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
