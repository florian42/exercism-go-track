package space

/*
Age calculates the age of a person on different planets in years.
*/
func Age(seconds float64, planet string) float64 {
	const earth float64 = 31557600
	switch planet {
	case "Earth":
		return seconds / earth
	case "Mercury":
		return seconds / (earth * 0.2408467)
	case "Venus":
		return seconds / (earth * 0.61519726)
	case "Mars":
		return seconds / (earth * 1.8808158)
	case "Jupiter":
		return seconds / (earth * 11.862615)
	case "Saturn":
		return seconds / (earth * 29.447498)
	case "Uranus":
		return seconds / (earth * 84.016846)
	case "Neptune":
		return seconds / (earth * 164.79132)
	default:
		return -1
	}
}
