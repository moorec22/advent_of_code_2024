// Advent of Code, 2024, Day 21
//
// https://adventofcode.com/2024/day/21
//
// Part 1: This problem is a lot to take in. I decided the best way to think
// through this is with an incremental solution.
//
// First we start with what the robot operating the number pad has to press.
// this is trivial, it has to press the numbers.
//
// The robot operating the first directional pad has to do the following. For
// each button the robot operating the number pad has to press, the robot must
// get the numerical operator from their current position to the next position,
// then press A to press the button. I first find what buttons need to be
// pressed for this, but do not choose order. I return a list of maps. Each map
// maps directions to number of times we need to move in that direction, for the
// move in that position in the list. The A press for the robot operating the
// first directional pad is implicit. An empty map represents an A pressed without
// having to move.
//
// For the robot operating the second directional pad, for each directional key
// the first directional operator has to press, we need to move the robot to that
// button, and then press it. Immediately, we see that the shortest path will
// mean grouping the same buttons together so we can press A over and over without
// moving the robot unnecessarily. Again we use the directional map.
//
// In fact, we can use this same philosophy for any number of directional maps,
// converting a list of directional maps to a list of directional maps representing
// the pressing of the original map.
//
// To choose the order of which directions to press: there will only ever be four
// directional combinations: one of up and down, and one of left and right. Since
// we want to avoid the empty space, and the directions are independent of each other,
// we opt to go up or down first, and then go left or right.
//
// Now knowing what order to choose and knowing it will produce a smallest-character
// path, I decided to rewrite and simplify to strings for each layer.
package day21

import (
	"advent/util"
	"bufio"
	"fmt"
	"strconv"
)

var NumberPadPositions = map[rune]*util.Vector{
	'7': util.NewVector(0, 0),
	'8': util.NewVector(1, 0),
	'9': util.NewVector(2, 0),
	'4': util.NewVector(0, 1),
	'5': util.NewVector(1, 1),
	'6': util.NewVector(2, 1),
	'1': util.NewVector(0, 2),
	'2': util.NewVector(1, 2),
	'3': util.NewVector(2, 2),
	'0': util.NewVector(1, 3),
	'A': util.NewVector(2, 3),
}

var NumberPadGap = util.NewVector(0, 3)

var DirectionalPadPositions = map[rune]*util.Vector{
	'^': util.NewVector(1, 0),
	'A': util.NewVector(2, 0),
	'<': util.NewVector(0, 1),
	'v': util.NewVector(1, 1),
	'>': util.NewVector(2, 1),
}

var DirectionalPadGap = util.NewVector(0, 0)

// FOR TESTING
var DirectionalVectors = map[rune]*util.Vector{
	'^': util.NewVector(0, -1),
	'<': util.NewVector(-1, 0),
	'v': util.NewVector(0, 1),
	'>': util.NewVector(1, 0),
}

type Day21Solution struct {
	codes []string
}

func NewDay21Solution(filename string) (*Day21Solution, error) {
	codes := make([]string, 0)
	err := util.ProcessFile(filename, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			codes = append(codes, scanner.Text())
		}
		return scanner.Err()
	})
	return &Day21Solution{codes}, err
}

func (s *Day21Solution) PartOneAnswer() (int, error) {
	return s.getCodeComplexitySum(s.codes, 2)
}

func (s *Day21Solution) PartTwoAnswer() (int, error) {
	return s.getCodeComplexitySum(s.codes, 25)
}

// getCodeComplexitySum returns the sum of the complexities of each
// code. It returns an error if one of the codes is malformed.
func (s *Day21Solution) getCodeComplexitySum(codes []string, numDirectionalPads int) (int, error) {
	sum := 0
	for _, code := range codes {
		complexity, err := s.getCodeComplexity(code, numDirectionalPads)
		if err != nil {
			return sum, err
		}
		sum += complexity
	}
	return sum, nil
}

// getCodeComplexity returns the complexity of the given code, as dictated
// by the problem. It returns an error if the code is malformed.
func (s *Day21Solution) getCodeComplexity(code string, numDirectionalPads int) (int, error) {
	numericalCode, err := s.getNumericalCode(code)
	if err != nil {
		return 0, err
	}
	moves, err := s.getRobotMoves(code, NumberPadPositions, NumberPadGap)
	if err != nil {
		return 0, err
	}
	for range numDirectionalPads {
		moves, err = s.getPossibleRobotMoves(moves, DirectionalPadPositions, DirectionalPadGap)
		if err != nil {
			return 0, err
		}
	}
	shortest, err := s.getShortest(moves)
	return numericalCode * len(shortest), err
}

// getPossibleRobotMoves takes a set of codes, and returns all possible moves
// a robot can do to make those codes, given the positions and the gap.
func (s *Day21Solution) getPossibleRobotMoves(codes []string, positions map[rune]*util.Vector, gap *util.Vector) ([]string, error) {
	possibleMoves := make([]string, 0)
	for _, code := range codes {
		possibleMove, err := s.getRobotMoves(code, positions, gap)
		if err != nil {
			return possibleMoves, err
		}
		possibleMoves = append(possibleMoves, possibleMove...)
	}
	return possibleMoves, nil
}

// getRobotMoves takes the code the pad operator has to press. It returns a
// string representing a combination of directional keys to press to get the
// code entered, with some constraints:
//   - any directions needed to be pressed multiple times will be grouped
//   - we always go up and down first, then left to right
//
// It returns an error if the code includes keys not found in positions.
func (s *Day21Solution) getRobotMoves(code string, positions map[rune]*util.Vector, gap *util.Vector) ([]string, error) {
	currentButton := 'A'
	paths := util.NewArrayQueue[string]()
	paths.Insert("")
	for _, nextButton := range code {
		newPaths, err := s.possiblePresses(currentButton, nextButton, positions, gap)
		if err != nil {
			return paths.ToArray(), err
		}
		size := paths.Size()
		for range size {
			prefix := paths.Remove()
			for _, newPath := range newPaths {
				paths.Insert(prefix + newPath)
			}
		}
		currentButton = nextButton
	}
	return paths.ToArray(), nil
}

// possiblePresses takes a start and an end, and finds all the possible
// presses that can be made to get there, excluding ones that we know to be too
// long.
func (s *Day21Solution) possiblePresses(start, end rune, positions map[rune]*util.Vector, gap *util.Vector) ([]string, error) {
	startPosition, ok := positions[start]
	if !ok {
		return nil, fmt.Errorf("code not found in positions: '%v'", start)
	}
	endPosition, ok := positions[end]
	if !ok {
		return nil, fmt.Errorf("code not found in positions: '%v'", end)
	}
	distance := endPosition.GetManhattanDistance(startPosition)
	xString, yString := s.getDirectionalSegments(distance)
	if xString == "" {
		return []string{yString + "A"}, nil
	} else if yString == "" {
		return []string{xString + "A"}, nil
	} else if util.NewVector(startPosition.X, endPosition.Y).Equals(gap) {
		return []string{xString + yString + "A"}, nil
	} else if util.NewVector(endPosition.X, startPosition.Y).Equals(gap) {
		return []string{yString + xString + "A"}, nil
	} else {
		return []string{xString + yString + "A", yString + xString + "A"}, nil
	}
}

// getDirectionalCharacters takes a vector, and returns the segments needed
// to traverse horizontally and vertically. If no traversal is needed for one,
// an empty string is chosen.
func (s *Day21Solution) getDirectionalSegments(v *util.Vector) (string, string) {
	xRune := rune(-1)
	yRune := rune(-1)
	if v.X < 0 {
		xRune = '<'
	} else if v.X > 0 {
		xRune = '>'
	}
	if v.Y < 0 {
		yRune = '^'
	} else if v.Y > 0 {
		yRune = 'v'
	}
	xString := ""
	yString := ""
	for range util.IntAbs(v.X) {
		xString += string(xRune)
	}
	for range util.IntAbs(v.Y) {
		yString += string(yRune)
	}
	return xString, yString
}

// Takes a code in the form [0-9]*A, and returns the numerical value preceeding
// A. Returns an error if the conversion fails.
func (s *Day21Solution) getNumericalCode(code string) (int, error) {
	return strconv.Atoi(code[:len(code)-1])
}

// FOR TESTING, converts button presses to what the robot does
func (s *Day21Solution) getPresses(moves string, positions map[rune]*util.Vector) string {
	currentPosition := positions['A']
	reversePositions := make(map[util.Vector]rune)
	for k, v := range positions {
		reversePositions[*v] = k
	}
	presses := ""
	for _, move := range moves {
		if move == 'A' {
			presses += string(reversePositions[*currentPosition])
		} else {
			currentPosition = currentPosition.Add(DirectionalVectors[move])
		}
	}
	return presses
}

// getShortest returns the shortest string in codes.
func (s *Day21Solution) getShortest(codes []string) (string, error) {
	if len(codes) == 0 {
		return "", fmt.Errorf("no codes in array")
	}
	shortest := ""
	for _, code := range codes {
		if shortest == "" || len(code) < len(shortest) {
			shortest = code
		}
	}
	return shortest, nil
}
