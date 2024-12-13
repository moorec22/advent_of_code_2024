// Advent of Code, 2024, Day 7
//
// https://adventofcode.com/2024/day/7
//
// Part 1 Idea: Naive solution, try all operator possibilities
package day07

import (
	"advent/util"
	"bufio"
	"strconv"
	"strings"
)

type Equation struct {
	left  int
	right []int
}

type Operator interface {
	apply(int, int) int
}

type Add struct{}

func (a Add) apply(a1, a2 int) int {
	return a1 + a2
}

type Multiply struct{}

func (m Multiply) apply(a1, a2 int) int {
	return a1 * a2
}

type Concatenate struct{}

func (c Concatenate) apply(a1, a2 int) int {
	strA1 := strconv.Itoa(a1)
	strA2 := strconv.Itoa(a2)
	// we can throw away the error, being certain that the conversion will work
	concat, _ := strconv.Atoi(strA1 + strA2)
	return concat
}

type Day07Solution struct {
	equations []Equation
}

func NewDay07Solution(filepath string) (*Day07Solution, error) {
	equations, err := getEquations(filepath)
	return &Day07Solution{equations}, err
}

func (s *Day07Solution) PartOneAnswer() (int, error) {
	validEquations := s.validEquations([]Operator{Add{}, Multiply{}})
	return s.leftSideSum(validEquations), nil
}

func (s *Day07Solution) PartTwoAnswer() (int, error) {
	validEquations := s.validEquations([]Operator{Add{}, Multiply{}, Concatenate{}})
	return s.leftSideSum(validEquations), nil
}

// leftSideSum returns the sum of the left side of the equations.
func (s *Day07Solution) leftSideSum(equations []Equation) int {
	sum := 0
	for _, e := range equations {
		sum += e.left
	}
	return sum
}

// validEquations returns a slice of equations that can be made valid using
// any combination of the operators.
func (s *Day07Solution) validEquations(operators []Operator) []Equation {
	validEquations := make([]Equation, 0)
	for _, e := range s.equations {
		if s.validEquation(e, operators) {
			validEquations = append(validEquations, e)
		}
	}
	return validEquations
}

// validEquation returns true if the given equation can be made valid using
// any combination of the operators.
func (s *Day07Solution) validEquation(e Equation, operators []Operator) bool {
	return s.validEquationHelper(e, operators, 0, 0)
}

// validEquationHelper is a helper function for validEquation. e is the
// equation, operators is the list of operators to try, currentVal is the
// current partial value, and i is the current index of the right side.
// The function returns true if it can be made valid using any combination
// of the operators for the remainder of the function. If the function has
// reached the end of the right side, it returns true if the current value
// is equal to the left side of the equation.
func (s *Day07Solution) validEquationHelper(e Equation, operators []Operator, currentVal, i int) bool {
	if i == len(e.right) {
		return currentVal == e.left
	}
	for _, op := range operators {
		if s.validEquationHelper(e, operators, op.apply(currentVal, e.right[i]), i+1) {
			return true
		}
	}
	return false
}

// getEquations reads the file at the given filepath and returns a slice of
// Equation structs. The file needs to be in the format specified by the
// problem.
func getEquations(filename string) ([]Equation, error) {
	equations := []Equation{}
	err := util.ProcessFile(filename, func(s *bufio.Scanner) error {
		for s.Scan() {
			line := s.Text()
			lineStrs := strings.Split(line, ":")
			leftNum, err := strconv.Atoi(lineStrs[0])
			if err != nil {
				return err
			}
			rightStrs := strings.Split(strings.TrimSpace(lineStrs[1]), " ")
			rightNums := []int{}
			for _, rightStr := range rightStrs {
				rightNum, err := strconv.Atoi(rightStr)
				if err != nil {
					return err
				}
				rightNums = append(rightNums, rightNum)
			}
			equations = append(equations, Equation{left: leftNum, right: rightNums})
		}
		return nil
	})
	return equations, err
}
