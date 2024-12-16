// Advent of Code, 2024, Day 15
//
// https://adventofcode.com/2024/day/15
//
// Part 1: I decided to start by running the simulation, step by step.
package day15

import (
	"advent/util"
	"bufio"
	"fmt"
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

func (d *Day15Solution) PartOneAnswer() (int, error) {
	storageMap := d.storageMap.Copy()
	err := d.makeMoves(storageMap, d.robotPosition, d.instructions)
	if err != nil {
		return 0, err
	}
	gpsCoordinates := d.getGpsCoordinates(storageMap)
	return util.SliceSum(gpsCoordinates), nil
}

func (d *Day15Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}

// makeMoves will make the moves specified by the moves slice, in order. It
// returns an error if there is no robot at robotPos.
func (d *Day15Solution) makeMoves(storageMap util.Matrix[rune], robotPos *util.Vector, moves []rune) error {
	var err error
	for _, move := range moves {
		robotPos, err = d.makeMove(storageMap, robotPos, move)
		if err != nil {
			return err
		}
	}
	return nil
}

// makeMove will move the robot at robotPos, returning an error if there is no
// robot at robotPos. The move made is a move on storageMap specified by the
// problem statement. storageMap is modified by the next move.
func (d *Day15Solution) makeMove(storageMap util.Matrix[rune], robotPos *util.Vector, move rune) (*util.Vector, error) {
	if storageMap.Get(robotPos) != RobotRune {
		return robotPos, fmt.Errorf("no robot at position %v", robotPos)
	}
	dir := RuneDirections[move]
	movePos := robotPos.Add(dir)
	el := storageMap.Get(movePos)
	moveFunc, err := d.getMoveFunction(el)
	if err != nil {
		return robotPos, err
	}
	if moveFunc(storageMap, robotPos, dir) {
		robotPos = movePos
	}
	return robotPos, nil
}

// getMoveFunction returns the appropriate function for the rune r
func (d *Day15Solution) getMoveFunction(r rune) (func(util.Matrix[rune], *util.Vector, *util.Vector) bool, error) {
	switch r {
	case EmptyRune:
		return d.makeEmptyMove, nil
	case WallRune:
		return d.makeWallMove, nil
	case BoxRune:
		return d.makeBoxMove, nil
	default:
		return nil, fmt.Errorf("no function for rune: %v", r)
	}
}

// makeBoxMove moves the object and the box as specified by the problem. It
// looks for an empty space at some point in the direction given. If it finds
// one, it moves the object and all boxes between the object and the empty space
// to that position. Otherwise, it does nothing. It returns true if a move was
// made. It assumes obPos is a moveable object. Otherwise, behavior is
// unspecified.
func (d *Day15Solution) makeBoxMove(storageMap util.Matrix[rune], obPos, dir *util.Vector) bool {
	nextPos := obPos.Add(dir)
	nextOb := storageMap.Get(nextPos)
	switch nextOb {
	case EmptyRune:
		return d.makeEmptyMove(storageMap, obPos, dir)
	case WallRune:
		return false
	case BoxRune:
		if d.makeBoxMove(storageMap, nextPos, dir) {
			d.makeEmptyMove(storageMap, obPos, dir)
			return true
		}
	}
	return false
}

// makeEmptyMove moves obPos to nextPos, leaving obPos empty. It assumes obPos
// is moveable and nextPos is empty, otherwise behavior is unspecified.
func (d *Day15Solution) makeEmptyMove(storageMap util.Matrix[rune], obPos, dir *util.Vector) bool {
	nextPos := obPos.Add(dir)
	ob := storageMap.Get(obPos)
	storageMap.Set(obPos, EmptyRune)
	storageMap.Set(nextPos, ob)
	return true
}

// makeWallMove does nothing, as happens when the robot is facing a wall. It
// returns false.
func (d *Day15Solution) makeWallMove(storageMap util.Matrix[rune], obPos, dir *util.Vector) bool {
	return false
}

// getGpsCoordiantes returns a list of all GPS coordinates for boxes found in
// storageMap, as specified by the problem statement.
func (d *Day15Solution) getGpsCoordinates(storageMap util.Matrix[rune]) []int {
	gpsCoordinates := make([]int, 0)
	for i := range storageMap {
		for j := range storageMap[i] {
			if storageMap[i][j] == BoxRune {
				gpsCoordinates = append(gpsCoordinates, 100*i+j)
			}
		}
	}
	return gpsCoordinates
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
