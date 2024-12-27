// Advent of Code, 2024, Day 20
//
// https://adventofcode.com/2024/day/20
//
// Part 1: I decided to preprocess the solution by running a djikstra's search
// for the shortest path from the end to each cell reachable, so we can run
// the following algorithm:
//   - for each cell 'cell':
//   - for each reachable cell 'nextCell' with a manhattan distance of <= 2
//   - if nextCell.distanceToEnd - cell.distanceToEnd > 100:
//   - shortcuts++
//
// Part 2: With the solution of part 1, all we have to do is change '2' to
// a variable, and set it to 20 for part 2.
package day20

import (
	"advent/util"
)

const ShortcutThreshold = 100
const PartOneCheatTime = 2
const PartTwoCheatTime = 20
const WallCell = '#'

type RacetrackCell struct {
	sym           rune
	distanceToEnd int
	pos           *util.Vector
}

func (c RacetrackCell) Compare(other RacetrackCell) int {
	return c.distanceToEnd - other.distanceToEnd
}

type Day20Solution struct {
	racetrack  util.Matrix[rune]
	start, end *util.Vector
}

func NewDay20Solution(filename string) (*Day20Solution, error) {
	racetrack, err := util.ParseMatrixFromFile(filename, func(r rune) rune {
		return r
	})
	start, end := findStartAndEnd(racetrack)
	return &Day20Solution{racetrack, start, end}, err
}

func (s *Day20Solution) PartOneAnswer() (int, error) {
	racetrackSearch := s.getRacetrackSearch(s.racetrack)
	s.shortestPathsToEnd(racetrackSearch, s.end)
	return s.getShortcutCount(racetrackSearch, PartOneCheatTime, ShortcutThreshold), nil
}

func (s *Day20Solution) PartTwoAnswer() (int, error) {
	racetrackSearch := s.getRacetrackSearch(s.racetrack)
	s.shortestPathsToEnd(racetrackSearch, s.end)
	return s.getShortcutCount(racetrackSearch, PartTwoCheatTime, ShortcutThreshold), nil
}

func (s *Day20Solution) shortestPathsToEnd(racetrack util.Matrix[RacetrackCell], end *util.Vector) {
	visited := make(map[util.Vector]bool)
	toVisit := util.NewArrayPriorityQueue[RacetrackCell]()
	endCell := racetrack.Get(end)
	endCell.distanceToEnd = 0
	racetrack.Set(end, endCell)
	toVisit.Insert(endCell)
	for !toVisit.IsEmpty() {
		cell := toVisit.Remove()
		if visited[*cell.pos] {
			continue
		}
		visited[*cell.pos] = true
		for neighborPos := range s.getEmptyNeighbors(racetrack, cell.pos) {
			neighbor := racetrack.Get(neighborPos)
			if neighbor.distanceToEnd == -1 || cell.distanceToEnd+1 < neighbor.distanceToEnd {
				neighbor.distanceToEnd = cell.distanceToEnd + 1
				racetrack.Set(neighbor.pos, neighbor)
				toVisit.Insert(neighbor)
			}
		}
	}
}

// getShortcutCount takes a racetrack, and a threshold to consider, and returns
// all the shortcuts that save at least threshold picoseconds
func (s *Day20Solution) getShortcutCount(racetrack util.Matrix[RacetrackCell], cheatTime, threshold int) int {
	count := 0
	for _, row := range racetrack {
		for _, cell := range row {
			if cell.sym != WallCell {
				possibleCheatCells := s.getCheatMoves(racetrack, cell.pos, cheatTime)
				for nextCell, distanceToCell := range possibleCheatCells {
					timeSaved := cell.distanceToEnd - nextCell.distanceToEnd - distanceToCell
					if nextCell.sym != WallCell && timeSaved >= threshold {
						count++
					}
				}
			}
		}
	}
	return count
}

// getEmtpyNeighbors takes a racetrack, and a position to consider. It returns the neighbors with
// empty squares.
func (s *Day20Solution) getEmptyNeighbors(racetrack util.Matrix[RacetrackCell], pos *util.Vector) map[*util.Vector]bool {
	neighbors := make(map[*util.Vector]bool)
	for _, d := range util.SimpleDirections {
		neighbor := racetrack.Get(pos.Add(d))
		if neighbor.sym != WallCell {
			neighbors[neighbor.pos] = true
		}
	}
	return neighbors
}

// getCheatMoves takes a racetrack, a position, and a cheat time to consider.
// It returns the positions of all positions that may be considered for a cheat
// that lasts cheatTime picoseconds.
func (s *Day20Solution) getCheatMoves(racetrack util.Matrix[RacetrackCell], pos *util.Vector, cheatTime int) map[RacetrackCell]int {
	neighbors := make(map[RacetrackCell]int)
	for i := pos.X - cheatTime; i <= pos.X+cheatTime; i++ {
		for j := pos.Y - cheatTime; j <= pos.Y+cheatTime; j++ {
			cellPos := util.NewVector(i, j)
			if !pos.Equals(cellPos) && racetrack.PosInBounds(cellPos) {
				cell := racetrack.Get(cellPos)
				distance := cell.pos.GetManhattanDistance(pos).GetManhattanMagnitude()
				if cell.sym != WallCell && distance <= cheatTime {
					neighbors[cell] = distance
				}
			}
		}
	}
	return neighbors
}

func (s *Day20Solution) getRacetrackSearch(racetrack util.Matrix[rune]) util.Matrix[RacetrackCell] {
	racetrackSearch := util.NewMatrix[RacetrackCell]()
	for i, row := range racetrack {
		racetrackSearch = append(racetrackSearch, make([]RacetrackCell, len(row)))
		for j, cell := range row {
			racetrackSearch[i][j] = RacetrackCell{cell, -1, util.NewVector(i, j)}
		}
	}
	return racetrackSearch
}

// findStartAndEnd finds the symbols S and E on the racetrack. It then returns
// the positions of both in the form (startVector, endVector)
func findStartAndEnd(racetrack util.Matrix[rune]) (*util.Vector, *util.Vector) {
	var start, end *util.Vector
	for i := range racetrack {
		for j, cell := range racetrack[i] {
			if cell == 'S' {
				start = util.NewVector(i, j)
			} else if cell == 'E' {
				end = util.NewVector(i, j)
			}
		}
	}
	return start, end
}
