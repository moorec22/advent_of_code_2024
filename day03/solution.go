// Advent of Code, 2024, Day 3
//
// https://adventofcode.com/2024/day/3
//
// For this problem I chose to implement a simple interpreter that can handle
// the instructions given in the input file. The interpreter will scan the
// input for the following instructions (dismissing any junk):
//
// - Multiply: multiplies the current value by the given number
// - Do: sets the interpreter to execute Multiply instructions
// - Don't: sets the interpreter to ignore Multiply instructions
package day03

import (
	"advent/day03/interpreter"
	"advent/util"
	"bufio"
)

func PartOneAnswer(filepath string) (int, error) {
	matchers := []interpreter.Matcher{
		interpreter.NewMultiplyMatcher(),
	}
	return runProgramWithMatchers(filepath, matchers)
}

func PartTwoAnswer(filepath string) (int, error) {
	matchers := []interpreter.Matcher{
		interpreter.NewMultiplyMatcher(),
		interpreter.NewDoMatcher(),
		interpreter.NewDontMatcher(),
	}
	return runProgramWithMatchers(filepath, matchers)
}

// runProgramWithMatchers runs the interpreter with the given matchers and
// returns the final value of the program.
func runProgramWithMatchers(filepath string, matchers []interpreter.Matcher) (int, error) {
	answer := 0
	err := util.ProcessFile(filepath, func(s *bufio.Scanner) error {
		var err error
		answer, err = interpreter.RunProgram(s, matchers)
		if err != nil {
			return err
		}
		return nil
	})
	return answer, err
}
