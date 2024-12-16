// Advent of Code, 2024, Day 15
//
// https://adventofcode.com/2024/day/15
package day15

import (
	"advent/util"
	"bufio"
)

type Day15Solution struct {
	storageMap   util.Matrix[rune]
	instructions []rune
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
	return &Day15Solution{storageMap, instructions}, err
}

func (d *Day15Solution) PartOneAnswer() (int, error) {
	return 0, nil
}

func (d *Day15Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}
