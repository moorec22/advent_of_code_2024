// Advent of Code, 2024, Day 8
//
// https://adventofcode.com/2024/day/8
package day08

import (
	"advent/util"
)

type Day08Solution struct {
	antennaMap util.Matrix[rune]
}

func NewDay08Solution(filepath string) (*Day08Solution, error) {
	antennaMap, err := util.ParseMatrix(filepath)
	antennaMap.Print(func(r rune) string {
		return string(r)
	})
	return &Day08Solution{antennaMap}, err
}

func (s *Day08Solution) PartOneAnswer() (int, error) {
	return 0, nil
}

func (s *Day08Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}
