package day03

import (
	"advent/util"
	"bufio"
	"reflect"
)

func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

func PartOneAnswer(filepath string) (int, error) {
	answer := 0
	err := util.ProcessFile(filepath, func(s *bufio.Scanner) error {
		var err error
		matchers := []Matcher{
			NewMultiplyMatcher(),
		}
		answer, err = runProgram(s, matchers)
		if err != nil {
			return err
		}
		return nil
	})
	return answer, err
}

func PartTwoAnswer(filepath string) (int, error) {
	answer := 0
	err := util.ProcessFile(filepath, func(s *bufio.Scanner) error {
		var err error
		matchers := []Matcher{
			NewMultiplyMatcher(),
			NewDoMatcher(),
			NewDontMatcher(),
		}
		answer, err = runProgram(s, matchers)
		if err != nil {
			return err
		}
		return nil
	})
	return answer, err
}

func runProgram(s *bufio.Scanner, matchers []Matcher) (int, error) {
	programState := NewProgramState()
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		var instruction Instruction = NewStartInstruction()
		remainder := line
		var err error
		for !isNil(instruction) {
			programState = instruction.Execute(programState)
			instruction, remainder, err = getNextInstruction(remainder, matchers)
			if err != nil {
				return 0, err
			}
		}
	}
	return programState.Answer, nil
}

func getNextInstruction(s string, matchers []Matcher) (Instruction, string, error) {
	var nextMatcher Matcher = nil
	var nextMatch []int = nil
	for _, matcher := range matchers {
		match := matcher.NextMatch(s)
		if match != nil && (nextMatch == nil || match[0] < nextMatch[0]) {
			nextMatch = match
			nextMatcher = matcher
		}
	}
	if isNil(nextMatcher) {
		return nil, s, nil
	} else {
		instruction, err := nextMatcher.Parse(s[nextMatch[0]:nextMatch[1]])
		if err != nil {
			return nil, s, err
		}
		return instruction, s[nextMatch[1]:], nil
	}
}
