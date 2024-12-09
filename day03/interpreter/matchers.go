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
	NextMatch(string) []int
	Parse(string) (Instruction, error)
}

type MultiplyMatcher struct {
	regex *regexp.Regexp
}

func NewMultiplyMatcher() *MultiplyMatcher {
	return &MultiplyMatcher{
		regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`),
	}
}

func (m *MultiplyMatcher) NextMatch(s string) []int {
	return m.regex.FindStringIndex(s)
}

func (m *MultiplyMatcher) Parse(s string) (Instruction, error) {
	if !strings.HasPrefix(s, "mul(") || !strings.HasSuffix(s, ")") {
		return nil, fmt.Errorf("invalid instruction: %s", s)
	}
	numbers := strings.Split(s[4:len(s)-1], ",")
	if len(numbers) != 2 {
		return nil, fmt.Errorf("invalid instruction: %s", s)
	}
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
	regex *regexp.Regexp
}

func NewDoMatcher() *DoMatcher {
	return &DoMatcher{
		regexp.MustCompile(`do\(\)`),
	}
}

func (m *DoMatcher) NextMatch(s string) []int {
	return m.regex.FindStringIndex(s)
}

func (m *DoMatcher) Parse(s string) (Instruction, error) {
	if s != "do()" {
		return nil, fmt.Errorf("invalid instruction: %s", s)
	}
	return NewDoInstruction(), nil
}

type DontMatcher struct {
	regex *regexp.Regexp
}

func NewDontMatcher() *DontMatcher {
	return &DontMatcher{
		regexp.MustCompile(`don't\(\)`),
	}
}

func (m *DontMatcher) NextMatch(s string) []int {
	return m.regex.FindStringIndex(s)
}

func (m *DontMatcher) Parse(s string) (Instruction, error) {
	if s != "don't()" {
		return nil, fmt.Errorf("invalid instruction: %s", s)
	}
	return NewDontInstruction(), nil
}
