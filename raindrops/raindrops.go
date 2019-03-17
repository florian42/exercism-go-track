/*
Package raindrops implements a function to convert a number to a string, the
contents of which depend on the number's factors
*/
package raindrops

import (
	"strconv"
)

/*
Convert converts a number to a string, the contents of which depend on the number's factors.
*/
func Convert(input int) string {
	const pling = "Pling"
	const plang = "Plang"
	const plong = "Plong"

	output := ""
	if input%3 == 0 {
		output += pling
	}
	if input%5 == 0 {
		output += plang
	}
	if input%7 == 0 {
		output += plong
	}

	if output == "" {
		return strconv.Itoa(input)
	}

	return output
}
