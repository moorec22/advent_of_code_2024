package util

type Vector struct {
	X, Y int
}

func NewVector(x, y int) Vector {
	return Vector{X: x, Y: y}
}

func (p *Vector) Add(other Vector) Vector {
	return Vector{X: p.X + other.X, Y: p.Y + other.Y}
}
