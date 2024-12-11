// Advent of Code, 2024, Day 6
//
// https://adventofcode.com/2024/day/6
package day06

import (
	"advent/util"
	"bufio"
	"fmt"
)

func PartOneAnswer(filepath string) (int, error) {
	labMap, guard, err := getLabMapAndGuard(filepath)
	fmt.Println(labMap, guard)
	return 0, err
}

func PartTwoAnswer(filepath string) (int, error) {
	return 0, nil
}

func getLabMapAndGuard(filename string) (util.Matrix[rune], util.Position, error) {
	labMap := make(util.Matrix[rune], 0)
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

func isGuard(r rune) bool {
	return r == '^' || r == 'v' || r == '<' || r == '>'
}
