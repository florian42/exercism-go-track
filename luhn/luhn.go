package luhn

import (
	"strconv"
	"strings"
)

// Valid given a number determine whether or not it is valid per the Luhn formula.
func Valid(input string) bool {
	input = spaceFieldsJoin(input)
	if len(input) <= 1 {
		return false
	}
	if _, err := strconv.Atoi(input); err != nil {
		return false
	}
	sum := 0
	var alternate bool
	for pos := len(input) - 1; pos > -1; pos-- {
		num := int(input[pos] - 0x30)
		if alternate {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		alternate = !alternate
		sum += num
	}
	return sum%10 == 0
}

func spaceFieldsJoin(str string) string {
	return strings.Join(strings.Fields(str), "")
}
