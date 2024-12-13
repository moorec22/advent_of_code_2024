// A distance between two points in a two dimensional grid, defined in two
// dimensions.
package util

type ManhattanDistance struct {
	X, Y int
}

func NewManhattanDistance(x, y int) ManhattanDistance {
	return ManhattanDistance{X: x, Y: y}
}

func (m *ManhattanDistance) Negate() ManhattanDistance {
	return NewManhattanDistance(-m.X, -m.Y)
}

// Unit returns the smallest manhattan distance with the same direction as
// the original.
func (m *ManhattanDistance) Unit() ManhattanDistance {
	gcd := GreatestCommonDivisor(IntAbs(m.X), IntAbs(m.Y))
	return NewManhattanDistance(m.X/gcd, m.Y/gcd)
}
