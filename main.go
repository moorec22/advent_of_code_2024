package main

import (
	"advent/day01"
	"advent/util"
	"flag"
	"fmt"
)

const FilePrefix = "day01/files/"
const TestFile = "test.txt"
const InputFile = "input.txt"

func main() {
	testFlag := setUpTestFlag()
	filepath := getFilepath(*testFlag)

	var solution util.Solution

	solution, err := day01.NewDay01Solution(filepath)

	if err != nil {
		fmt.Printf("Error creating solution: %s\n", err)
		return
	}

	answer, err := solution.PartOneAnswer()
	if err != nil {
		fmt.Printf("Error getting answer for part 1: %s\n", err)
		return
	}
	fmt.Printf("Part 1 answer: %d\n", answer)

	answer, err = solution.PartTwoAnswer()
	if err != nil {
		fmt.Printf("Error getting answer for part 2: %s\n", err)
		return
	}
	fmt.Printf("Part 2 answer: %d\n", answer)
}

func setUpTestFlag() *bool {
	testFlag := flag.Bool("t", false, "run with test.txt")
	flag.Parse()
	return testFlag
}

func getFilepath(testFlag bool) string {
	if testFlag {
		return FilePrefix + TestFile
	}
	return FilePrefix + InputFile
}
