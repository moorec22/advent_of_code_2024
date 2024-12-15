package main

import (
	"advent/day01"
	"advent/day02"
	"advent/day03"
	"advent/day04"
	"advent/day05"
	"advent/day06"
	"advent/day07"
	"advent/day08"
	"advent/day09"
	"advent/day10"
	"advent/day11"
	"advent/util"
	"flag"
	"fmt"
)

const FilePrefix = "day%s/files/%s.txt"
const TestFileName = "test"
const InputFileName = "input"

func main() {
	testFlag, dayFlag := setUpFlags()
	if *dayFlag <= 0 {
		fmt.Println("Day number must be greater than 0")
		return
	}
	filepath := getFilepath(*dayFlag, *testFlag)

	factory, ok := SolutionFactories[*dayFlag]
	if !ok {
		fmt.Printf("No solution found for day %d\n", *dayFlag)
		return
	}

	solution, err := factory(filepath)
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

// setUpTestFlag sets up the test flag and the day number, and returns them.
func setUpFlags() (*bool, *int) {
	testFlag := flag.Bool("t", false, "run with test.txt")
	dayFlag := flag.Int("d", -1, "day number")
	flag.Parse()
	return testFlag, dayFlag
}

func getFilepath(day int, testFlag bool) string {
	filename := InputFileName
	if testFlag {
		filename = TestFileName
	}
	return fmt.Sprintf(FilePrefix, getTwoDigitNumber(day), filename)
}

func getTwoDigitNumber(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%d", n)
	}
	return fmt.Sprintf("%d", n)
}

type SolutionFactory func(string) (util.Solution, error)

var SolutionFactories = map[int]SolutionFactory{
	1:  Day01SolutionFactory,
	2:  Day02SolutionFactory,
	3:  Day03SolutionFactory,
	4:  Day04SolutionFactory,
	5:  Day05SolutionFactory,
	6:  Day06SolutionFactory,
	7:  Day07SolutionFactory,
	8:  Day08SolutionFactory,
	9:  Day09SolutionFactory,
	10: Day10SolutionFactory,
	11: Day11SolutionFactory,
}

func Day01SolutionFactory(filepath string) (util.Solution, error) {
	return day01.NewDay01Solution(filepath)
}

func Day02SolutionFactory(filepath string) (util.Solution, error) {
	return day02.NewDay02Solution(filepath)
}

func Day03SolutionFactory(filepath string) (util.Solution, error) {
	return day03.NewDay03Solution(filepath)
}

func Day04SolutionFactory(filepath string) (util.Solution, error) {
	return day04.NewDay04Solution(filepath)
}

func Day05SolutionFactory(filepath string) (util.Solution, error) {
	return day05.NewDay05Solution(filepath)
}

func Day06SolutionFactory(filepath string) (util.Solution, error) {
	return day06.NewDay06Solution(filepath)
}

func Day07SolutionFactory(filepath string) (util.Solution, error) {
	return day07.NewDay07Solution(filepath)
}

func Day08SolutionFactory(filepath string) (util.Solution, error) {
	return day08.NewDay08Solution(filepath)
}

func Day09SolutionFactory(filepath string) (util.Solution, error) {
	return day09.NewDay09Solution(filepath)
}

func Day10SolutionFactory(filepath string) (util.Solution, error) {
	return day10.NewDay10Solution(filepath)
}

func Day11SolutionFactory(filepath string) (util.Solution, error) {
	return day11.NewDay11Solution(filepath)
}
