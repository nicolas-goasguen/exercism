// Package weather is a package for weather.
package weather

var (
	// CurrentCondition is a variable that stores last known condition.
	CurrentCondition string
	// CurrentLocation is a variable that stores last location.
	CurrentLocation string
)

// Forecast is a function that gives forecast for a specified location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
