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
