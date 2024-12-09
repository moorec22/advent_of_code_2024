package util

// intAbs returns the absolute value of a given integer.
func IntAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
