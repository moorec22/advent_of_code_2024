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
//     based on their X and Yumn, we could quickly check if it's worth putting
//     a new obstacle in the guard's path at each step.
//   - keeping track of coordinates where the guard has turned. This helps indicate
//     if the guard is looping. [DONE]
package day06

import (
	"advent/util"
	"fmt"
)

type Day06Solution struct {
	initialLabMap      util.Matrix[rune]
	initialGuardVector util.Vector
}

func NewDay06Solution(filepath string) (*Day06Solution, error) {
	labMap, guard, err := getLabMapAndGuard(filepath)
	return &Day06Solution{initialLabMap: labMap, initialGuardVector: guard}, err
}

func (s *Day06Solution) PartOneAnswer() (int, error) {
	labMap := s.getMatrixCopy(s.initialLabMap)
	seenVectors, err := s.trackGuard(labMap, s.initialGuardVector)
	return len(seenVectors), err
}

func (s *Day06Solution) PartTwoAnswer() (int, error) {
	labMap := s.getMatrixCopy(s.initialLabMap)
	seenVectors, err := s.trackGuard(labMap, s.initialGuardVector)
	if err != nil {
		return 0, err
	}
	labMap = s.getMatrixCopy(s.initialLabMap)
	steps, err := s.countLoops(labMap, s.initialGuardVector, seenVectors)
	return steps, err
}

// getLabMapAndGuard returns the lab map and the guard's Vector in the map,
// or an error if the file cannot be processed. The file is in the format
// described in the prompt.
func getLabMapAndGuard(filename string) (util.Matrix[rune], util.Vector, error) {
	labMap, err := util.ParseMatrix(filename, func(r rune) rune {
		return r
	})
	if err != nil {
		return labMap, util.Vector{}, err
	}
	for i, X := range labMap {
		for j, r := range X {
			if isGuard(r) {
				return labMap, util.NewVector(i, j), nil
			}
		}
	}
	return labMap, util.Vector{}, fmt.Errorf("no guard found in lab map")
}

// trackGuard returns all known locations the guard visits on their path. It is not
// guaranteed that labMap will be unchanged by this function.
func (s *Day06Solution) trackGuard(labMap util.Matrix[rune], guardPos util.Vector) (map[util.Vector]bool, error) {
	var err error
	seenVectors := make(map[util.Vector]bool)
	for labMap.PosInBounds(guardPos) {
		if _, ok := seenVectors[guardPos]; !ok {
			seenVectors[guardPos] = true
		}
		guardPos, err = s.moveToNextVector(labMap, guardPos)
		if err != nil {
			return seenVectors, err
		}
	}
	return seenVectors, nil
}

// countLoops returns the number of obstacles that would cause the guard to loop. It needs the lab map,
// the starting Vector of the guard, and all Vectors the guard is seen at on her original path.
// It is not guaranteed that labMap will be unchanged by this function.
func (s *Day06Solution) countLoops(labMap util.Matrix[rune], guardPos util.Vector, seenVectors map[util.Vector]bool) (int, error) {
	delete(seenVectors, guardPos)
	guard := labMap.Get(guardPos)
	obstacleVectorCount := 0
	for pos := range seenVectors {
		labMap.Set(pos, Obstacle)
		looping, err := s.isLooping(labMap, guardPos)
		if err != nil {
			return 0, err
		}
		if looping {
			obstacleVectorCount++
		}
		labMap.Set(pos, Empty)
		labMap.Set(guardPos, guard)
	}
	return obstacleVectorCount, nil
}

// isLooping returns true if the guard is looping in the labMap, false otherwise. An error is returned
// if there is a problem moving the guard.
func (s *Day06Solution) isLooping(labMap util.Matrix[rune], guardPos util.Vector) (bool, error) {
	seenTurns := make(map[util.Vector]util.Vector)
	for labMap.PosInBounds(guardPos) {
		obstacleDirection, ok := seenTurns[guardPos]
		currentDirection, err := s.getGuardDirection(labMap.Get(guardPos))
		if err != nil {
			return false, err
		}
		if ok && s.isInFrontOfObstacle(labMap, guardPos) && obstacleDirection == currentDirection {
			// the guard has turned here before -- she is looping
			labMap.Set(guardPos, Empty)
			return true, nil
		}
		seenTurns[guardPos] = currentDirection
		guardPos, err = s.moveToNextVector(labMap, guardPos)
		if err != nil {
			return false, err
		}
	}
	return false, nil
}

// moveToNextVector moves the guard to the next Vector in the labMap. If the
// guard encounters an obstacle, she turns right and steps one forward. Otherwise,
// she steps one forward in the direction she is currently moving.
func (s *Day06Solution) moveToNextVector(labMap util.Matrix[rune], guardPos util.Vector) (util.Vector, error) {
	guard := labMap.Get(guardPos)
	nextVector, err := s.getNextVector(guardPos, guard)
	if err != nil {
		return guardPos, err
	}
	labMap.Set(guardPos, Empty)
	if labMap.PosInBounds(nextVector) {
		if isObstacle(labMap.Get(nextVector)) {
			newDirection, err := s.getRightTurn(guard)
			if err != nil {
				return guardPos, err
			}
			labMap.Set(guardPos, newDirection)
			return s.moveToNextVector(labMap, guardPos)
		} else {
			labMap.Set(nextVector, guard)
		}
	}
	return nextVector, nil
}

// isInFrontOfObstacle returns true if the guard is in front of an obstacle in the labMap.
func (s *Day06Solution) isInFrontOfObstacle(labMap util.Matrix[rune], guardPos util.Vector) bool {
	guard := labMap.Get(guardPos)
	nextPos, err := s.getNextVector(guardPos, guard)
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
	_, ok := guardDirections[r]
	return ok
}

// getGuardDirection returns the direction of the guard represented by r, or an
// error if r is not a guard.
func (s *Day06Solution) getGuardDirection(r rune) (util.Vector, error) {
	vector, ok := guardDirections[r]
	if !ok {
		return util.NewVector(0, 0), fmt.Errorf("rune %c is not a guard", r)
	}
	return vector, nil
}

func (s *Day06Solution) getNextVector(currentPos util.Vector, guard rune) (util.Vector, error) {
	d, err := s.getGuardDirection(guard)
	return currentPos.Add(d), err
}

// getRightTurn returns the direction the guard should turn to if it encounters
// an obstacle.
func (s *Day06Solution) getRightTurn(r rune) (rune, error) {
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

func (s *Day06Solution) getMatrixCopy(m util.Matrix[rune]) util.Matrix[rune] {
	copy := util.NewMatrix[rune]()
	for _, X := range m {
		newX := make([]rune, len(X))
		for i, r := range X {
			newX[i] = r
		}
		copy = append(copy, newX)
	}
	return copy
}
