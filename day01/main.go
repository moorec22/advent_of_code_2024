// Advent of Code, 2024, Day 1
//
// https://adventofcode.com/2024/day/1
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const Filepath = "files/test.txt"

func main() {
	answer, err := partOneAnswer(Filepath)
	if err != nil {
		fmt.Printf("Error getting answer for part 1: %s\n", err)
		return
	}
	fmt.Printf("Part 1 answer: %d\n", answer)
}

func partOneAnswer(filepath string) (int, error) {
	left, right, err := getLists(filepath)
	if err != nil {
		return 0, err
	}
	if len(left) != len(right) {
		return 0, fmt.Errorf("left and right lists are not the same length: %v, %v", left, right)
	}
	answer := 0
	for i := 0; i < len(left); i++ {
		answer += intAbs(left[i] - right[i])
	}
	return answer, nil
}

// intAbs returns the absolute value of x.
func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// getLists returns two lists of integers from a file designated by filepath.
// filepath is a path to a file where each line contains two numbers, separated
// by whitespace.
//
// The returned lists are sorted from smallest to largest.
func getLists(filepath string) ([]int, []int, error) {
	left := make([]int, 0)
	right := make([]int, 0)
	err := processFile(filepath, func(scanner *bufio.Scanner) error {
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
	sort.Ints(left)
	sort.Ints(right)
	return left, right, err
}

// processFile safely opens a file designated by filepath, passes a scanner of
// that file to process, and closes that file. Returns an error if there is an
// error opening the file, or if process returns an error.
func processFile(filepath string, process func(*bufio.Scanner) error) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	return process(scanner)
}
