// Advent of Code, 2024, Day 1
//
// https://adventofcode.com/2024/day/1
package day01

import (
	"advent/util"
	"bufio"
	"fmt"
	"sort"
)

const Filepath = "files/test.txt"

type Day01Solution struct {
	left  []int
	right []int
}

func NewDay01Solution(filepath string) (*Day01Solution, error) {
	left, right, err := getLists(filepath)
	if err != nil {
		return nil, err
	}
	sort.Ints(left)
	sort.Ints(right)
	return &Day01Solution{left, right}, nil
}

func (s *Day01Solution) PartOneAnswer() (int, error) {
	answer := 0
	for i := 0; i < len(s.left); i++ {
		answer += util.IntAbs(s.left[i] - s.right[i])
	}
	return answer, nil
}

func (s *Day01Solution) PartTwoAnswer() (int, error) {
	frequencies := s.getFrequencies(s.right)
	answer := 0
	for _, i := range s.left {
		frequency, ok := frequencies[i]
		if !ok {
			frequency = 0
		}
		answer += i * frequency
	}
	return answer, nil
}

// getFrequencies returns a map where keys are values in `a`, and values are
// the count of occurrences of that value in `a`.
func (s *Day01Solution) getFrequencies(a []int) map[int]int {
	frequencies := make(map[int]int)
	for _, i := range a {
		if _, ok := frequencies[i]; !ok {
			frequencies[i] = 0
		}
		frequencies[i]++
	}
	return frequencies
}

// getLists returns two lists of integers from a file designated by filepath.
// filepath is a path to a file where each line contains two numbers, separated
// by whitespace.
//
// It can be assumed that the returned lists are the same length.
func getLists(filepath string) ([]int, []int, error) {
	left := make([]int, 0)
	right := make([]int, 0)
	err := util.ProcessFile(filepath, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			var l, r int
			_, err := fmt.Sscanf(scanner.Text(), "%d %d", &l, &r)
			if err != nil {
				return err
			}
			left = append(left, l)
			right = append(right, r)
		}
		return nil
	})
	return left, right, err
}
