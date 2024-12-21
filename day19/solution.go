// Advent of Code, 2024, Day 19
//
// https://adventofcode.com/2024/day/19
//
// Part 1: Deciding if a combination of towel stripe colors is possible or not
// can be accomplished with recursion. For each letter, decide if it is an end
// of a design or not. If it is, try using that design and recurse on the
// remaining string. Otherwise, continue to the next letter. This gives a runtime
// of O(k2^n) where k is the number of designs and n is the length of the string.
// This is the naive solution, but it works for part 1.
//
// Part 2: We cannot find all possible combinations of towels in a reasonable
// amount of time with our solution from part 1.
//
// Optimizations:
//   - use a trie to find all prefixes faster
//   - dynamic programming. Let's take a string s, and assume that for all
//     prefixes of s, we know how many ways they can be arranged. Then for
//     each suffix of s (t) in patterns, we can add the number of ways to
//     arrange s[0:len(s)-len(t)] to the number of ways to arrange t.
//
// In the end, the dynaming programming solution was enough to solve for part
// 2. If we use a trie that starts at the end of each suffix, and keep track
// of the node, we could speed up the search for each suffix.
package day19

import (
	"advent/util"
	"bufio"
	"strings"
)

type Day19Solution struct {
	patterns       map[string]bool
	desiredDesigns []string
}

func NewDay19Solution(filename string) (*Day19Solution, error) {
	patterns := make(map[string]bool)
	desiredDesigns := make([]string, 0)
	err := util.ProcessFile(filename, func(scanner *bufio.Scanner) error {
		// first line is the patterns
		scanner.Scan()
		patternsList := strings.Split(scanner.Text(), ", ")
		for _, pattern := range patternsList {
			patterns[pattern] = true
		}
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
	return s.totalNumArrangementsPossible(s.desiredDesigns, s.patterns), nil
}

// numDesignsPossible returns the number of designs in designs that can be arranged
// using patterns from patterns.
func (s *Day19Solution) numDesignsPossible(designs []string, patterns map[string]bool) int {
	count := 0
	for _, design := range designs {
		if s.designIsPossible(design, patterns) {
			count++
		}
	}
	return count
}

// totalNumArrangementsPossible returns the total number of arrangements possible for each design
// in designs, summed together.
func (s *Day19Solution) totalNumArrangementsPossible(designs []string, patterns map[string]bool) int {
	total := 0
	for _, design := range designs {
		total += s.numArrangementsPossible(design, patterns)
	}
	return total
}

// designIsPossible returns true if and only if the design can be arranged using
// the patterns.
func (s *Day19Solution) designIsPossible(design string, patterns map[string]bool) bool {
	return s.numArrangementsPossible(design, patterns) > 0
}

// numArrangementsPossible returns the number of ways to arrange the design using patterns.
func (s *Day19Solution) numArrangementsPossible(design string, patterns map[string]bool) int {
	counts := make([]int, len(design)+1)
	// this is the priming step; 0 is not a valid substring, but a way to start
	// the dynamic programming.
	counts[0] = 1
	for i := 1; i <= len(design); i++ {
		for pattern := range patterns {
			if strings.HasSuffix(design[:i], pattern) {
				counts[i] += counts[i-len(pattern)]
			}
		}
	}
	return counts[len(design)]
}
