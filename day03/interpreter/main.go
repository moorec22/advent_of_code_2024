package interpreter

import (
	"advent/util"
	"bufio"
)

func RunProgram(s *bufio.Scanner, matchers []Matcher) (int, error) {
	programState := NewProgramState()
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		var instruction Instruction = NewStartInstruction()
		remainder := line
		var err error
		for !util.IsNil(instruction) {
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
	if util.IsNil(nextMatcher) {
		return nil, s, nil
	} else {
		instruction, err := nextMatcher.Parse(s[nextMatch[0]:nextMatch[1]])
		if err != nil {
			return nil, s, err
		}
		return instruction, s[nextMatch[1]:], nil
	}
}
