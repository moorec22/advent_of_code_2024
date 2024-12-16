// Advent of Code, 2024, Day 13
//
// https://adventofcode.com/2024/day/13
//
// Part 1:
// Each solution, representing how many times to press A and how many times to
// press B such that we spend the least tokens, can be represented in systems
// of equations:
//
//  1. S_X = (A_X)A + (B_X)B
//  2. S_Y = (A_Y)A + (B_Y)B
//
// where (S_X, S_Y) are the coordinates of the solution, (A_X, A_Y) are how far
// the claw is moved in the X and Y directions when A is pressed, (B_X, B_Y)
// are how far the claw is moved in the X and Y directions when B is pressed,
// and A and B are the number of times A and B are pressed.
//
// We can solve this system of equations by multiplying the first equation by
// B_Y and the second equation by B_X, and then subtracting the two equations.
// This gives us:
//
//	B_Y * S_X = B_Y * A_X * A + B_Y * B_X * B
//	B_X * S_Y = B_X * A_Y * A + B_X * B_Y * B
//	(B_Y * S_X) - (B_X * S_Y) = (B_Y * A_X - B_X * A_Y) * A
//	A = (B_Y * S_X - B_X * S_Y) / (B_Y * A_X - B_X * A_Y)
//
// If in this last equation the donimator is 0, then we have an infinite number
// of solutions. At that point, we have ot factor in the tokens spent as a way
// to find the minimum solution. However, there are no such cases in the input.
//
// We can then find B:
//
//	B = (S_X - A_X * A) / B_X
//
// If A or B in these last two equations are not integers, then we have no
// solution.
//
// Part 2: Thankfully, because we've used a system of linear equations to solve
// the problem, we can easily extend the solution to part 2 by making the adjustment
// to the answer necessary.
package day13

import (
	"advent/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const ButtonLinePrefixLength = 10
const SolutionLinePrefixLength = 7
const Part2Adjustment = 10000000000000

type Equation struct {
	Solution int
	ButtonA  int
	ButtonB  int
}

type EquationSystem struct {
	XEquation Equation
	YEquation Equation
}

type Day13Solution struct {
	equationSystems []*EquationSystem
}

func NewDay13Solution(filepath string) (*Day13Solution, error) {
	equationSystems := make([]*EquationSystem, 0)
	err := util.ProcessFile(filepath, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			equationSystem, err := getEquationSystem(scanner)
			if err != nil {
				return err
			}
			equationSystems = append(equationSystems, equationSystem)
		}
		return nil
	})
	return &Day13Solution{equationSystems}, err
}

func (s *Day13Solution) PartOneAnswer() (int, error) {
	return s.getFewestTokensNeeded(s.equationSystems, false), nil
}

func (s *Day13Solution) PartTwoAnswer() (int, error) {
	return s.getFewestTokensNeeded(s.equationSystems, true), nil
}

// getFewestTokensNeeded returns the fewest number of tokens needed to solve the
// equation system. If adjustSolution is true, the solution is adjusted by the
// Part2Adjustment constant.
func (s *Day13Solution) getFewestTokensNeeded(equationSystems []*EquationSystem, adjustSolution bool) int {
	tokens := 0
	for _, equationSystem := range equationSystems {
		if adjustSolution {
			equationSystem = adjustEquationSystem(*equationSystem, Part2Adjustment)
		}
		a, solvable := s.getA(equationSystem)
		if !solvable {
			continue
		}
		b, solvable := s.getB(equationSystem, a)
		if !solvable {
			continue
		}
		if !adjustSolution && (a > 100 || b > 100) {
			continue
		}
		tokens += 3*a + b
	}
	return tokens
}

// getA solves for A by evaluating
// A = (B_Y * S_X - B_X * S_Y) / (B_Y * A_X - B_X * A_Y).
// If A is not an integer, the second return value is false.
func (s *Day13Solution) getA(equationSystem *EquationSystem) (int, bool) {
	numerator := equationSystem.YEquation.ButtonB*equationSystem.XEquation.Solution -
		equationSystem.XEquation.ButtonB*equationSystem.YEquation.Solution
	denominator := equationSystem.YEquation.ButtonB*equationSystem.XEquation.ButtonA -
		equationSystem.XEquation.ButtonB*equationSystem.YEquation.ButtonA
	// If the denominator is 0, we have an infinite number of solutions.
	if numerator%denominator != 0 {
		return 0, false
	}
	return numerator / denominator, true
}

// given A, solves for B in the equation system.
// If B is not an integer, the second return value is false.
func (s *Day13Solution) getB(equationSystem *EquationSystem, a int) (int, bool) {
	numerator := equationSystem.XEquation.Solution - equationSystem.XEquation.ButtonA*a
	denominator := equationSystem.XEquation.ButtonB
	return numerator / denominator, numerator%denominator == 0
}

// getEquationSystem returns an EquationSystem from the given scanner. If the
// scanner reaches the end of the file or if the file is not in the format
// specified by the problem, an error is returned.
func getEquationSystem(scanner *bufio.Scanner) (*EquationSystem, error) {
	aButtonLine := scanner.Text()
	nextLine := scanner.Scan()
	if !nextLine {
		return nil, fmt.Errorf("unexpected end of file")
	}
	bButtonLine := scanner.Text()
	nextLine = scanner.Scan()
	if !nextLine {
		return nil, fmt.Errorf("unexpected end of file")
	}
	solutionLine := scanner.Text()
	aButtonValues, err := getButtonValues(aButtonLine)
	if err != nil {
		return nil, err
	}
	bButtonValues, err := getButtonValues(bButtonLine)
	if err != nil {
		return nil, err
	}
	solutionValues, err := getSolutionValues(solutionLine)
	if err != nil {
		return nil, err
	}
	xEquation := getEquation(aButtonValues, bButtonValues, solutionValues, 0)
	yEquation := getEquation(aButtonValues, bButtonValues, solutionValues, 1)
	scanner.Scan()
	return &EquationSystem{xEquation, yEquation}, nil
}

// getButtonValues returns the numbers from a string of the form "ccccccccccccN, ccN, ccN, ..."
// where cc is a two-character prefix and N is a number.
func getButtonValues(line string) ([]int, error) {
	return getValues(line[ButtonLinePrefixLength:])
}

// getSolutionValues returns the numbers from a string of the form "cccccccccN, ccN, ccN, ..."
// where cc is a two-character prefix and N is a number.
func getSolutionValues(line string) ([]int, error) {
	return getValues(line[SolutionLinePrefixLength:])
}

// getValues returns the numbers from a string of the form "ccN, ccN, ccN, ..."
// where cc is a two-character prefix and N is a number.
func getValues(line string) ([]int, error) {
	parts := strings.Split(line, ", ")
	values := make([]int, len(parts))
	for i, part := range parts {
		value, err := strconv.Atoi(part[2:])
		if err != nil {
			return nil, fmt.Errorf("invalid button value: %s", part)
		}
		values[i] = value
	}
	return values, nil
}

// getEquation returns the equation from the given arrays at position x.
func getEquation(aButton, bButton, solution []int, x int) Equation {
	return Equation{
		Solution: solution[x],
		ButtonA:  aButton[x],
		ButtonB:  bButton[x],
	}
}

// adjustEquationSystem returns a new EquationSystem with the given adjustment
// made to the solution values
func adjustEquationSystem(equationSystem EquationSystem, adjustment int) *EquationSystem {
	return &EquationSystem{
		XEquation: adjustEquation(equationSystem.XEquation, adjustment),
		YEquation: adjustEquation(equationSystem.YEquation, adjustment),
	}
}

// adjustEquation returns a new Equation with the given adjustment made to the
// solution value.
func adjustEquation(equation Equation, adjustment int) Equation {
	return Equation{
		Solution: equation.Solution + adjustment,
		ButtonA:  equation.ButtonA,
		ButtonB:  equation.ButtonB,
	}
}
