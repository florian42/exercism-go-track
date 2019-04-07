/*
Package isogram implements a method to determine if a given word is an isogram.
*/
package isogram

import (
	"strings"
	"unicode"
)

/*
IsIsogram determines if a word or phrase is an isogram.
*/
func IsIsogram(input string) bool {
	if input == "" {
		return true
	}
	input = strings.ToLower(input)
	for i, char := range input {
		lastIndex := strings.LastIndex(input, string(char))
		if lastIndex > i == true && unicode.IsLetter(char) == true {
			return false
		}
	}
	return true
}
