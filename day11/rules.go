package day11

import "strconv"

// A rule to apply a change to a number.
type Rule interface {
	// Returns true if the rule is applicable to the number.
	IsApplicable(int) bool
	// Applies the rule and returns resulting numbers. If the rule is not
	// applicable, behavior is not defined.
	Apply(int) []int
}

// Zero rule only applies to 0, and returns 1
type ZeroRule struct{}

func NewZeroRule() *ZeroRule {
	return &ZeroRule{}
}

func (r *ZeroRule) IsApplicable(n int) bool {
	return n == 0
}

func (r *ZeroRule) Apply(int) []int {
	return []int{1}
}

type SplitRule struct{}

func NewSplitRule() *SplitRule {
	return &SplitRule{}
}

func (r *SplitRule) IsApplicable(n int) bool {
	return r.countDigits(n)%2 == 0
}

func (r *SplitRule) Apply(n int) []int {
	nStr := strconv.Itoa(n)
	leftHalf := nStr[0 : len(nStr)/2]
	rightHalf := nStr[len(nStr)/2:]
	// we can be sure these are numbers.
	leftInt, _ := strconv.Atoi(leftHalf)
	rightInt, _ := strconv.Atoi(rightHalf)
	return []int{leftInt, rightInt}
}

func (r *SplitRule) countDigits(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

type MultRule struct{}

func NewMultRule() *MultRule {
	return &MultRule{}
}

func (r *MultRule) IsApplicable(n int) bool {
	return true
}

func (r *MultRule) Apply(n int) []int {
	return []int{n * 2024}
}
