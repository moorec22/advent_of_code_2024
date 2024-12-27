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
//
// INCOMPLETE: I couldn't get the right answer for part 2, and after several
// days could not figure out why. I ended up using the solution found
// here https://github.com/AllanTaylor314/AdventOfCode/blob/main/2024/21.py
// for now.
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

type Transition struct {
	from rune
	to   rune
}

var DirectionalTransitionSequences = map[Transition]string{
	{'A', 'A'}: "A",
	{'A', '^'}: "<A",
	{'A', '<'}: "v<<A",
	{'A', 'v'}: "<vA",
	{'A', '>'}: "vA",

	{'^', 'A'}: ">A",
	{'^', '^'}: "A",
	{'^', '>'}: "v>A",
	{'^', '<'}: "v<A",

	{'<', 'A'}: ">>^A",
	{'<', '^'}: ">^A",
	{'<', 'v'}: ">A",
	{'<', '<'}: "A",

	{'v', 'A'}: "^>A",
	{'v', 'v'}: "A",
	{'v', '>'}: ">A",
	{'v', '<'}: "<A",

	{'>', 'A'}: "^A",
	{'>', '^'}: "<^A",
	{'>', 'v'}: "<A",
	{'>', '>'}: "A",
}

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
	possibleMoves, err := s.getNumericalRobotMoves(code)
	if err != nil {
		return 0, err
	}
	minCount := -1
	for _, moves := range possibleMoves {
		transitions := s.getTransitionMap(moves)
		for range numDirectionalPads {
			transitions, err = s.getNextDirectionalMoves(transitions)
			if err != nil {
				return 0, err
			}
		}
		count := s.getTotalCount(transitions)
		if minCount == -1 || count < minCount {
			minCount = count
		}
	}
	return numericalCode * minCount, err
}

// getNumericalPossibleRobotMoves takes a numerical code, and returns an
// optimal directional pad sequence to make those moves. It returns an error
// if any digit is not valid.
func (s *Day21Solution) getNumericalRobotMoves(code string) ([]string, error) {
	currentDigit := 'A'
	moves := util.NewArrayQueue[string]()
	moves.Insert("")
	for _, nextDigit := range code {
		sequences, err := s.getNumberPadSequences(currentDigit, nextDigit)
		if err != nil {
			return nil, err
		}
		size := moves.Size()
		for range size {
			baseSequence := moves.Remove()
			for _, sequence := range sequences {
				moves.Insert(baseSequence + sequence)
			}
		}
		currentDigit = nextDigit
	}
	return moves.ToArray(), nil
}

// getNextDirectionalMoves takes a map of transitions to the number of
// transitions made, and returns moves made by the next robot up in the form
// of a transition map for that robot.
func (s *Day21Solution) getNextDirectionalMoves(transitions map[Transition]int) (map[Transition]int, error) {
	newTransitions := make(map[Transition]int)
	for transition, count := range transitions {
		sequence, ok := DirectionalTransitionSequences[transition]
		if !ok {
			return newTransitions, fmt.Errorf("transition not found in sequences: %s\v", sequence)
		}
		subTransitions := s.getTransitionMap(sequence)
		for subTransition, subCount := range subTransitions {
			oldCount, ok := newTransitions[subTransition]
			if !ok {
				oldCount = 0
			}
			newTransitions[subTransition] = oldCount + subCount*count
		}
	}
	return newTransitions, nil
}

// buttonPresses takes a start and end, and returns an optimal seqeuence from
// the start button to the end button.
func (s *Day21Solution) getNumberPadSequences(start, end rune) ([]string, error) {
	startPosition, ok := NumberPadPositions[start]
	if !ok {
		return nil, fmt.Errorf("code not found in positions: '%v'", start)
	}
	endPosition, ok := NumberPadPositions[end]
	if !ok {
		return nil, fmt.Errorf("code not found in positions: '%v'", end)
	}
	distance := endPosition.GetManhattanDistance(startPosition)
	xString, yString := s.getDirectionalSegments(distance)
	// order: left, (up, right), down
	if xString == "" {
		return []string{yString + "A"}, nil
	} else if yString == "" {
		return []string{xString + "A"}, nil
	}
	if startPosition.Y == 3 && endPosition.X == 0 {
		// we need to go up first, otherwise we'll hit the gap
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

func (s *Day21Solution) getTransitionMap(moves string) map[Transition]int {
	transitions := make(map[Transition]int)
	currentButton := 'A'
	for _, nextButton := range moves {
		transition := Transition{currentButton, nextButton}
		_, ok := transitions[transition]
		if !ok {
			transitions[transition] = 0
		}
		transitions[transition]++
		currentButton = nextButton
	}
	return transitions
}

func (s *Day21Solution) getTotalCount(transitions map[Transition]int) int {
	count := 0
	for _, v := range transitions {
		count += v
	}
	return count
}
