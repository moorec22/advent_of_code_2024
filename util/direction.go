package util

// TODO: Use direction constants in other solutions
var (
	UpDirection    = NewPosition(-1, 0)
	DownDirection  = NewPosition(1, 0)
	RightDirection = NewPosition(0, 1)
	LeftDirection  = NewPosition(0, -1)
)

var (
	UpRightDirection   = NewPosition(-1, 1)
	DownRightDirection = NewPosition(1, 1)
	DownLeftDirection  = NewPosition(1, -1)
	UpLeftDirection    = NewPosition(-1, -1)
)

// Directions are ordered in a clockwise manner starting from UpDirection.
var SimpleDirections = []Position{UpDirection, RightDirection, DownDirection, LeftDirection}
var AllDirections = []Position{UpDirection, UpRightDirection, RightDirection, DownRightDirection,
	DownDirection, DownLeftDirection, LeftDirection, UpLeftDirection}
