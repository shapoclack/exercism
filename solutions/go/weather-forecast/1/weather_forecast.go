// Package weather contains func that shows current weather in a current location.
package weather

var (
    // CurrentCondition represents a whether in a string format.
	CurrentCondition string 
    // CurrentLocation represents a city in a string format.
	CurrentLocation  string 
)
// Forecast function returns information about current wheather in current city as string value.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
