package util

type Vector struct {
	X, Y int
}

func NewVector(x, y int) Vector {
	return Vector{X: x, Y: y}
}

func (v *Vector) Add(other Vector) Vector {
	return Vector{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v *Vector) Negate() Vector {
	return NewVector(-v.X, -v.Y)
}

// Unit returns the smallest manhattan distance with the same direction as
// the original.
func (v *Vector) Unit() Vector {
	gcd := GreatestCommonDivisor(IntAbs(v.X), IntAbs(v.Y))
	return NewVector(v.X/gcd, v.Y/gcd)
}

func (v *Vector) GetManhattanDistance(other Vector) Vector {
	return NewVector(v.X-other.X, v.Y-other.Y)
}

func (v *Vector) AddManhattanDistance(distance Vector) Vector {
	return Vector{X: v.X + distance.X, Y: v.Y + distance.Y}
}
