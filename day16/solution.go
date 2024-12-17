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
	return s.findLeastCost()
}

func (s *Day16Solution) PartTwoAnswer() (int, error) {
	return s.findCellCount()
}

// findLeastCost returns the least cost to reach the end cell from the start.
func (s *Day16Solution) findLeastCost() (int, error) {
	mazeInfo := s.getMazeInfoMap()
	err := s.findLeastCostHelper(mazeInfo, s.start, s.end, util.RightDirection, make(map[util.Vector]bool), 0)
	if err != nil {
		return -1, err
	}
	cost, err := s.mapMin(mazeInfo.Get(s.end).foundCosts)
	return cost, err
}

// findCellCount returns the number of cells found on any of the least cost paths.
func (s *Day16Solution) findCellCount() (int, error) {
	mazeInfo := s.getMazeInfoMap()
	err := s.findLeastCostHelper(mazeInfo, s.start, s.end, util.RightDirection, make(map[util.Vector]bool), 0)
	return 0, err
}

// findLeastCostHelper fills in mazeSearch with the least cost to reach each
// cell, facing each direction. It returns an error if the search fails.
func (s *Day16Solution) findLeastCostHelper(mazeSearch util.Matrix[CellInfo], start, end, dir *util.Vector, visited map[util.Vector]bool, curCost int) error {
	foundCost, ok := mazeSearch.Get(start).foundCosts[*dir]
	if start.Equals(end) {
		if !ok || curCost < foundCost {
			mazeSearch.Get(start).foundCosts[*dir] = curCost
		}
		return nil
	} else if ok && curCost >= foundCost {
		return nil
	}
	mazeSearch.Get(start).foundCosts[*dir] = curCost
	// let's first try moving forward
	newPos := start.Add(dir)
	if mazeSearch.Get(newPos).sym != WallRune {
		err := s.findLeastCostHelper(mazeSearch, newPos, end, dir, visited, curCost+MoveCost)
		if err != nil {
			return err
		}
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

// getNeighbors returns the neighbors of the current position that can be moved
// to. A position can be moved to if it is an empty position in front of, to the
// right, or to the left of pos facing curDir. An error is returend if curDir is
// not a simple direction.
func (s *Day16Solution) getNeighbors(mazeSearch util.Matrix[CellInfo], pos, curDir *util.Vector) (map[*util.Vector]*util.Vector, error) {
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

// getMazeInfoMap returns a matrix of CellInfo structs, one for each cell in
// the maze. The CellInfo struct contains the rune of the cell, and a map of
// directions to the least cost found to reach the end from that cell facing
// that direction.
func (s *Day16Solution) getMazeInfoMap() util.Matrix[CellInfo] {
	mazeInfo := make(util.Matrix[CellInfo], len(s.maze))
	for i, row := range s.maze {
		mazeInfo[i] = make([]CellInfo, len(row))
		for j, cell := range row {
			mazeInfo[i][j] = CellInfo{cell, make(map[util.Vector]int)}
		}
	}
	return mazeInfo
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
