// Advent of Code, 2024, Day 17
//
// https://adventofcode.com/2024/day/17
package day17

import (
	"advent/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const RegisterPrefix = "Register X:"
const ProgramPrefix = "Program:"

type Day17Solution struct {
	program *Program
}

func NewDay17Solution(filename string) (*Day17Solution, error) {
	registers := Registers{}
	program := make([]int, 0)
	err := util.ProcessFile(filename, func(scanner *bufio.Scanner) error {
		var err error
		registers := []*int{&registers.A, &registers.B, &registers.C}
		for i := 0; i < RegisterCount; i++ {
			scanner.Scan()
			line := scanner.Text()
			val, err := parseRegister(line)
			if err != nil {
				return err
			}
			*registers[i] = val
		}
		scanner.Scan()
		scanner.Scan()
		line := scanner.Text()
		program, err = parseProgram(line)
		return err
	})
	return &Day17Solution{NewProgram(program, NewProgramState(registers))}, err
}

func (s *Day17Solution) PartOneAnswer() (int, error) {
	output, err := s.program.Run()
	if err != nil {
		return 0, err
	}
	fmt.Println("Registers:", s.program.State.Registers)
	fmt.Println("Output:", s.arrToString(output))
	return 0, nil
}

func (s *Day17Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}

func (s *Day17Solution) arrToString(arr []int) string {
	if len(arr) == 0 {
		return ""
	}
	str := strconv.Itoa(arr[0])
	for i := 1; i < len(arr); i++ {
		str += "," + strconv.Itoa(arr[i])
	}
	return str
}

// parseRegister takes an input in the form "Register X: <val>" and returns the
// value.
func parseRegister(input string) (int, error) {
	valStr := input[len(RegisterPrefix):]
	valStr = strings.TrimSpace(valStr)
	return strconv.Atoi(valStr)
}

// parseProgram takes an input in the form "Program: <prg>" and returns the
// program as a slice of integers.
func parseProgram(input string) ([]int, error) {
	prgStr := strings.TrimPrefix(input, ProgramPrefix)
	prgStr = strings.TrimSpace(prgStr)
	prgStrs := strings.Split(prgStr, ",")
	prg := make([]int, len(prgStrs))
	for i, prgStr := range prgStrs {
		val, err := strconv.Atoi(prgStr)
		if err != nil {
			return nil, err
		}
		prg[i] = val
	}
	return prg, nil
}
