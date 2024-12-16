// Advent of Code, 2024, Day 14
//
// https://adventofcode.com/2024/day/14
//
// Part 1: The robots movements are linear, so we can calculate the
// position they will be at after any number of steps in constant time.
// This will give a straightforward solution to part 1 in O(n) time.
//
// Part 2: To find the christmas tree, I made an educated guess that there
// would be a line of robots representing the trunk. It turns out there's
// a frame too! I just increment the robots and check for a long line,
// returning the first time I find one.
package day14

import (
	"advent/util"
	"bufio"
	"fmt"
	"strings"
)

const PartOneSteps = 100
const HallwayWidth = 101
const HallwayHeight = 103

type RobotInfo struct {
	pos *util.Vector
	vel *util.Vector
}

type Day14Solution struct {
	robotInfos []*RobotInfo
}

func NewDay14Solution(filepath string) (*Day14Solution, error) {
	robotInfos := make([]*RobotInfo, 0)
	err := util.ProcessFile(filepath, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			line := scanner.Text()
			robotInfo, err := getRobotStartingInfo(line)
			if err != nil {
				return err
			}
			robotInfos = append(robotInfos, robotInfo)
		}
		return nil
	})
	return &Day14Solution{robotInfos}, err
}

func (s *Day14Solution) PartOneAnswer() (int, error) {
	robotInfos := s.stateAfterXSteps(s.robotInfos, PartOneSteps)
	return s.getSafetyFactor(robotInfos), nil
}

func (s *Day14Solution) PartTwoAnswer() (int, error) {
	return s.christmasTreeSteps(s.robotInfos), nil
}

// stateAfterXSteps returns the state of the robots after steps steps.
func (s *Day14Solution) stateAfterXSteps(robotInfos []*RobotInfo, steps int) []*RobotInfo {
	bounds := util.NewVector(HallwayWidth, HallwayHeight)
	newRobotInfos := make([]*RobotInfo, len(robotInfos))
	for i, robotInfo := range robotInfos {
		newRobotInfo := &RobotInfo{
			robotInfo.pos.Add(robotInfo.vel.ScalarMultiply(steps)).MathModulo(bounds),
			robotInfo.vel,
		}
		newRobotInfos[i] = newRobotInfo
	}
	return newRobotInfos
}

// christmasTreeSteps returns the number of steps it takes for the robots to
// form a christmas tree.
func (s *Day14Solution) christmasTreeSteps(robotInfos []*RobotInfo) int {
	steps := 0
	for !s.isChristmasTree(robotInfos) {
		newRobotInfos := s.stateAfterXSteps(robotInfos, 1)
		steps++
		robotInfos = newRobotInfos
	}
	printState(robotInfos)
	return steps
}

// isChristmasTree returns whether the robots are in the shape of a christmas
// tree.
func (s *Day14Solution) isChristmasTree(robotInfos []*RobotInfo) bool {
	return s.hasLineOfSize(robotInfos, 30)
}

// getSafetyFactor returns the safety factor of the robots, calculated by
// counting how many robots are in each quadrant and multiplying those four
// numbers together.
func (s *Day14Solution) getSafetyFactor(robotInfos []*RobotInfo) int {
	quadrantCounts := make([]int, 4)
	for _, robotInfo := range robotInfos {
		quadrant := s.getQuadrant(robotInfo.pos)
		if quadrant > -1 {
			quadrantCounts[quadrant]++
		}
	}
	return s.sliceProduct(quadrantCounts)
}

// getQuadrant returns the quadrant of a given position. If the position is on
// the exact middle line vertically or horizontally, it is considered to be in
// no quadrant and -1 is returned.
func (s *Day14Solution) getQuadrant(pos *util.Vector) int {
	equator := (HallwayWidth - 1) / 2
	meridian := (HallwayHeight - 1) / 2
	if pos.X < equator && pos.Y < meridian {
		return 0
	} else if pos.X > equator && pos.Y < meridian {
		return 1
	} else if pos.X < equator && pos.Y > meridian {
		return 2
	} else if pos.X > equator && pos.Y > meridian {
		return 3
	}
	return -1
}

// sliceProduct returns the product of all the numbers in a slice.
func (s *Day14Solution) sliceProduct(slice []int) int {
	product := 1
	for _, num := range slice {
		product *= num
	}
	return product
}

// getRobotStartingInfo returns a RobotInfo struct from a line of input.
func getRobotStartingInfo(line string) (*RobotInfo, error) {
	parts := strings.Split(line, " ")
	pos, err := util.ParseVector(parts[0][2:])
	if err != nil {
		return &RobotInfo{}, err
	}
	vel, err := util.ParseVector(parts[1][2:])
	if err != nil {
		return &RobotInfo{}, err
	}
	return &RobotInfo{pos, vel}, nil
}

// printState prints the a map with the robots' positions.
func printState(robotInfos []*RobotInfo) {
	positions := getPositions(robotInfos)
	for y := 0; y < HallwayHeight; y++ {
		for x := 0; x < HallwayWidth; x++ {
			if count, ok := positions[*util.NewVector(x, y)]; ok {
				fmt.Print(count)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// hasLineOfSize returns whether the robots form a line of size size.
func (s *Day14Solution) hasLineOfSize(robotInfos []*RobotInfo, size int) bool {
	positions := getPositions(robotInfos)
	for y := 0; y < HallwayHeight; y++ {
		lineSize := 0
		for x := 0; x < HallwayWidth; x++ {
			if _, ok := positions[*util.NewVector(x, y)]; ok {
				lineSize++
				if lineSize >= size {
					return true
				}
			} else {
				lineSize = 0
			}
		}
	}
	return false
}

// getPositions returns a map of positions to the number of robots at that
// position.
func getPositions(robotInfos []*RobotInfo) map[util.Vector]int {
	positions := make(map[util.Vector]int)
	for _, robotInfo := range robotInfos {
		if _, ok := positions[*robotInfo.pos]; !ok {
			positions[*robotInfo.pos] = 0
		}
		positions[*robotInfo.pos]++
	}
	return positions
}
