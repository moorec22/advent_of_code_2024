// A matrix is a two-dimensional array of values. This file contains a generic
// implementation of a matrix, as well as some utility functions for working with
// matrices.
package util

import "fmt"

type Matrix[T any] [][]T

func NewMatrix[T any]() Matrix[T] {
	return make([][]T, 0)
}

// Get returns the value at the given vector in the matrix.
func (m Matrix[T]) Get(pos *Vector) T {
	return m[pos.X][pos.Y]
}

// Set sets the value at the given vector in the matrix.
func (m Matrix[T]) Set(pos *Vector, val T) {
	m[pos.X][pos.Y] = val
}

// PosInBounds returns true if the given vector is within the bounds of the matrix.
func (m Matrix[T]) PosInBounds(pos *Vector) bool {
	return pos.X >= 0 && pos.X < len(m) && pos.Y >= 0 && pos.Y < len((m)[0])
}

// Copy copies the matrix and returns a new matrix with the same values. It
// does not do a deep copy of the values T in each cell.
func (m Matrix[T]) Copy() Matrix[T] {
	newMatrix := make([][]T, len(m))
	for i, row := range m {
		newMatrix[i] = make([]T, len(row))
		copy(newMatrix[i], row)
	}
	return newMatrix
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
