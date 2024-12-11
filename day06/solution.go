// Advent of Code, 2024, Day 6
//
// https://adventofcode.com/2024/day/6
//
// Part 1: Naive solution, just follow the guard step by step
package day06

import (
	"advent/util"
	"bufio"
	"fmt"
)

type Direction int

const NewPosition = '.'
const VisitedPosition = 'X'

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

func PartOneAnswer(filepath string) (int, error) {
	labMap, guard, err := getLabMapAndGuard(filepath)
	if err != nil {
		return 0, err
	}
	steps, err := trackGuard(labMap, guard)
	return steps, err
}

func PartTwoAnswer(filepath string) (int, error) {
	return 0, nil
}

// getLabMapAndGuard returns the lab map and the guard's position in the map,
// or an error if the file cannot be processed. The file is in the format
// described in the prompt.
func getLabMapAndGuard(filename string) (util.Matrix[rune], util.Position, error) {
	labMap := util.NewMatrix[rune]()
	guardPosition := util.Position{}
	err := util.ProcessFile(filename, func(s *bufio.Scanner) error {
		for s.Scan() {
			line := s.Text()
			row := make([]rune, 0)
			for j, r := range line {
				row = append(row, r)
				if isGuard(r) {
					guardPosition = util.Position{Row: len(labMap), Col: j}
				}
			}
			labMap = append(labMap, row)
		}
		return nil
	})
	return labMap, guardPosition, err
}

// trackGuard returns the number of steps the guard walks in the map before
// leaving the map
func trackGuard(labMap util.Matrix[rune], guardPos util.Position) (int, error) {
	steps := 0
	var err error
	var isNewPos bool
	for labMap.PosInBounds(guardPos) {
		guardPos, isNewPos, err = moveToNextPosition(labMap, guardPos)
		if err != nil {
			return 0, err
		}
		if isNewPos {
			steps++
		}
	}
	return steps, nil
}

// isGuard returns true if r is '^', 'v', '<', or '>'
func isGuard(r rune) bool {
	_, ok := guardPositions[r]
	return ok
}

// moveToNextPosition moves the guard to the next position in the labMap. If the
// guard encounters an obstacle, she turns right and steps one forward. Otherwise,
// she steps one forward in the direction she is currently moving. It returns true
// if and only if the returned position has never been visited by the guard.
func moveToNextPosition(labMap util.Matrix[rune], guardPos util.Position) (util.Position, bool, error) {
	guard := labMap.Get(guardPos)
	guardDirection, err := getGuardDirection(guard)
	if err != nil {
		return guardPos, false, err
	}
	nextPosition, err := getNextPosition(guardPos, guardDirection)
	if err != nil {
		return guardPos, false, err
	}
	labMap.Set(guardPos, VisitedPosition)
	isNewPosition := true
	if labMap.PosInBounds(nextPosition) {
		if isObstacle(labMap.Get(nextPosition)) {
			newDirection, err := getRightTurn(guard)
			if err != nil {
				return guardPos, false, err
			}
			labMap.Set(guardPos, newDirection)
			return moveToNextPosition(labMap, guardPos)
		} else {
			isNewPosition = labMap.Get(nextPosition) == NewPosition
			labMap.Set(nextPosition, guard)
		}
	}
	return nextPosition, isNewPosition, nil
}

// isObstacle returns true if r is not a '.'
func isObstacle(r rune) bool {
	return r != NewPosition && r != VisitedPosition
}

// getGuardDirection returns the direction of the guard represented by r, or an
// error if r is not a guard.
func getGuardDirection(r rune) (Direction, error) {
	position, ok := guardPositions[r]
	if !ok {
		return -1, fmt.Errorf("rune %c is not a guard", r)
	}
	return position, nil
}

func getNextPosition(currentPos util.Position, d Direction) (util.Position, error) {
	switch d {
	case Up:
		return util.Position{Row: currentPos.Row - 1, Col: currentPos.Col}, nil
	case Right:
		return util.Position{Row: currentPos.Row, Col: currentPos.Col + 1}, nil
	case Down:
		return util.Position{Row: currentPos.Row + 1, Col: currentPos.Col}, nil
	case Left:
		return util.Position{Row: currentPos.Row, Col: currentPos.Col - 1}, nil
	default:
		return util.Position{}, fmt.Errorf("invalid direction %d", d)
	}

}

// getRightTurn returns the direction the guard should turn to if it encounters
// an obstacle.
func getRightTurn(r rune) (rune, error) {
	switch r {
	case '^':
		return '>', nil
	case '>':
		return 'v', nil
	case 'v':
		return '<', nil
	case '<':
		return '^', nil
	default:
		return ' ', fmt.Errorf("rune %c is not a guard", r)
	}
}

// Print prints the matrix to the console.
func PrintMatrix(m util.Matrix[rune]) {
	for _, row := range m {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println()
	}
	fmt.Println()
}
