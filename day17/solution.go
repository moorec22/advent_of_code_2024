// Advent of Code, 2024, Day 17
//
// https://adventofcode.com/2024/day/17
//
// Part 1: I wrote an interpreter reflected by the problem statement. To find
// the answer, the input program is run and output.
//
// Part 2: On the outset, this seemed a bit trickier. I decided to solve just the
// test case given to us, which as a program does the following:
//   - A = A / 8
//   - print A % 8
//   - if A is not 0, repeat
//
// We can work backwards. We just want A % 8 to equal the instruction. For example,
// the last instruction is 0, so let's take the smallest value such that A % 8 = 0.
// Which is 0.
//
// Thinking further, we see that every program is either a one-shot, or a single
// loop that repeats until A is 0. We also see that there are two kinds of operations:
//   - division: TODO: reversible?
//   - bitwise XOR: reversible. x ^ y ^ y = x
//
// We know this is not a solvable problem in the general case. So, let's focus
// on JUST the input case. Some notes:
//   - the program is only dependent on A: B and C are written at the beginning
//     of each loop.
//   - the program only uses the first 10 bits of A.
//   - the program only uses the last 3-10 bits of A.
//   - A is shifted right by 3 each time.
//
// Idea for a recursive program to solve:
//   - find a 10 bit integer that outputs the last number.
//   - shift that number up by 3 bits.
//   - alter the bottom 3 bits to output the second to last number.
//   - repeat until we have the full program.
//
// to simplify, we can try for all 7-bit (0-127) numbers and recurse from
// the very bottom of the program.
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
	output, err := s.program.Copy().Run(true)
	if err != nil {
		return 0, err
	}
	fmt.Println("Registers:", s.program.State.Registers)
	fmt.Println("Output:", s.arrToString(output))
	return 0, nil
}

func (s *Day17Solution) PartTwoAnswer() (int, error) {
	for i := 0; i < 128; i++ {
		a, found := s.findSelfPrintingProgram(s.program, len(s.program.Instructions)-1, i)
		if found {
			return a, nil
		}
	}
	return 0, fmt.Errorf("no self-printing program found")
}

// findSelfPrintingProgram finds a program with an initial value for register A
// such that the program prints itself.
//
// NOTE: This only functions for programs that:
// - only depend on A as input for the loop
// - only use the last 3-10 bits of A for each loop
func (s *Day17Solution) findSelfPrintingProgram(p *Program, pos, a int) (int, bool) {
	if pos < 0 {
		return a, true
	}
	a = a << 3
	for i := 0; i < 8; i++ {
		p.State.InstructionPointer = 0
		p.State.Registers.A = a | i
		output, _ := p.Run(false)
		if len(output) == 1 && output[0] == p.Instructions[pos] {
			a, found := s.findSelfPrintingProgram(p, pos-1, a|i)
			if found {
				return a, true
			}
		}
	}
	return 0, false
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
