// Advent of Code, 2024, Day 11
//
// https://adventofcode.com/2024/day/11
//
// Part 1: Initially, I implemented the problem using slices. I'd iterate
// over the stones, and apply the rules to each stone. This worked fine for
// part 1, but became a problem in part 2.
//
// Part 2: The slice approach led to unacceptable performance for part 2. I
// realized that the order of the stones didn't matter, and decided instead
// to use a map, from stone to number of stones of that number. I also used
// a cache, to store the results of applying rules to a stone, to avoid
// recalculating the same rule for the same stone multiple times. This yielded
// the answer for both parts much more quickly.
package day11

import (
	"advent/util"
	"bufio"
	"strconv"
	"strings"
)

const PartOneIteration = 25
const PartTwoIteration = 75

type StoneCache map[int][]int

type Day11Solution struct {
	// stone number mapped to the count of that stone
	initialStones map[int]int
}

func NewDay11Solution(filepath string) (*Day11Solution, error) {
	initialStones := make(map[int]int)
	err := util.ProcessFile(filepath, func(scan *bufio.Scanner) error {
		for scan.Scan() {
			line := strings.TrimSpace(scan.Text())
			lineParts := strings.Fields(line)
			for _, part := range lineParts {
				num, err := strconv.Atoi(part)
				if err != nil {
					return err
				}
				addStones(initialStones, num, 1)
			}
		}
		return nil
	})
	return &Day11Solution{initialStones}, err
}

func (s *Day11Solution) PartOneAnswer() (int, error) {
	stones := s.applyStandardRulesTimes(s.initialStones, PartOneIteration)
	return s.totalCounts(stones), nil
}

func (s *Day11Solution) PartTwoAnswer() (int, error) {
	stones := s.applyStandardRulesTimes(s.initialStones, PartTwoIteration)
	return s.totalCounts(stones), nil
}

// applyStandardRulesTimes applies the standard rules to the initial stones
// the given number of times, and returns the resulting stones. The map
// passed in is not modified.
func (s *Day11Solution) applyStandardRulesTimes(initialStones map[int]int, times int) map[int]int {
	stones := copySet(initialStones)
	rules := []Rule{
		NewZeroRule(),
		NewSplitRule(),
		NewMultRule(),
	}
	return s.applyRulesTimes(rules, stones, times, make(StoneCache))
}

// applyRulesTimes applies the given rules to the stones the given number of
// times, and returns the resulting stones.
func (s *Day11Solution) applyRulesTimes(rules []Rule, stones map[int]int, times int, cache StoneCache) map[int]int {
	for i := 0; i < times; i++ {
		newStones := make(map[int]int)
		for stone, count := range stones {
			for _, rule := range rules {
				if rule.IsApplicable(stone) {
					for _, newStone := range s.applyRuleWithCache(stone, cache, rule) {
						addStones(newStones, newStone, count)
					}
					break
				}
			}
		}
		stones = newStones
	}
	return stones
}

// applyRuleWithCache applies the given rule to the stone, using the cache to
// store the results of applying the rule to a stone. If the result is already
// in the cache, it is returned. Otherwise, the rule is applied and the result
// is stored in the cache.
func (s *Day11Solution) applyRuleWithCache(stone int, cache StoneCache, rule Rule) []int {
	if result, ok := cache[stone]; ok {
		return result
	}
	cache[stone] = rule.Apply(stone)
	return cache[stone]
}

// addStones adds the given stone to stones with the given count.
func addStones(stones map[int]int, stone, count int) {
	if _, ok := stones[stone]; !ok {
		stones[stone] = 0
	}
	stones[stone] += count
}

// totalCounts returns the sum of all values in the map
func (s *Day11Solution) totalCounts(dict map[int]int) int {
	total := 0
	for _, count := range dict {
		total += count
	}
	return total
}

// copySet returns a copy of the given map.
func copySet[K comparable, V any](m map[K]V) map[K]V {
	newMap := make(map[K]V)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}
