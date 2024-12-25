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
	"advent/day12"
	"advent/day13"
	"advent/day14"
	"advent/day15"
	"advent/day16"
	"advent/day17"
	"advent/day18"
	"advent/day19"
	"advent/day20"
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
	12: Day12SolutionFactory,
	13: Day13SolutionFactory,
	14: Day14SolutionFactory,
	15: Day15SolutionFactory,
	16: Day16SolutionFactory,
	17: Day17SolutionFactory,
	18: Day18SolutionFactory,
	19: Day19SolutionFactory,
	20: Day20SolutionFactory,
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

func Day12SolutionFactory(filepath string) (util.Solution, error) {
	return day12.NewDay12Solution(filepath)
}

func Day13SolutionFactory(filepath string) (util.Solution, error) {
	return day13.NewDay13Solution(filepath)
}

func Day14SolutionFactory(filepath string) (util.Solution, error) {
	return day14.NewDay14Solution(filepath)
}

func Day15SolutionFactory(filepath string) (util.Solution, error) {
	return day15.NewDay15Solution(filepath)
}

func Day16SolutionFactory(filepath string) (util.Solution, error) {
	return day16.NewDay16Solution(filepath)
}

func Day17SolutionFactory(filepath string) (util.Solution, error) {
	return day17.NewDay17Solution(filepath)
}

func Day18SolutionFactory(filepath string) (util.Solution, error) {
	return day18.NewDay18Solution(filepath)
}

func Day19SolutionFactory(filepath string) (util.Solution, error) {
	return day19.NewDay19Solution(filepath)
}

func Day20SolutionFactory(filepath string) (util.Solution, error) {
	return day20.NewDay20Solution(filepath)
}
