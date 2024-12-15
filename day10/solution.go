// Advent of Code: Day 10
//
// https://adventofcode.com/2024/day/10
//
// Part 1 Idea: For each trailhead, recursively follow the trail until a peak
// is reached. We never walk downhill, so we know we cannot reach loops.
//
// Part 2 Idea: given our part 1 solution, the only thing we need to change is
// to count each trailhead to peak pair non-uniquely.
package day10

import (
	"advent/util"
)

const Trailhead = '0'
const Peak = '9'

type Day10Solution struct {
	trailMap util.Matrix[rune]
}

func NewDay10Solution(filepath string) (*Day10Solution, error) {
	trailMap, err := util.ParseMatrix(filepath, func(r rune) rune {
		return r
	})
	return &Day10Solution{trailMap}, err
}

func (s *Day10Solution) PartOneAnswer() (int, error) {
	return s.countReachablePeaks(s.trailMap, true), nil
}

func (s *Day10Solution) PartTwoAnswer() (int, error) {
	return s.countReachablePeaks(s.trailMap, false), nil
}

// countReachablePeaks counts all peaks reachable from a trailhead. If unique
// is true, then each trailhead to peak pair is only counted once, no matter
// how many ways there are to reach it. Otherwise, all trails are counted.
func (s *Day10Solution) countReachablePeaks(trailMap util.Matrix[rune], unique bool) int {
	reachablePeaks := 0
	for i, row := range trailMap {
		for j, cell := range row {
			if cell == Trailhead {
				reachablePeaks += s.countReachablePeaksFrom(trailMap, util.NewVector(i, j), unique)
			}
		}
	}
	return reachablePeaks
}

// countReachablePeaksFrom counts all peaks reachable from a trailhead at p.
// If unique is true, then each trailhead to peak pair is only counted once,
// no matter how many ways there are to reach it. Otherwise, all trails are
// counted.
func (s *Day10Solution) countReachablePeaksFrom(trailMap util.Matrix[rune],
	p util.Vector, unique bool) int {
	reachablePeaks := s.getEndOfTrailsFrom(trailMap, p)
	if unique {
		return s.countUnique(reachablePeaks)
	} else {
		return len(reachablePeaks)
	}
}

// countTrailEndsFrom returns the ending position for every trail reachable from p.
func (s *Day10Solution) getEndOfTrailsFrom(trailMap util.Matrix[rune], p util.Vector) []util.Vector {
	elevation := trailMap.Get(p)
	if elevation == Peak {
		return []util.Vector{p}
	}
	reachablePeaks := make([]util.Vector, 0)
	for _, dir := range util.SimpleDirections {
		newPos := p.Add(dir)
		if trailMap.PosInBounds(newPos) && trailMap.Get(newPos) == elevation+1 {
			reachablePeaks = append(reachablePeaks, s.getEndOfTrailsFrom(trailMap, newPos)...)
		}
	}
	return reachablePeaks
}

func (s *Day10Solution) countUnique(positions []util.Vector) int {
	unique := make(map[util.Vector]bool)
	for _, p := range positions {
		unique[p] = true
	}
	return len(unique)
}
