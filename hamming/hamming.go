/*
Package hamming implements distance function to calculate the hamming distance.
*/
package hamming

import "errors"

/*
Distance comparing two DNA strands and counting how many of the nucleotides are
different from their equivalent in the other string.
*/
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("the Hamming distance is only defined for sequences of equal length")
	}
	sum := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			sum++
		}
	}
	return sum, nil
}
