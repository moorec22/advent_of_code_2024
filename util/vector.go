package util

import (
	"strconv"
	"strings"
)

type Vector struct {
	X, Y int
}

func NewVector(x, y int) *Vector {
	return &Vector{X: x, Y: y}
}

// parseVector returns a Vector from a string in the format x,y. If the string
// is not in the correct format, an error is returned.
func ParseVector(s string) (*Vector, error) {
	parts := strings.Split(s, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return &Vector{}, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return &Vector{}, err
	}
	return NewVector(x, y), nil
}

func (v *Vector) Add(other *Vector) *Vector {
	return &Vector{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v *Vector) ScalarMultiply(scalar int) *Vector {
	return &Vector{X: v.X * scalar, Y: v.Y * scalar}
}

func (v *Vector) Modulo(other *Vector) *Vector {
	return &Vector{X: v.X % other.X, Y: v.Y % other.Y}
}

func (v *Vector) MathModulo(other *Vector) *Vector {
	return &Vector{X: MathModulo(v.X, other.X), Y: MathModulo(v.Y, other.Y)}
}

func (v *Vector) Negate() *Vector {
	return NewVector(-v.X, -v.Y)
}

// Unit returns the smallest manhattan distance with the same direction as
// the original.
func (v *Vector) Unit() *Vector {
	gcd := GreatestCommonDivisor(IntAbs(v.X), IntAbs(v.Y))
	return NewVector(v.X/gcd, v.Y/gcd)
}

func (v *Vector) GetManhattanDistance(other *Vector) *Vector {
	return NewVector(v.X-other.X, v.Y-other.Y)
}

func (v *Vector) AddManhattanDistance(distance *Vector) *Vector {
	return &Vector{X: v.X + distance.X, Y: v.Y + distance.Y}
}
