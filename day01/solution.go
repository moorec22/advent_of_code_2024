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

func PartOneAnswer(filepath string) (int, error) {
	left, right, err := getLists(filepath)
	if err != nil {
		return 0, err
	}
	sort.Ints(left)
	sort.Ints(right)
	answer := 0
	for i := 0; i < len(left); i++ {
		answer += util.IntAbs(left[i] - right[i])
	}
	return answer, nil
}

func PartTwoAnswer(filepath string) (int, error) {
	left, right, err := getLists(filepath)
	if err != nil {
		return 0, err
	}
	frequencies := getFrequencies(right)
	answer := 0
	for _, i := range left {
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
func getFrequencies(a []int) map[int]int {
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
