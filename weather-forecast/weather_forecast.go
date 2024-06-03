// Package weather is used for retriving weather information.
package weather

// CurrentCondition for representing present condition.
var CurrentCondition string

// CurrentLocation for representing present location.
var CurrentLocation string

// Forecast will take city and condition as args and ouput the correspondig forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
