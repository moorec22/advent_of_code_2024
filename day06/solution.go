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
//   - keeping track of coordinates where the guard has turned. This helps indicate
//     if the guard is looping. [DONE]
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
	seenPositions, err := trackGuard(labMap, guard)
	return len(seenPositions), err
}

func PartTwoAnswer(filepath string) (int, error) {
	labMap, guard, err := getLabMapAndGuard(filepath)
	if err != nil {
		return 0, err
	}
	seenPositions, err := trackGuard(getMatrixCopy(labMap), guard)
	if err != nil {
		return 0, err
	}
	steps, err := countLoops(labMap, guard, seenPositions)
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

// trackGuard returns all known locations the guard visits on their path.
func trackGuard(labMap util.Matrix[rune], guardPos util.Position) (map[util.Position]bool, error) {
	var err error
	seenPositions := make(map[util.Position]bool)
	for labMap.PosInBounds(guardPos) {
		if _, ok := seenPositions[guardPos]; !ok {
			seenPositions[guardPos] = true
		}
		guardPos, err = moveToNextPosition(labMap, guardPos)
		if err != nil {
			return seenPositions, err
		}
	}
	return seenPositions, nil
}

// countLoops returns the number of obstacles that would cause the guard to loop. It needs the lab map,
// the starting position of the guard, and all positions the guard is seen at on her original path.
func countLoops(labMap util.Matrix[rune], guardPos util.Position, seenPositions map[util.Position]bool) (int, error) {
	delete(seenPositions, guardPos)
	guard := labMap.Get(guardPos)
	obstaclePositionCount := 0
	for pos := range seenPositions {
		labMap.Set(pos, Obstacle)
		looping, err := isLooping(labMap, guardPos)
		if err != nil {
			return 0, err
		}
		if looping {
			obstaclePositionCount++
		}
		labMap.Set(pos, Empty)
		labMap.Set(guardPos, guard)
	}
	return obstaclePositionCount, nil
}

// isLooping returns true if the guard is looping in the labMap, false otherwise. An error is returned
// if there is a problem moving the guard.
func isLooping(labMap util.Matrix[rune], guardPos util.Position) (bool, error) {
	seenTurns := make(map[util.Position]Direction)
	for labMap.PosInBounds(guardPos) {
		obstacleDirection, ok := seenTurns[guardPos]
		currentDirection, err := getGuardDirection(labMap.Get(guardPos))
		if err != nil {
			return false, err
		}
		if ok && isInFrontOfObstacle(labMap, guardPos) && obstacleDirection == currentDirection {
			// the guard has turned here before -- she is looping
			labMap.Set(guardPos, Empty)
			return true, nil
		}
		seenTurns[guardPos] = currentDirection
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

// isInFrontOfObstacle returns true if the guard is in front of an obstacle in the labMap.
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

func getMatrixCopy(m util.Matrix[rune]) util.Matrix[rune] {
	copy := util.NewMatrix[rune]()
	for _, row := range m {
		newRow := make([]rune, len(row))
		for i, r := range row {
			newRow[i] = r
		}
		copy = append(copy, newRow)
	}
	return copy
}
