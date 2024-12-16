package util

// SliceSum returns the sum of all the integers in the slice.
func SliceSum(slice []int) int {
	sum := 0
	for _, i := range slice {
		sum += i
	}
	return sum
}

// SliceProduct returns the product of all the integers in the slice.
func SliceProduct(slice []int) int {
	product := 1
	for _, i := range slice {
		product *= i
	}
	return product
}
