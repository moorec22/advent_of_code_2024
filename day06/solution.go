// Advent of Code, 2024, Day 6
//
// https://adventofcode.com/2024/day/6
//
// Part 1: Naive solution, just follow the guard step by step
// Part 2: Could fundamentally implement naive solution, of placing an obstacle
// in every step of the guard's path, and then seeing if the guard is stuck in
// a loop.
//
// Optimization ideas:
//   - could save a matrix of vectors, indicating the distance to an obstacle in
//     each direction. This would help for both parts.
//   - an obstacle will only put the guard in a loop if there is another obstacle
//     in her path when she turns right. If we store a data structure of obstacles
//     based on their row and column, we could quickly check if it's worth putting
//     a new obstacle in the guard's path at each step.
//   - keeping track of not only where a guard has been, but what direction that
//     guard stepped in last time. This would help us determine if the guard is
//     looping.
//   - keeping track of coordinates where the guard has turned. This helps indicate
//     if the guard is looping.
package day06

import (
	"advent/util"
	"bufio"
	"fmt"
)

func PartOneAnswer(filepath string) (int, error) {
	labMap, guard, err := getLabMapAndGuard(filepath)
	if err != nil {
		return 0, err
	}
	steps, err := trackGuard(labMap, guard)
	return steps, err
}

func PartTwoAnswer(filepath string) (int, error) {
	labMap, guard, err := getLabMapAndGuard(filepath)
	if err != nil {
		return 0, err
	}
	steps, err := countLoops(labMap, guard)
	return steps, err
}

// getLabMapAndGuard returns the lab map and the guard's position in the map,
// or an error if the file cannot be processed. The file is in the format
// described in the prompt.
func getLabMapAndGuard(filename string) (util.Matrix[rune], util.Position, error) {
	labMap := util.NewMatrix[rune]()
	guardPosition := util.Position{}
	err := util.ProcessFile(filename, func(s *bufio.Scanner) error {
		for s.Scan() {
			line := s.Text()
			row := make([]rune, 0)
			for j, r := range line {
				row = append(row, r)
				if isGuard(r) {
					guardPosition = util.Position{Row: len(labMap), Col: j}
				}
			}
			labMap = append(labMap, row)
		}
		return nil
	})
	return labMap, guardPosition, err
}

// trackGuard returns the number of steps the guard walks in the map before
// leaving the map
func trackGuard(labMap util.Matrix[rune], guardPos util.Position) (int, error) {
	var err error
	seenPositions := make(map[util.Position]bool)
	for labMap.PosInBounds(guardPos) {
		if _, ok := seenPositions[guardPos]; !ok {
			seenPositions[guardPos] = true
		}
		guardPos, err = moveToNextPosition(labMap, guardPos)
		if err != nil {
			return 0, err
		}
	}
	return len(seenPositions), nil
}

func countLoops(labMap util.Matrix[rune], guardPos util.Position) (int, error) {
	loops := 0
	for labMap.PosInBounds(guardPos) {
		nextPos, err := getNextPosition(guardPos, labMap.Get(guardPos))
		if err != nil {
			return 0, err
		}
		if labMap.PosInBounds(nextPos) && !isObstacle(labMap.Get(nextPos)) {
			guard := labMap.Get(guardPos)
			labMap.Set(nextPos, 'O')
			isLoop, err := isLooping(labMap, guardPos)
			if err != nil {
				return 0, err
			}
			if isLoop {
				printMatrix(tracePath(labMap, guardPos))
				fmt.Println()
				loops++
			}
			labMap.Set(nextPos, Empty)
			labMap.Set(guardPos, guard)
		}
		guardPos, err = moveToNextPosition(labMap, guardPos)
		if err != nil {
			return 0, err
		}
	}
	return loops, nil
}

func isLooping(labMap util.Matrix[rune], guardPos util.Position) (bool, error) {
	seenTurns := make(map[util.Position]bool)
	var err error
	for labMap.PosInBounds(guardPos) {
		_, ok := seenTurns[guardPos]
		if ok && isInFrontOfObstacle(labMap, guardPos) {
			// the guard has turned here before -- she is looping
			return true, nil
		}
		seenTurns[guardPos] = true
		guardPos, err = moveToNextPosition(labMap, guardPos)
		if err != nil {
			return false, err
		}
	}
	return false, nil
}

// moveToNextPosition moves the guard to the next position in the labMap. If the
// guard encounters an obstacle, she turns right and steps one forward. Otherwise,
// she steps one forward in the direction she is currently moving.
func moveToNextPosition(labMap util.Matrix[rune], guardPos util.Position) (util.Position, error) {
	guard := labMap.Get(guardPos)
	nextPosition, err := getNextPosition(guardPos, guard)
	if err != nil {
		return guardPos, err
	}
	labMap.Set(guardPos, Empty)
	if labMap.PosInBounds(nextPosition) {
		if isObstacle(labMap.Get(nextPosition)) {
			newDirection, err := getRightTurn(guard)
			if err != nil {
				return guardPos, err
			}
			labMap.Set(guardPos, newDirection)
			return moveToNextPosition(labMap, guardPos)
		} else {
			labMap.Set(nextPosition, guard)
		}
	}
	return nextPosition, nil
}

func isInFrontOfObstacle(labMap util.Matrix[rune], guardPos util.Position) bool {
	guard := labMap.Get(guardPos)
	nextPos, err := getNextPosition(guardPos, guard)
	if err != nil {
		return false
	}
	return labMap.PosInBounds(nextPos) && isObstacle(labMap.Get(nextPos))
}

// isObstacle returns true if r is not a '.'
func isObstacle(r rune) bool {
	return r == Obstacle || r == 'O'
}

// isGuard returns true if r is '^', 'v', '<', or '>'
func isGuard(r rune) bool {
	_, ok := guardPositions[r]
	return ok
}

// getGuardDirection returns the direction of the guard represented by r, or an
// error if r is not a guard.
func getGuardDirection(r rune) (Direction, error) {
	position, ok := guardPositions[r]
	if !ok {
		return -1, fmt.Errorf("rune %c is not a guard", r)
	}
	return position, nil
}

func getNextPosition(currentPos util.Position, guard rune) (util.Position, error) {
	d, err := getGuardDirection(guard)
	return currentPos.Add(unitVectors[d]), err
}

// getRightTurn returns the direction the guard should turn to if it encounters
// an obstacle.
func getRightTurn(r rune) (rune, error) {
	switch r {
	case '^':
		return '>', nil
	case '>':
		return 'v', nil
	case 'v':
		return '<', nil
	case '<':
		return '^', nil
	default:
		return ' ', fmt.Errorf("rune %c is not a guard", r)
	}
}

func tracePath(labMap util.Matrix[rune], guardPos util.Position) util.Matrix[rune] {
	trace := util.NewMatrix[rune]()
	for _, row := range labMap {
		traceRow := make([]rune, len(row))
		copy(traceRow, row)
		trace = append(trace, traceRow)
	}
	seenTurns := make(map[util.Position]bool)
	for trace.PosInBounds(guardPos) {
		if isInFrontOfObstacle(trace, guardPos) {
			trace[guardPos.Row][guardPos.Col] = '+'
			_, ok := seenTurns[guardPos]
			if ok {
				// the guard has turned here before -- she is looping
				return trace
			}
			seenTurns[guardPos] = true
		} else {
			direction, _ := getGuardDirection(trace.Get(guardPos))
			if direction == Up || direction == Down {
				trace[guardPos.Row][guardPos.Col] = '|'
			} else {
				trace[guardPos.Row][guardPos.Col] = '-'
			}
		}
		guardPos, _ = moveToNextPosition(trace, guardPos)
	}
	return trace
}

func printMatrix(m util.Matrix[rune]) {
	for _, row := range m {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}
