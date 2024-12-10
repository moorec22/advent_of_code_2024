package main

import (
	"advent/day04"
	"fmt"
)

const Filepath = "day04/files/input.txt"

type Solution interface {
	PartOneAnswer(filepath string) (int, error)
	PartTwoAnswer(filepath string) (int, error)
}

func main() {
	answer, err := day04.PartOneAnswer(Filepath)
	if err != nil {
		fmt.Printf("Error getting answer for part 1: %s\n", err)
		return
	}
	fmt.Printf("Part 1 answer: %d\n", answer)

	answer, err = day04.PartTwoAnswer(Filepath)
	if err != nil {
		fmt.Printf("Error getting answer for part 2: %s\n", err)
		return
	}
	fmt.Printf("Part 2 answer: %d\n", answer)
}
