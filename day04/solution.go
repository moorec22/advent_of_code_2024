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
	"bufio"
	"strings"
)

const Word = "XMAS"

type Position struct {
	i, j int
}

// Tracker keeps track of the progress of a word in a given direction
type Tracker struct {
	// trackers for the word in the forward direction
	forwardHorizontal, forwardVertical, forwardDiagonalLeft, forwardDiagonalRight int
	// trackers for the word in the backward direction
	backwardHorizontal, backwardVertical, backwardDiagonalLeft, backwardDiagonalRight int
}

func NewTracker() Tracker {
	return Tracker{
		forwardHorizontal:     -1,
		forwardVertical:       -1,
		forwardDiagonalLeft:   -1,
		forwardDiagonalRight:  -1,
		backwardHorizontal:    -1,
		backwardVertical:      -1,
		backwardDiagonalLeft:  -1,
		backwardDiagonalRight: -1,
	}
}

// CountFrequencies returns the number of counters in t equal to i
func (t *Tracker) CountFrequencies(i int) int {
	count := 0
	if t.forwardHorizontal == i {
		count++
	}
	if t.forwardVertical == i {
		count++
	}
	if t.forwardDiagonalLeft == i {
		count++
	}
	if t.forwardDiagonalRight == i {
		count++
	}
	if t.backwardHorizontal == i {
		count++
	}
	if t.backwardVertical == i {
		count++
	}
	if t.backwardDiagonalLeft == i {
		count++
	}
	if t.backwardDiagonalRight == i {
		count++
	}
	return count
}

func PartOneAnswer(filepath string) (int, error) {
	matrix := getMatrix(filepath)
	return countWords(Word, matrix), nil
}

func PartTwoAnswer(filepath string) (int, error) {
	matrix := getMatrix(filepath)
	return countXmases(matrix), nil
}

// getMatrix returns a matrix of runes from a file. Each row is a line in the
// file.
func getMatrix(filepath string) [][]rune {
	matrix := make([][]rune, 0)
	util.ProcessFile(filepath, func(s *bufio.Scanner) error {
		for s.Scan() {
			line := s.Text()
			row := make([]rune, 0)
			for _, r := range line {
				row = append(row, r)
			}
			matrix = append(matrix, row)
		}
		return nil
	})
	return matrix
}

// countWords returns the number of times the word appears in the matrix.
// For correct behavior, word must not have any repeating substrings.
func countWords(word string, matrix [][]rune) int {
	count := 0

	// let's start by making the dynamic programming table
	trackers := make([][]Tracker, len(matrix))
	for i := 0; i < len(matrix); i++ {
		trackers[i] = make([]Tracker, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			trackers[i][j] = NewTracker()
			forwardLetterIndex := strings.IndexRune(word, matrix[i][j])
			backwardLetterIndex := len(word) - forwardLetterIndex - 1
			if forwardLetterIndex == 0 {
				trackers[i][j].forwardHorizontal = 0
				trackers[i][j].forwardVertical = 0
				trackers[i][j].forwardDiagonalLeft = 0
				trackers[i][j].forwardDiagonalRight = 0
			} else if forwardLetterIndex > 0 {
				// This letter is in word, so we need to update the trackers
				if (positionInBounds(Position{i - 1, j}, trackers) && trackers[i-1][j].forwardVertical+1 == forwardLetterIndex) {
					trackers[i][j].forwardVertical = forwardLetterIndex
				}
				if (positionInBounds(Position{i, j - 1}, trackers) && trackers[i][j-1].forwardHorizontal+1 == forwardLetterIndex) {
					trackers[i][j].forwardHorizontal = forwardLetterIndex
				}
				if (positionInBounds(Position{i - 1, j - 1}, trackers) && trackers[i-1][j-1].forwardDiagonalRight+1 == forwardLetterIndex) {
					trackers[i][j].forwardDiagonalRight = forwardLetterIndex
				}
				if (positionInBounds(Position{i - 1, j + 1}, trackers) && trackers[i-1][j+1].forwardDiagonalLeft+1 == forwardLetterIndex) {
					trackers[i][j].forwardDiagonalLeft = forwardLetterIndex
				}
			}
			if backwardLetterIndex == 0 {
				trackers[i][j].backwardHorizontal = 0
				trackers[i][j].backwardVertical = 0
				trackers[i][j].backwardDiagonalLeft = 0
				trackers[i][j].backwardDiagonalRight = 0
			} else if backwardLetterIndex > 0 {
				if backwardLetterIndex == 0 || (positionInBounds(Position{i - 1, j}, trackers) && trackers[i-1][j].backwardVertical+1 == backwardLetterIndex) {
					trackers[i][j].backwardVertical = backwardLetterIndex
				}
				if backwardLetterIndex == 0 || (positionInBounds(Position{i, j - 1}, trackers) && trackers[i][j-1].backwardHorizontal+1 == backwardLetterIndex) {
					trackers[i][j].backwardHorizontal = backwardLetterIndex
				}
				if backwardLetterIndex == 0 || (positionInBounds(Position{i - 1, j - 1}, trackers) && trackers[i-1][j-1].backwardDiagonalRight+1 == backwardLetterIndex) {
					trackers[i][j].backwardDiagonalRight = backwardLetterIndex
				}
				if backwardLetterIndex == 0 || (positionInBounds(Position{i - 1, j + 1}, trackers) && trackers[i-1][j+1].backwardDiagonalLeft+1 == backwardLetterIndex) {
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
func countXmases(matrix [][]rune) int {
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if isXmasCenter(i, j, matrix) {
				count++
			}
		}
	}
	return count
}

// isXmasCenter returns true if the given position is the center of an X-MAS.
func isXmasCenter(i, j int, matrix [][]rune) bool {
	if matrix[i][j] != 'A' {
		return false
	}
	if i-1 < 0 || i+1 >= len(matrix) || j-1 < 0 || j+1 >= len(matrix[i]) {
		return false
	}
	return isLeftDiagonalXmasCenter(i, j, matrix) && isRightDiagonalXmasCenter(i, j, matrix)
}

// isLeftDiagonalXmasCenter returns true if the given position has an M and an S,
// in either order, at matrix[i-1][j-1] and matrix[i+1][j+1].
func isLeftDiagonalXmasCenter(i, j int, matrix [][]rune) bool {
	return (matrix[i-1][j-1] == 'M' && matrix[i+1][j+1] == 'S') ||
		(matrix[i-1][j-1] == 'S' && matrix[i+1][j+1] == 'M')
}

// isRightDiagonalXmasCenter returns true if the given position has an M and an S,
// in either order, at matrix[i-1][j+1] and matrix[i+1][j-1].
func isRightDiagonalXmasCenter(i, j int, matrix [][]rune) bool {
	return (matrix[i-1][j+1] == 'M' && matrix[i+1][j-1] == 'S') ||
		(matrix[i-1][j+1] == 'S' && matrix[i+1][j-1] == 'M')
}

// positionInBounds returns true if the given position is within the bounds of
// the matrix.
func positionInBounds(p Position, trackers [][]Tracker) bool {
	return p.i >= 0 && p.i < len(trackers) && p.j >= 0 && p.j < len(trackers[p.i])
}
