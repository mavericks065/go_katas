package fizzbuzzSolution

import "strconv"

// constants to check if the input is modulo 3,5 or 15.
const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

// GetResult should return fizz, buzz or fizzbuzz depending
// on if the input is modulo 3, 5 or 15.
func GetResult(input int) string {
	if input%FIZZBUZZ == 0 {
		return "fizzbuzz"
	}
	if input%FIZZ == 0 {
		return "fizz"
	}
	if input%BUZZ == 0 {
		return "buzz"
	}
	return strconv.Itoa(input)
}
