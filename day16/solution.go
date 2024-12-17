// Advent of Code, 2024, Day 16
//
// https://adventofcode.com/2024/day/16
//
// Part 1: I decided to represent the searchable maze as a 2D vector of
// MazeCell. With this representation, each cell stores the best found cost to
// reach it, and whether it has been visited. The search algorithm is a simple
// recursive depth-first search that fills in the cost to reach each cell,
// moving on if the found cost to reach the cell is less than the current cost.
// We then return the least cost found to reach the end cell.
package day16

import (
	"advent/util"
	"fmt"
)

const MoveCost = 1
const TurnCost = 1000

const WallRune = '#'

// MazeCell represents a cell in the maze in the search map. It has a visited
// flag, and the cost to reach that cell.
type MazeCell struct {
	sym rune
	// the cost to reach this cell facing the direction of the key
	dirCost map[*util.Vector]int
	visited bool
	// the final cost, only to be used by the end cell.
	endCost int
}

// MoveInfo represents the position to move to, the direction it is in, and
// the cost to move there.
type MoveInfo struct {
	cost     int
	pos, dir *util.Vector
}

type Day16Solution struct {
	maze       util.Matrix[rune]
	start, end *util.Vector
}

func NewDay16Solution(filename string) (*Day16Solution, error) {
	maze, err := util.ParseMatrixFromFile[rune](filename, func(r rune) rune {
		return r
	})
	start, end := getStartAndEnd(maze)
	return &Day16Solution{maze, start, end}, err
}

func (s *Day16Solution) PartOneAnswer() (int, error) {
	return s.findLeastCostPath(), nil
}

func (s *Day16Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}

// findLeastCostPath returns the least cost to reach the end cell from the start.
func (s *Day16Solution) findLeastCostPath() int {
	mazeSearch := s.getMazeSearch(s.maze)
	s.findLeastCostPathHelper(mazeSearch, s.start, s.end, util.RightDirection, 0)
	return mazeSearch.Get(s.end).endCost
}

// findLeastCostHelper fills in mazeSearch with the least cost to reach each cell.
func (s *Day16Solution) findLeastCostPathHelper(mazeSearch util.Matrix[MazeCell], start, end, dir *util.Vector, cost int) error {
	cell := mazeSearch.Get(start)
	if start.Equals(end) {
		if !cell.visited || cost < cell.endCost {
			cell.endCost = cost
			mazeSearch.Set(start, cell)
		}
		return nil
	}
	neighbors, err := s.getNeighbors(mazeSearch, start, dir)
	if err != nil {
		return nil
	}
	for newDir, newPos := range neighbors {
		newCost := cost
		if newDir != dir {
			newCost += TurnCost
		}
		v, ok := cell.dirCost[newDir]
		if !cell.visited || !ok || newCost < v {
			cell.dirCost[newDir] = newCost
			cell.visited = true
			mazeSearch.Set(start, cell)
			s.findLeastCostPathHelper(mazeSearch, newPos, end, newDir, newCost+MoveCost)
		}
	}
	return nil
}

// getNeighbors returns the neighbors of the current position that can be moved
// to. A position can be moved to if it is an empty position in front of, to the
// right, or to the left of pos facing curDir. An error is returend if curDir is
// not a simple direction.
func (s *Day16Solution) getNeighbors(mazeSearch util.Matrix[MazeCell], pos, curDir *util.Vector) (map[*util.Vector]*util.Vector, error) {
	validDirections, err := s.getValidDirections(curDir)
	if err != nil {
		return nil, err
	}
	neighbors := make(map[*util.Vector]*util.Vector)
	for _, d := range validDirections {
		newPos := pos.Add(d)
		if mazeSearch.Get(newPos).sym != WallRune {
			neighbors[d] = newPos
		}
	}
	return neighbors, nil
}

// getValidDirections returns the valid directions that can be moved to from the
// current direction. Valid directions are all directions other than the cardinal
// opposite of curDirection. An error is returned if curDirection is not a simple
// direction.
func (s *Day16Solution) getValidDirections(curDirection *util.Vector) ([]*util.Vector, error) {
	validDirections := make([]*util.Vector, 0)
	oppositeDirection, err := s.getOppositeDirection(curDirection)
	if err != nil {
		return nil, err
	}
	for _, d := range util.SimpleDirections {
		if d != oppositeDirection {
			validDirections = append(validDirections, d)
		}
	}
	return validDirections, nil
}

// getOppositeDirection returns the opposite direction of the current
// direction. If curDirection is not a simple direction, an error is
// returned.
func (s *Day16Solution) getOppositeDirection(curDirection *util.Vector) (*util.Vector, error) {
	switch curDirection {
	case util.UpDirection:
		return util.DownDirection, nil
	case util.DownDirection:
		return util.UpDirection, nil
	case util.LeftDirection:
		return util.RightDirection, nil
	case util.RightDirection:
		return util.LeftDirection, nil
	default:
		return nil, fmt.Errorf("invalid direction: %v", curDirection)
	}
}

// getMazeSearch returns a copy of the maze with the cost to reach each cell.
// The initial MazeCell has a cost of -1, and visited set to false.
func (s *Day16Solution) getMazeSearch(maze util.Matrix[rune]) util.Matrix[MazeCell] {
	searchMap := make(util.Matrix[MazeCell], len(maze))
	for i, row := range maze {
		searchMap[i] = make([]MazeCell, len(row))
		for j, sym := range row {
			searchMap[i][j] = MazeCell{sym, map[*util.Vector]int{}, false, -1}
		}
	}
	return searchMap
}

// getStartAndEnd returns the start and end positions in the maze.
func getStartAndEnd(maze util.Matrix[rune]) (*util.Vector, *util.Vector) {
	var start, end *util.Vector
	for i, row := range maze {
		for j, cell := range row {
			if cell == 'S' {
				start = util.NewVector(i, j)
			} else if cell == 'E' {
				end = util.NewVector(i, j)
			}
		}
	}
	return start, end
}
