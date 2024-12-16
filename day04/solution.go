// Advent of Code, 2024, Day 4
//
// https://adventofcode.com/2024/day/4
//
// For the part one of this solution I decided to use a dynamic programming
// approach to solve the problem. The idea is to keep track of the progress of
// the word in the matrix in all directions. For each cell in the matrix, we
// keep track of the progress of the word in the forward and backward
// directions in the horizontal, vertical, and diagonal directions. We then
// count the number of times the word appears in the matrix by counting the
// number of trackers that have reached the end of the word. It may be that
// looking for all X'es and seeing if they're part of the word would be simpler!
// This is only a little faster.
//
// Part two is more straightforward, any time we encounter an A we see if it is
// the center of an X-MAS.
package day04

import (
	"advent/util"
	"strings"
)

const Word = "XMAS"

type Day04Solution struct {
	wordSearch util.Matrix[rune]
}

func NewDay04Solution(filepath string) (*Day04Solution, error) {
	wordSearch, err := util.ParseMatrixFromFile(filepath, func(r rune) rune {
		return r
	})
	return &Day04Solution{wordSearch}, err
}

func (s *Day04Solution) PartOneAnswer() (int, error) {
	return s.countWords(Word, s.wordSearch), nil
}

func (s *Day04Solution) PartTwoAnswer() (int, error) {
	return s.countXmases(s.wordSearch), nil
}

// countWords returns the number of times the word appears in the matrix.
// For correct behavior, word must not have any repeating substrings.
func (s *Day04Solution) countWords(word string, matrix util.Matrix[rune]) int {
	count := 0

	// let's start by making the dynamic programming table
	trackers := make(util.Matrix[Tracker], len(matrix))
	for i := 0; i < len(matrix); i++ {
		trackers[i] = make([]Tracker, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			p := util.NewVector(i, j)
			trackers.Set(p, NewTracker())
			forwardLetterIndex := strings.IndexRune(word, matrix.Get(p))
			backwardLetterIndex := len(word) - forwardLetterIndex - 1
			if forwardLetterIndex == 0 {
				trackers[i][j].forwardHorizontal = 0
				trackers[i][j].forwardVertical = 0
				trackers[i][j].forwardDiagonalLeft = 0
				trackers[i][j].forwardDiagonalRight = 0
			} else if forwardLetterIndex > 0 {
				// This letter is in word, so we need to update the trackers
				if trackers.PosInBounds(util.NewVector(i-1, j)) && trackers[i-1][j].forwardVertical+1 == forwardLetterIndex {
					trackers[i][j].forwardVertical = forwardLetterIndex
				}
				if trackers.PosInBounds(util.NewVector(i, j-1)) && trackers[i][j-1].forwardHorizontal+1 == forwardLetterIndex {
					trackers[i][j].forwardHorizontal = forwardLetterIndex
				}
				if trackers.PosInBounds(util.NewVector(i-1, j-1)) && trackers[i-1][j-1].forwardDiagonalRight+1 == forwardLetterIndex {
					trackers[i][j].forwardDiagonalRight = forwardLetterIndex
				}
				if trackers.PosInBounds(util.NewVector(i-1, j+1)) && trackers[i-1][j+1].forwardDiagonalLeft+1 == forwardLetterIndex {
					trackers[i][j].forwardDiagonalLeft = forwardLetterIndex
				}
			}
			if backwardLetterIndex == 0 {
				trackers[i][j].backwardHorizontal = 0
				trackers[i][j].backwardVertical = 0
				trackers[i][j].backwardDiagonalLeft = 0
				trackers[i][j].backwardDiagonalRight = 0
			} else if backwardLetterIndex > 0 {
				if backwardLetterIndex == 0 || (trackers.PosInBounds(util.NewVector(i-1, j)) && trackers[i-1][j].backwardVertical+1 == backwardLetterIndex) {
					trackers[i][j].backwardVertical = backwardLetterIndex
				}
				if backwardLetterIndex == 0 || (trackers.PosInBounds(util.NewVector(i, j-1)) && trackers[i][j-1].backwardHorizontal+1 == backwardLetterIndex) {
					trackers[i][j].backwardHorizontal = backwardLetterIndex
				}
				if backwardLetterIndex == 0 || (trackers.PosInBounds(util.NewVector(i-1, j-1)) && trackers[i-1][j-1].backwardDiagonalRight+1 == backwardLetterIndex) {
					trackers[i][j].backwardDiagonalRight = backwardLetterIndex
				}
				if backwardLetterIndex == 0 || (trackers.PosInBounds(util.NewVector(i-1, j+1)) && trackers[i-1][j+1].backwardDiagonalLeft+1 == backwardLetterIndex) {
					trackers[i][j].backwardDiagonalLeft = backwardLetterIndex
				}
			}
			count += trackers[i][j].CountFrequencies(len(word) - 1)
		}
	}
	return count
}

// countXmases returns the number of X-MASes in the matrix. As defined in the
// problem, an X-MAS is a 3x3 matrix with an A in the center and an M and an S
// in the diagonals. The M and S can be top to bottom or bottom to top.
func (s *Day04Solution) countXmases(matrix util.Matrix[rune]) int {
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if s.isXmasCenter(util.NewVector(i, j), matrix) {
				count++
			}
		}
	}
	return count
}

// isXmasCenter returns true if the given Vector is the center of an X-MAS.
func (s *Day04Solution) isXmasCenter(p *util.Vector, matrix util.Matrix[rune]) bool {
	if matrix.Get(p) != 'A' {
		return false
	}
	if p.X-1 < 0 || p.X+1 >= len(matrix) || p.Y-1 < 0 || p.Y+1 >= len(matrix[p.X]) {
		return false
	}
	return s.isLeftDiagonalXmasCenter(p, matrix) && s.isRightDiagonalXmasCenter(p, matrix)
}

// isLeftDiagonalXmasCenter returns true if the given Vector has an M and an S,
// in either order, at matrix[i-1][j-1] and matrix[i+1][j+1].
func (s *Day04Solution) isLeftDiagonalXmasCenter(p *util.Vector, matrix util.Matrix[rune]) bool {
	topRight := util.NewVector(p.X-1, p.Y-1)
	bottomLeft := util.NewVector(p.X+1, p.Y+1)
	return (matrix.Get(topRight) == 'M' && matrix.Get(bottomLeft) == 'S') ||
		(matrix.Get(topRight) == 'S' && matrix.Get(bottomLeft) == 'M')
}

// isRightDiagonalXmasCenter returns true if the given Vector has an M and an S,
// in either order, at matrix[i-1][j+1] and matrix[i+1][j-1].
func (s *Day04Solution) isRightDiagonalXmasCenter(p *util.Vector, matrix util.Matrix[rune]) bool {
	topLeft := util.NewVector(p.X-1, p.Y+1)
	bottomRight := util.NewVector(p.X+1, p.Y-1)
	return (matrix.Get(topLeft) == 'M' && matrix.Get(bottomRight) == 'S') ||
		(matrix.Get(topLeft) == 'S' && matrix.Get(bottomRight) == 'M')
}
