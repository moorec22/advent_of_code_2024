// Advent of Code, 2024, Day 2
//
// https://adventofcode.com/2024/day/2
package main

import (
	"advent/processing"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const Filepath = "files/input.txt"

func main() {
	parOne, err := safeCount(false)
	if err != nil {
		fmt.Printf("Error counting safe numbers: %s\n", err)
		return
	}
	fmt.Printf("Safe without dampener: %d\n", parOne)

	partTwo, err := safeCount(true)
	if err != nil {
		fmt.Printf("Error counting safe numbers with dampener: %s\n", err)
		return
	}
	fmt.Printf("Safe with dampener: %d\n", partTwo)
}

func safeCount(problemDampener bool) (int, error) {
	safeCount := 0
	err := processing.ProcessFile(Filepath, func(s *bufio.Scanner) error {
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

func isSafe(parts []int, problemDampener bool) bool {
	return isSafeAscending(parts, problemDampener) || isSafeDescending(parts, problemDampener)
}

func isSafeAscending(parts []int, problemDampener bool) bool {
	return isSafeDirection(parts, true, problemDampener)
}

func isSafeDescending(parts []int, problemDampener bool) bool {
	return isSafeDirection(parts, false, problemDampener)
}

func isSafeDirection(parts []int, wantAscending, problemDampener bool) bool {
	for i := 0; i < len(parts)-1; i++ {
		first := parts[i]
		second := parts[i+1]
		safe, ascending := checkNumbers(first, second)
		if !safe || wantAscending != ascending {
			if problemDampener {
				check := isSafeDirection(copyWithoutIndex(parts, i), wantAscending, false)
				if !check {
					check = isSafeDirection(copyWithoutIndex(parts, i+1), wantAscending, false)
				}
				// fmt.Printf("HERE: %v %t\n", parts, check)
				return check
			} else {
				return false
			}
		}
	}
	return true
}

// checkNumbers checks if first and second are safe. Returns two booleans:
// the first is true if the numbers are safe, the second is true if the numbers
// are ascending.
func checkNumbers(first, second int) (bool, bool) {
	diff := intAbs(second - first)
	return diff >= 1 && diff <= 3, second > first
}

func intAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// returns a copy of s without the element at index
func copyWithoutIndex(s []int, index int) []int {
	newSlice := make([]int, len(s)-1)
	copy(newSlice, s[:index])
	copy(newSlice[index:], s[index+1:])
	return newSlice
}
