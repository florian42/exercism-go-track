package grains

import (
	"errors"
)

//Total calculates the number of grains of wheat on a chessboard given that the number on each square doubles.
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		square, err := Square(i)
		if err == nil {
			total += square
		}
	}
	return total
}

// Square calculates the number of grains on the xth square
func Square(x int) (uint64, error) {
	if x <= 0 || x > 64 {
		return 0, errors.New("input of 0 and smaller and input bigger than 64 is not allowed")
	}
	return 1 << (uint(x) - 1), nil
}
