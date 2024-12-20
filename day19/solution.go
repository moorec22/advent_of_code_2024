// Advent of Code, 2024, Day 19
//
// https://adventofcode.com/2024/day/19
//
// Part 1: Deciding if a combination of towel stripe colors is possible or not
// can be accomplished with recursion. For each letter, decide if it is an end
// of a design or not. If it is, try using that design and recurse on the
// remaining string. Otherwise, continue to the next letter. This gives a runtime
// of O(k2^n) where k is the number of designs and n is the length of the string.
package day19

import (
	"advent/util"
	"bufio"
	"strings"
)

type Day19Solution struct {
	patterns       []string
	desiredDesigns []string
}

func NewDay19Solution(filename string) (*Day19Solution, error) {
	var patterns []string
	desiredDesigns := make([]string, 0)
	err := util.ProcessFile(filename, func(scanner *bufio.Scanner) error {
		// first line is the patterns
		scanner.Scan()
		patterns = strings.Split(scanner.Text(), ", ")
		// then comes a blank line
		scanner.Scan()
		// and then the desired designs
		for scanner.Scan() {
			desiredDesigns = append(desiredDesigns, scanner.Text())
		}
		return scanner.Err()
	})
	return &Day19Solution{patterns, desiredDesigns}, err
}

func (s *Day19Solution) PartOneAnswer() (int, error) {
	return s.numDesignsPossible(s.desiredDesigns, s.patterns), nil
}

func (s *Day19Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}

func (s *Day19Solution) numDesignsPossible(designs []string, patterns []string) int {
	count := 0
	for _, design := range designs {
		if s.designIsPossible(design, patterns) {
			count++
		}
	}
	return count
}

func (s *Day19Solution) designIsPossible(design string, patterns []string) bool {
	if len(design) == 0 {
		return true
	}
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) && s.designIsPossible(design[len(pattern):], patterns) {
			return true
		}
	}
	return false
}
