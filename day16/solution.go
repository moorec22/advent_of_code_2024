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
)

const MoveCost = 1
const TurnCost = 1000

const WallRune = '#'

// MazeCell represents a cell in the maze in the search map. It has a visited
// flag, and the cost to reach that cell.
type MazeCell struct {
	sym     rune
	cost    int
	visited bool
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
	s.findLeastCostPathHelper(mazeSearch, s.start, util.RightDirection, 0)
	return mazeSearch.Get(s.end).cost
}

// findLeastCostHelper fills in mazeSearch with the least cost to reach each cell.
func (s *Day16Solution) findLeastCostPathHelper(mazeSearch util.Matrix[MazeCell], start, dir *util.Vector, cost int) {
	cell := mazeSearch.Get(start)
	if cell.visited && cell.cost <= cost {
		return
	}
	mazeSearch.Set(start, MazeCell{cell.sym, cost, true})
	neighbors := s.getNeighbors(mazeSearch, start, dir)
	for _, n := range neighbors {
		s.findLeastCostPathHelper(mazeSearch, n.pos, n.dir, cost+n.cost)
	}
}

// getNeighbors returns the neighbors of the current position that can be moved
// to.
func (s *Day16Solution) getNeighbors(mazeSearch util.Matrix[MazeCell], pos, curDir *util.Vector) []MoveInfo {
	neighbors := make([]MoveInfo, 0)
	for _, d := range util.SimpleDirections {
		newPos := pos.Add(d)
		if mazeSearch.Get(newPos).sym != WallRune {
			cost := MoveCost
			if d != curDir {
				cost += TurnCost
			}
			neighbors = append(neighbors, MoveInfo{cost, newPos, d})
		}
	}
	return neighbors
}

// getMazeSearch returns a copy of the maze with the cost to reach each cell.
// The initial MazeCell has a cost of -1, and visited set to false.
func (s *Day16Solution) getMazeSearch(maze util.Matrix[rune]) util.Matrix[MazeCell] {
	searchMap := make(util.Matrix[MazeCell], len(maze))
	for i, row := range maze {
		searchMap[i] = make([]MazeCell, len(row))
		for j, sym := range row {
			searchMap[i][j] = MazeCell{sym, -1, false}
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
