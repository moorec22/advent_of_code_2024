// Advent of Code, 2024, Day 16
//
// https://adventofcode.com/2024/day/16
package day16

import (
	"advent/util"
)

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
	return 0, nil
}

func (s *Day16Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}

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
