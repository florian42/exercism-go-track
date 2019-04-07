/*
Package scrabble implements a function to calculate the scrabble score for a given word.
*/
package scrabble

import "unicode"

/*
Score given a word, computes the scrabble score for that word.
*/
func Score(input string) int {
	sum := 0
	for _, char := range input {
		switch unicode.ToUpper(char) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			sum++
		case 'D', 'G':
			sum += 2
		case 'B', 'C', 'M', 'P':
			sum += 3
		case 'F', 'H', 'V', 'W', 'Y':
			sum += 4
		case 'K':
			sum += 5
		case 'J', 'X':
			sum += 8
		case 'Q', 'Z':
			sum += 10
		}
	}
	return sum
}
