package diffsquares

/*
Difference finds the difference between the square of the sum and the sum of the squares of the first N natural numbers
*/
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

/*
SquareOfSum calculates the square of the sum of the first N natural numbers.
*/
func SquareOfSum(n int) int {
	squareOfSum := n * (n + 1) / 2
	squareOfSum *= squareOfSum
	return squareOfSum
}

/*
SumOfSquares calculates the sum of the squares of the first n natural numbers.
*/
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}
