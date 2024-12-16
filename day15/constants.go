package day15

import "advent/util"

const RobotRune = '@'

var RuneDirections = map[rune]*util.Vector{
	'^': util.UpDirection,
	'v': util.DownDirection,
	'>': util.RightDirection,
	'<': util.LeftDirection,
}

const (
	EmptyRune        = '.'
	WallRune         = '#'
	BoxRune          = 'O'
	WideBoxLeftRune  = '['
	WideBoxRightRune = ']'
)
