// Advent of Code, 2024, Day 15
//
// https://adventofcode.com/2024/day/15
//
// Part 1: I decided to start by running the simulation, step by step.
//
// Part 2: This part ended up being more cumbersome than I thought. I did my
// best to restrict the complexity to canBeMovedTo, which returns true if
// a position can be moved to. It contains the logic for checking if a wide
// box, and all following boxes, can be moved.
package day15

import (
	"advent/util"
	"bufio"
	"fmt"
	"slices"
)

type Day15Solution struct {
	storageMap    util.Matrix[rune]
	instructions  []rune
	robotPosition *util.Vector
}

func NewDay15Solution(filename string) (*Day15Solution, error) {
	storageMap := make(util.Matrix[rune], 0)
	instructions := make([]rune, 0)
	err := util.ProcessFile(filename, func(scanner *bufio.Scanner) error {
		var err error
		storageMap, err = util.ParseMatrixFromScanner(scanner, func(r rune) rune {
			return r
		})
		if err != nil {
			return err
		}
		for scanner.Scan() {
			line := scanner.Text()
			for _, r := range line {
				instructions = append(instructions, r)
			}
		}
		return nil
	})
	robotPosition := findRobot(storageMap)
	return &Day15Solution{storageMap, instructions, robotPosition}, err
}

func (s *Day15Solution) PartOneAnswer() (int, error) {
	storageMap := s.storageMap.Copy()
	err := s.makeMoves(storageMap, s.robotPosition, s.instructions)
	if err != nil {
		return 0, err
	}
	gpsCoordinates := s.getGpsCoordinates(storageMap)
	return util.SliceSum(gpsCoordinates), nil
}

func (s *Day15Solution) PartTwoAnswer() (int, error) {
	widerMap, robotPos := s.widenMap(s.storageMap)
	err := s.makeMoves(widerMap, robotPos, s.instructions)
	if err != nil {
		return 0, err
	}
	gpsCoordinates := s.getGpsCoordinates(widerMap)
	return util.SliceSum(gpsCoordinates), nil
}

// makeMoves will make the moves specified by the moves slice, in order. It
// returns an error if there is no robot at robotPos.
func (s *Day15Solution) makeMoves(storageMap util.Matrix[rune], robotPos *util.Vector, moves []rune) error {
	var err error
	for _, move := range moves {
		robotPos, err = s.makeMove(storageMap, robotPos, move)
		if err != nil {
			return err
		}
	}
	return nil
}

// makeMove will move the robot at robotPos, returning an error if there is no
// robot at robotPos. The move made is a move on storageMap specified by the
// problem statement. storageMap is modified by the next move.
func (s *Day15Solution) makeMove(storageMap util.Matrix[rune], robotPos *util.Vector, move rune) (*util.Vector, error) {
	if storageMap.Get(robotPos) != RobotRune {
		return robotPos, fmt.Errorf("no robot at position %v", robotPos)
	}
	dir := RuneDirections[move]
	movePos := robotPos.Add(dir)
	if !s.canBeMovedTo(storageMap, movePos, dir) {
		return robotPos, nil
	}
	el := storageMap.Get(movePos)
	moveFunc, err := s.getMoveFunction(el)
	if err != nil {
		return robotPos, err
	}
	if moveFunc(storageMap, robotPos, dir) {
		robotPos = movePos
	}
	return robotPos, nil
}

// getMoveFunction returns the appropriate function for the rune r
func (s *Day15Solution) getMoveFunction(r rune) (func(util.Matrix[rune], *util.Vector, *util.Vector) bool, error) {
	switch r {
	case EmptyRune:
		return s.makeEmptyMove, nil
	case WallRune:
		return s.makeWallMove, nil
	case BoxRune, WideBoxLeftRune, WideBoxRightRune:
		return s.makeBoxMove, nil
	default:
		return nil, fmt.Errorf("no function for rune: %v", r)
	}
}

// makeBoxMove moves the object and the box as specified by the problem. It
// looks for an empty space at some point in the direction given. If it finds
// one, it moves the object and all boxes between the object and the empty space
// to that position. Otherwise, it does nothing. It returns true if a move was
// made. It assumes obPos is a moveable object. Otherwise, behavior is
// unspecified. If obPos is part of a wide box, it is assumed the move can be
// made. Otherwise, behavior is unspecified.
func (s *Day15Solution) makeBoxMove(storageMap util.Matrix[rune], objPos, dir *util.Vector) bool {
	nextPos := objPos.Add(dir)
	nextObj := storageMap.Get(nextPos)
	switch nextObj {
	case EmptyRune:
		return s.makeEmptyMove(storageMap, objPos, dir)
	case WallRune:
		return false
	case BoxRune:
		if s.makeBoxMove(storageMap, nextPos, dir) {
			return s.makeEmptyMove(storageMap, objPos, dir)
		}
	case WideBoxLeftRune, WideBoxRightRune:
		if slices.Contains(util.VerticalDirections, dir) {
			otherSidePos := nextPos.Add(s.wideBoxOtherSideDirection(nextObj))
			if s.makeBoxMove(storageMap, nextPos, dir) && s.makeBoxMove(storageMap, otherSidePos, dir) {
				return s.makeEmptyMove(storageMap, objPos, dir)
			}
		} else if s.makeBoxMove(storageMap, nextPos, dir) {
			return s.makeEmptyMove(storageMap, objPos, dir)
		}
	}
	return false
}

// canBeMovedTo returns true if:
// - storageMap[pos] is empty
// - storageMap[pos] is a box and can be moved in the direction dir
// - storageMap[pos] is a wide box and can be moved in the direction dir
func (s *Day15Solution) canBeMovedTo(storageMap util.Matrix[rune], pos, dir *util.Vector) bool {
	obj := storageMap.Get(pos)
	switch obj {
	case EmptyRune:
		return true
	case BoxRune:
		return s.canBeMovedTo(storageMap, pos.Add(dir), dir)
	case WideBoxLeftRune, WideBoxRightRune:
		if slices.Contains(util.VerticalDirections, dir) {
			otherSidePos := pos.Add(s.wideBoxOtherSideDirection(obj))
			return s.canBeMovedTo(storageMap, pos.Add(dir), dir) && s.canBeMovedTo(storageMap, otherSidePos.Add(dir), dir)
		} else {
			return s.canBeMovedTo(storageMap, pos.Add(dir), dir)
		}
	default:
		return false
	}
}

// makeEmptyMove moves obPos to nextPos, leaving obPos empty. It assumes obPos
// is moveable and nextPos is empty, otherwise behavior is unspecified.
func (s *Day15Solution) makeEmptyMove(storageMap util.Matrix[rune], obPos, dir *util.Vector) bool {
	nextPos := obPos.Add(dir)
	ob := storageMap.Get(obPos)
	storageMap.Set(obPos, EmptyRune)
	storageMap.Set(nextPos, ob)
	return true
}

// makeWallMove does nothing, as happens when the robot is facing a wall. It
// returns false.
func (s *Day15Solution) makeWallMove(storageMap util.Matrix[rune], obPos, dir *util.Vector) bool {
	return false
}

// getGpsCoordiantes returns a list of all GPS coordinates for boxes found in
// storageMap, as specified by the problem statement.
func (s *Day15Solution) getGpsCoordinates(storageMap util.Matrix[rune]) []int {
	gpsCoordinates := make([]int, 0)
	for i := range storageMap {
		for j := range storageMap[i] {
			if storageMap[i][j] == BoxRune || storageMap[i][j] == WideBoxLeftRune {
				gpsCoordinates = append(gpsCoordinates, 100*i+j)
			}
		}
	}
	return gpsCoordinates
}

// widenMap makes every element in the storage map twice as wide, as follows:
//
//	'#' -> '##'
//	'.' -> '..'
//	'O' -> '[]'
//	'@' -> '@.'
//
// It returns the map, and the new position of the robot.
func (s *Day15Solution) widenMap(storageMap util.Matrix[rune]) (util.Matrix[rune], *util.Vector) {
	widerMap := util.NewMatrix[rune]()
	var robotPosition *util.Vector
	for i, row := range storageMap {
		widerRow := make([]rune, len(row)*2)
		for j, cell := range row {
			switch cell {
			case EmptyRune, WallRune:
				widerRow[2*j] = cell
				widerRow[2*j+1] = cell
			case RobotRune:
				robotPosition = util.NewVector(i, 2*j)
				widerRow[2*j] = cell
				widerRow[2*j+1] = EmptyRune
			case BoxRune:
				widerRow[2*j] = '['
				widerRow[2*j+1] = ']'
			}
		}
		widerMap = append(widerMap, widerRow)
	}
	return widerMap, robotPosition
}

// wideBoxOtherSideDireciton returns the direction to the other side of a wide
// box, or nil if r is not part of a wide box.
func (s *Day15Solution) wideBoxOtherSideDirection(r rune) *util.Vector {
	switch r {
	case WideBoxLeftRune:
		return util.RightDirection
	case WideBoxRightRune:
		return util.LeftDirection
	default:
		return nil
	}
}

// findRobot finds the robot designated by RobotRune in storageMap, and returns
// the position (or nil if not found)
func findRobot(storageMap util.Matrix[rune]) *util.Vector {
	for i, row := range storageMap {
		for j, cell := range row {
			if cell == RobotRune {
				return util.NewVector(i, j)
			}
		}
	}
	return nil
}
