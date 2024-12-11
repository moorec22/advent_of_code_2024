package util

type Matrix[T any] [][]T

type Position struct {
	Row, Col int
}

// PosInBounds returns true if the given position is within the bounds of the matrix.
func (m Matrix[T]) PosInBounds(pos Position) bool {
	return pos.Row >= 0 && pos.Row < len(m) && pos.Col >= 0 && pos.Col < len((m)[0])
}
