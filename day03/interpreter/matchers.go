package interpreter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Matcher is an interface defining matchers, which can be used to see the next match
// for a string, as well as parse instructions from a string to an instruction
type Matcher interface {
	// NextMatch returns the next match for the given string, or nil if there is no match
	NextMatch(string) []int
	// Parse parses the given string into an instruction, or returns an error if the string is invalid.
	// A valid string must be an exact match. In other words, isMatch must return true.
	Parse(string) (Instruction, error)
	// isMatch returns true iff the given string is a match, with no extra parts
	isMatch(string) bool
}

type BaseMatcher struct {
	regex *regexp.Regexp
}

func (m *BaseMatcher) NextMatch(s string) []int {
	return m.regex.FindStringIndex(s)
}

func (m *BaseMatcher) isMatch(s string) bool {
	nextMatch := m.NextMatch(s)
	return nextMatch != nil && nextMatch[0] == 0 && nextMatch[1] == len(s)
}

type MultiplyMatcher struct {
	BaseMatcher
}

func NewMultiplyMatcher() *MultiplyMatcher {
	return &MultiplyMatcher{BaseMatcher{regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)}}
}

func (m *MultiplyMatcher) Parse(s string) (Instruction, error) {
	if !m.isMatch(s) {
		return nil, fmt.Errorf("invalid instruction: %s", s)
	}
	numbers := strings.Split(s[4:len(s)-1], ",")
	left, err := strconv.Atoi(numbers[0])
	if err != nil {
		return nil, fmt.Errorf("invalid instruction: %s", s)
	}
	right, err := strconv.Atoi(numbers[1])
	if err != nil {
		return nil, fmt.Errorf("invalid instruction: %s", s)
	}
	return NewMultiplyInstruction(left, right), nil
}

type DoMatcher struct {
	BaseMatcher
}

func NewDoMatcher() *DoMatcher {
	return &DoMatcher{BaseMatcher{regexp.MustCompile(`do\(\)`)}}
}

func (m *DoMatcher) Parse(s string) (Instruction, error) {
	if !m.isMatch(s) {
		return nil, fmt.Errorf("invalid instruction: %s", s)
	}
	return NewDoInstruction(), nil
}

type DontMatcher struct {
	BaseMatcher
}

func NewDontMatcher() *DontMatcher {
	return &DontMatcher{BaseMatcher{regexp.MustCompile(`don't\(\)`)}}
}

func (m *DontMatcher) Parse(s string) (Instruction, error) {
	if !m.isMatch(s) {
		return nil, fmt.Errorf("invalid instruction: %s", s)
	}
	return NewDontInstruction(), nil
}
