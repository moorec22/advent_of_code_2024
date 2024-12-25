// Advent of Code, 2024, Day 21
//
// https://adventofcode.com/2024/day/21
package day21

import (
	"advent/util"
	"bufio"
)

type Day21Solution struct {
	codes []string
}

func NewDay21Solution(filename string) (*Day21Solution, error) {
	codes := make([]string, 0)
	err := util.ProcessFile(filename, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			codes = append(codes, scanner.Text())
		}
		return scanner.Err()
	})
	return &Day21Solution{codes}, err
}

func (s *Day21Solution) PartOneAnswer() (int, error) {
	return 0, nil
}

func (s *Day21Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}
