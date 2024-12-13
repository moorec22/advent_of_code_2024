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

var unitVectors = map[Direction]util.Position{
	Up:    {Row: -1, Col: 0},
	Right: {Row: 0, Col: 1},
	Down:  {Row: 1, Col: 0},
	Left:  {Row: 0, Col: -1},
}

// DistanceMap maps each direction to the distance to the next obstacle. -1
// indicates that there is no obstacle in that direction.
type DistanceMap map[Direction]int
