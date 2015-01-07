// Package damm supports the computation of a decimal checksum. It uses a
// method proposed by H. Michael Damm in 2004. The checksum doesn't change for
// leading zeroes.
//
// The function CheckDigit computes the check sum as string:
//
//   c, err := CheckDigit("12345678901")
//
// The function Validate checks whether the appended check digit is correct.
//
//   ok := Validate("123456789018")
//
// Information about the algorithm is available on Wikipedia.
//
// http://en.wikipedia.org/wiki/Damm_algorithm
//
package damm

import (
	"errors"
	"strconv"
)

// convert converts the digits from string into an int8 slice.
func convert(digits string) (a []int8, err error) {
	a = make([]int8, len(digits))
	for i, r := range digits {
		x := r - '0'
		if !(0 <= x && x <= 9) {
			return nil, errors.New(
				"digits strings must contain only digits")
		}
		a[i] = int8(x)
	}
	return a, nil
}

// quasi contains the quasi group used for computing the check digit.
var quasi = [10][10]int8{
	{0, 3, 1, 7, 5, 9, 8, 6, 4, 2},
	{7, 0, 9, 2, 1, 5, 4, 8, 6, 3},
	{4, 2, 0, 6, 8, 7, 1, 3, 5, 9},
	{1, 7, 5, 0, 9, 8, 3, 4, 2, 6},
	{6, 1, 2, 3, 0, 4, 5, 9, 7, 8},
	{3, 6, 7, 4, 2, 0, 9, 5, 8, 1},
	{5, 8, 6, 9, 7, 2, 0, 1, 3, 4},
	{8, 9, 4, 5, 3, 6, 2, 0, 1, 7},
	{9, 4, 3, 8, 6, 1, 7, 2, 0, 5},
	{2, 5, 8, 1, 4, 3, 6, 7, 9, 0},
}

// checkInt computes the check digit and returns it as integer.
func checkInt(a []int8) int {
	c := int8(0)
	for _, x := range a {
		c = quasi[c][x]
	}
	return int(c)
}

// CheckDigit computes the check digit and returns it as string. The function
// argument must only contain decimal digits.
func CheckDigit(digits string) (c string, err error) {
	a, err := convert(digits)
	if err != nil {
		return "", err
	}
	x := checkInt(a)
	c = strconv.Itoa(x)
	return
}

// Validate checks a number with the check digit appended. The function returns
// true only if the argument contains only decimal digits and the appended
// check digit is correct.
func Validate(digits string) bool {
	a, err := convert(digits)
	if err != nil {
		return false
	}
	return checkInt(a) == 0
}
