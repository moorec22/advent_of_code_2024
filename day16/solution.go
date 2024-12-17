// Advent of Code, 2024, Day 16
//
// https://adventofcode.com/2024/day/16
//
// Part 1: I decided to start with a simple recursive search, that solved the
// test cases but was not preformant. I then moved on to optimize the solution.
// Namely, if we keep track of the least cost we've found to get to the end, we
// can save ourselves the trouble of trying a path we already know to be more
// expensive. This by definition will truncate the long running paths, and
// greatly improve preformance time. With this optimization, I was able to find
// the answer.
package day16

import (
	"advent/util"
	"fmt"
)

const MoveCost = 1
const TurnCost = 1000

const WallRune = '#'

type CellInfo struct {
	sym        rune
	foundCosts map[util.Vector]int
}

type Day16Solution struct {
	maze       util.Matrix[CellInfo]
	start, end *util.Vector
	solved     bool
}

func NewDay16Solution(filename string) (*Day16Solution, error) {
	maze, err := util.ParseMatrixFromFile[rune](filename, func(r rune) rune {
		return r
	})
	start, end := getStartAndEnd(maze)
	mazeInfo := getMazeInfoMap(maze)
	return &Day16Solution{mazeInfo, start, end, false}, err
}

func (s *Day16Solution) PartOneAnswer() (int, error) {
	return s.findLeastCost()
}

func (s *Day16Solution) PartTwoAnswer() (int, error) {
	return s.findCellCount()
}

// findLeastCost returns the least cost to reach the end cell from the start.
func (s *Day16Solution) findLeastCost() (int, error) {
	if !s.solved {
		err := s.findLeastCostHelper(s.maze, s.start, s.end, util.RightDirection, make(map[util.Vector]bool), 0)
		if err != nil {
			return -1, err
		}
		s.solved = true
	}
	cost, err := s.mapMin(s.maze.Get(s.end).foundCosts)
	return cost, err
}

// findCellCount returns the number of cells found on any of the least cost paths.
func (s *Day16Solution) findCellCount() (int, error) {
	if !s.solved {
		err := s.findLeastCostHelper(s.maze, s.start, s.end, util.RightDirection, make(map[util.Vector]bool), 0)
		if err != nil {
			return -1, err
		}
	}
	return 0, nil
}

// findLeastCostHelper fills in mazeSearch with the least cost to reach each
// cell, facing each direction. It returns an error if the search fails.
func (s *Day16Solution) findLeastCostHelper(mazeSearch util.Matrix[CellInfo], start, end, dir *util.Vector, visited map[util.Vector]bool, curCost int) error {
	foundCosts := mazeSearch.Get(start).foundCosts
	if start.Equals(end) {
		if foundCost, err := s.mapMin(foundCosts); err != nil || curCost < foundCost {
			mazeSearch.Get(start).foundCosts[*dir] = curCost
		}
		return nil
	} else if foundCost, ok := foundCosts[*dir]; ok && curCost >= foundCost {
		return nil
	} else if visited[*start] {
		return nil
	}
	mazeSearch.Get(start).foundCosts[*dir] = curCost
	// let's first try moving forward
	newPos := start.Add(dir)
	if mazeSearch.Get(newPos).sym != WallRune {
		visited[*start] = true
		err := s.findLeastCostHelper(mazeSearch, newPos, end, dir, visited, curCost+MoveCost)
		if err != nil {
			return err
		}
		visited[*start] = false
	}
	// and then try turning
	leftTurn, err := s.getLeftTurn(dir)
	if err != nil {
		return err
	}
	err = s.findLeastCostHelper(mazeSearch, start, end, leftTurn, visited, curCost+TurnCost)
	if err != nil {
		return err
	}
	rightTurn, err := s.getRightTurn(dir)
	if err != nil {
		return err
	}
	err = s.findLeastCostHelper(mazeSearch, start, end, rightTurn, visited, curCost+TurnCost)
	if err != nil {
		return err
	}
	return nil
}

// getLeftTurn returns a new direction that is the result of turning left from
// the current direction. An error is returned if the current direction is not
// a simple direction.
func (s *Day16Solution) getLeftTurn(dir *util.Vector) (*util.Vector, error) {
	switch dir {
	case util.UpDirection:
		return util.LeftDirection, nil
	case util.DownDirection:
		return util.RightDirection, nil
	case util.LeftDirection:
		return util.DownDirection, nil
	case util.RightDirection:
		return util.UpDirection, nil
	default:
		return nil, fmt.Errorf("invalid direction: %v", dir)
	}
}

// getRightTurn returns a new direction that is the result of turning right from
// the current direction. An error is returned if the current direction is not
// a simple direction.
func (s *Day16Solution) getRightTurn(dir *util.Vector) (*util.Vector, error) {
	switch dir {
	case util.UpDirection:
		return util.RightDirection, nil
	case util.DownDirection:
		return util.LeftDirection, nil
	case util.LeftDirection:
		return util.UpDirection, nil
	case util.RightDirection:
		return util.DownDirection, nil
	default:
		return nil, fmt.Errorf("invalid direction: %v", dir)
	}
}

// mapMin returns the minimum value found in the map. It returns an error if
// the map is empty.
func (s *Day16Solution) mapMin(m map[util.Vector]int) (int, error) {
	if len(m) == 0 {
		return -1, fmt.Errorf("no minimum found")
	}
	min := -1
	for _, cost := range m {
		if min < 0 || cost < min {
			min = cost
		}
	}
	return min, nil
}

// getMazeInfoMap returns a matrix of CellInfo structs, one for each cell in
// the maze. The CellInfo struct contains the rune of the cell, and a map of
// directions to the least cost found to reach the end from that cell facing
// that direction.
func getMazeInfoMap(maze util.Matrix[rune]) util.Matrix[CellInfo] {
	mazeInfo := make(util.Matrix[CellInfo], len(maze))
	for i, row := range maze {
		mazeInfo[i] = make([]CellInfo, len(row))
		for j, cell := range row {
			mazeInfo[i][j] = CellInfo{cell, make(map[util.Vector]int)}
		}
	}
	return mazeInfo
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

func printPath(maze util.Matrix[rune], path map[util.Vector]bool) {
	for i, row := range maze {
		for j, cell := range row {
			pos := util.NewVector(i, j)
			if path[*pos] {
				fmt.Print("X")
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
}
