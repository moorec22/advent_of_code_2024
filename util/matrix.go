// A matrix is a two-dimensional array of values. This file contains a generic
// implementation of a matrix, as well as some utility functions for working with
// matrices.
package util

import "fmt"

type Matrix[T any] [][]T

func NewMatrix[T any]() Matrix[T] {
	return make([][]T, 0)
}

// Get returns the value at the given position in the matrix.
func (m Matrix[T]) Get(pos Position) T {
	return m[pos.Row][pos.Col]
}

// Set sets the value at the given position in the matrix.
func (m Matrix[T]) Set(pos Position, val T) {
	m[pos.Row][pos.Col] = val
}

// PosInBounds returns true if the given position is within the bounds of the matrix.
func (m Matrix[T]) PosInBounds(pos Position) bool {
	return pos.Row >= 0 && pos.Row < len(m) && pos.Col >= 0 && pos.Col < len((m)[0])
}

// Print prints the matrix to the console.
func (m Matrix[T]) Print(toString func(T) string) {
	for _, row := range m {
		for _, val := range row {
			fmt.Print(toString(val))
		}
		fmt.Println()
	}
	fmt.Println()
}
