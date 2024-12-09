package main

import (
	"advent/day02"
	"fmt"
)

const Filepath = "day02/files/test.txt"

func main() {
	answer, err := day02.PartOneAnswer(Filepath)
	if err != nil {
		fmt.Printf("Error getting answer for part 1: %s\n", err)
		return
	}
	fmt.Printf("Part 1 answer: %d\n", answer)

	answer, err = day02.PartTwoAnswer(Filepath)
	if err != nil {
		fmt.Printf("Error getting answer for part 2: %s\n", err)
		return
	}
	fmt.Printf("Part 2 answer: %d\n", answer)
}
