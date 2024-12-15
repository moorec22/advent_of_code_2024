package util

// TODO: Use direction constants in other solutions
var (
	UpDirection    = NewVector(-1, 0)
	DownDirection  = NewVector(1, 0)
	RightDirection = NewVector(0, 1)
	LeftDirection  = NewVector(0, -1)
)

var (
	UpRightDirection   = NewVector(-1, 1)
	DownRightDirection = NewVector(1, 1)
	DownLeftDirection  = NewVector(1, -1)
	UpLeftDirection    = NewVector(-1, -1)
)

// Directions are ordered in a clockwise manner starting from UpDirection.
var SimpleDirections = []Vector{UpDirection, RightDirection, DownDirection, LeftDirection}
var AllDirections = []Vector{UpDirection, UpRightDirection, RightDirection, DownRightDirection,
	DownDirection, DownLeftDirection, LeftDirection, UpLeftDirection}
