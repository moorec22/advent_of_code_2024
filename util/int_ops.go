package util

// intAbs returns the absolute value of a given integer.
func IntAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// GreatestCommonDivisor returns the greatest common divisor of two integers.
func GreatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}
	return GreatestCommonDivisor(b, a%b)
}
