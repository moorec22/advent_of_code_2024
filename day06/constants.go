// Constants and Types for day 6 solution
package day06

import "advent/util"

const Empty = '.'
const Obstacle = '#'

var guardDirections = map[rune]util.Vector{
	'^': util.UpDirection,
	'>': util.RightDirection,
	'v': util.DownDirection,
	'<': util.LeftDirection,
}

// DistanceMap maps each direction to the distance to the next obstacle. -1
// indicates that there is no obstacle in that direction.
type DistanceMap map[util.Vector]int
