// Advent of Code, 2024, Day 2
//
// https://adventofcode.com/2024/day/2
package day02

import (
	"advent/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func PartOneAnswer(filepath string) (int, error) {
	return safeCount(filepath, false)
}

func PartTwoAnswer(filepath string) (int, error) {
	return safeCount(filepath, true)
}

// safeCount returns the number of safe number lists in the file at filepath. A
// number list is safe if the difference between each number is between 1 and 3,
// inclusive, and if all numbers are either descending or ascending. If
// problemDampener is true, it will return true if removing one number from the
// list makes the remaining numbers safe.
func safeCount(filepath string, problemDampener bool) (int, error) {
	safeCount := 0
	err := util.ProcessFile(filepath, func(s *bufio.Scanner) error {
		for s.Scan() {
			line := s.Text()
			if line == "" {
				continue
			}
			numberList, err := getNumberList(line)
			if err != nil {
				fmt.Printf("Error getting number list for line %s: %s\n", line, err)
				continue
			}
			if isSafe(numberList, problemDampener) {
				safeCount++
			}
		}
		return nil
	})
	return safeCount, err
}

// isSafe returns whether the numbers in parts are safe, either in ascending or
// descending order. If problemDampener is true, it will return true if removing
// one number from parts makes the remaining numbers safe.
func isSafe(parts []int, problemDampener bool) bool {
	return isSafeAscending(parts, problemDampener) || isSafeDescending(parts, problemDampener)
}

// isSafeAscending returns whether the numbers in parts are safe in ascending
// order. If problemDampener is true, it will return true if removing one number
// from parts makes the remaining numbers safe.
func isSafeAscending(parts []int, problemDampener bool) bool {
	return isSafeDirection(parts, true, problemDampener)
}

// isSafeDescending returns whether the numbers in parts are safe in descending
// order. If problemDampener is true, it will return true if removing one number
// from parts makes the remaining numbers safe.
func isSafeDescending(parts []int, problemDampener bool) bool {
	return isSafeDirection(parts, false, problemDampener)
}

// isSafeDirection returns whether the numbers in parts are safe, ascending if
// wantAscending, otherwise descending. If problemDampener is true, it will
// return true if removing one number from parts makes the remaining numbers
// safe.
func isSafeDirection(parts []int, wantAscending, problemDampener bool) bool {
	for i := 0; i < len(parts)-1; i++ {
		first := parts[i]
		second := parts[i+1]
		safe, ascending := numbersAreSafe(first, second)
		if !safe || wantAscending != ascending {
			if problemDampener {
				return isSafeDirection(copyWithoutIndex(parts, i), wantAscending, false) ||
					isSafeDirection(copyWithoutIndex(parts, i+1), wantAscending, false)
			} else {
				return false
			}
		}
	}
	return true
}

// numbersAreSafe returns whether the two numbers are safe and whether they are
// ascending. The two numbers are safe if their difference is between 1 and 3,
// inclusive.
func numbersAreSafe(first, second int) (bool, bool) {
	diff := util.IntAbs(second - first)
	return diff >= 1 && diff <= 3, second > first
}

// getNumberList takes a string of space-separated numbers and returns a slice
// of those numbers.
func getNumberList(line string) ([]int, error) {
	parts := strings.Split(line, " ")
	numbers := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		numbers[i] = num
	}
	return numbers, nil
}

// copyWithoutIndex returns a copy of s without the element at index.
func copyWithoutIndex(s []int, index int) []int {
	newSlice := make([]int, len(s)-1)
	copy(newSlice, s[:index])
	copy(newSlice[index:], s[index+1:])
	return newSlice
}
