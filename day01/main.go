// Advent of Code, 2024, Day 1
//
// https://adventofcode.com/2024/day/1
package main

import (
	"bufio"
	"os"
)

const Filepath = "files/test.txt"

func main() {
}

func processFile(filepath string, process func(*bufio.Scanner) error) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := getFileScanner(file)

	return process(scanner)
}

func getFileScanner(file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	return scanner
}
