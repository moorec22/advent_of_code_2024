// Advent of Code, 2024, Day 22
//
// https://adventofcode.com/2024/day/22
package day22

import (
	"advent/util"
	"bufio"
	"strconv"
)

type Day22Solution struct {
	initialSecrets []int
}

func NewDay22Solution(filename string) (*Day22Solution, error) {
	initialSecrets := make([]int, 0)
	err := util.ProcessFile(filename, func(scanner *bufio.Scanner) error {
		for scanner.scan() {
			if scanner.Text() == "" {
				continue
			}
			secret, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return err
			}
			initialSecrets = append(initialSecrets, secret)
		}
		return scanner.Err()
	})
	return &Day22Solution{initialSecrets}, err
}

func (s *Day22Solution) PartOneAnswer() (int, error) {
	return 0, nil
}

func (s *Day22Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}
