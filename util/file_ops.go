package util

import (
	"bufio"
	"os"
)

// processFile safely opens a file designated by filepath, passes a scanner of
// that file to process, and closes that file. Returns an error if there is an
// error opening the file, or if process returns an error.
func ProcessFile(filepath string, process func(*bufio.Scanner) error) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	return process(scanner)
}

// ParseMatrixFromFile will parse a file of 2d runes into a matrix of any type.
// The function stops parsing a matrix as soon as a blank line is reached. The
// function will return an error if there is an error opening the file, or if
// there is an error parsing the file.
func ParseMatrixFromFile[T any](filepath string, toT func(rune) T) (Matrix[T], error) {
	matrix := NewMatrix[T]()
	err := ProcessFile(filepath, func(scanner *bufio.Scanner) error {
		var err error
		matrix, err = ParseMatrixFromScanner(scanner, toT)
		return err
	})
	return matrix, err
}

// ParseMatrixFromScanner will parse a file of 2d runes into a matrix of any
// type. The function stops parsing a matrix as soon as a blank line is reached.
// The function will return an error if there is there is an error parsing from
// the scanner.
func ParseMatrixFromScanner[T any](scanner *bufio.Scanner, toT func(rune) T) (Matrix[T], error) {
	matrix := NewMatrix[T]()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return matrix, nil
		}
		row := make([]T, len(line))
		for i, r := range line {
			row[i] = toT(r)
		}
		matrix = append(matrix, row)
	}
	return matrix, scanner.Err()
}
