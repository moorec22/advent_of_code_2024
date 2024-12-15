// Constants and Types for day 6 solution
package day06

import "advent/util"

type Direction int

const Empty = '.'
const Obstacle = '#'

const (
	Up    = iota
	Right = iota
	Down  = iota
	Left  = iota
)

var guardPositions = map[rune]Direction{
	'^': Up,
	'>': Right,
	'v': Down,
	'<': Left,
}

var unitVectors = map[Direction]util.Vector{
	Up:    {X: -1, Y: 0},
	Right: {X: 0, Y: 1},
	Down:  {X: 1, Y: 0},
	Left:  {X: 0, Y: -1},
}

// DistanceMap maps each direction to the distance to the next obstacle. -1
// indicates that there is no obstacle in that direction.
type DistanceMap map[Direction]int
