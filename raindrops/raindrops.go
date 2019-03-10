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
	factors := getFactors(input)
	output := ""
	for _, factor := range factors {
		switch factor {
		case 3:
			output += "Pling"
		case 5:
			output += "Plang"
		case 7:
			output += "Plong"
		}
	}
	if output == "" {
		output = strconv.Itoa((factors[len(factors)-1]))
	}
	return output
}

/**
getFactors returns all the factors of input with no remainder after division
*/
func getFactors(input int) []int {
	factors := make([]int, input)
	for i := 1; i <= input; i++ {
		if input%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
