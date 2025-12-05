// Package weather provides tools for analyzing weather conditions.
package weather

var (
    // CurrentCondition represents the current weather condition in a location.
	CurrentCondition string
    // CurrentLocation represents the location of the checked weather.
	CurrentLocation  string
)

// Forecast returns a string describing a location and its current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
