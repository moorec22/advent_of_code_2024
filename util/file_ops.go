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

func ParseMatrix(filepath string) (Matrix[rune], error) {
	matrix := NewMatrix[rune]()
	err := ProcessFile(filepath, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			line := scanner.Text()
			row := make([]rune, len(line))
			for i, r := range line {
				row[i] = r
			}
			matrix = append(matrix, row)
		}
		return scanner.Err()
	})
	return matrix, err
}
