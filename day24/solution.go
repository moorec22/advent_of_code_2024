// Advent of Code, 2024, Day 24
//
// https://adventofcode.com/2024/day/24
//
// Part 1: The biggest challenge initially I think is how to store these gates.
// My idea is to have a gate interface, for all types of gates. These gates do
// not store state. Then, we'll have a circuit class, that stores all gates and
// caches state for each gate. The circuit can compute the final output of all
// z gates.
package day24

import (
	"advent/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Day24Solution struct {
	circuit   *Circuit
	variables map[string]bool
}

func NewDay24Solution(filename string) (*Day24Solution, error) {
	initialStates := make(map[string]bool)
	gates := make(map[string]*Gate)
	variables := make(map[string]bool)
	err := util.ProcessFile(filename, func(scanner *bufio.Scanner) error {
		// starting with inputs
		for scanner.Scan() {
			if scanner.Text() == "" {
				break
			}
			parts := strings.Split(scanner.Text(), ": ")
			state, err := strconv.Atoi(parts[1])
			if err != nil || (state != 0 && state != 1) {
				return fmt.Errorf("can't parse into boolean: '%s'", parts[1])
			}
			initialStates[parts[0]] = state == 1
			variables[parts[0]] = true
		}
		// and get gates
		for scanner.Scan() {
			parts := strings.Split(scanner.Text(), "->")
			gateParts := strings.Split(strings.TrimSpace(parts[0]), " ")
			gateFunc, ok := GateFunctions[gateParts[1]]
			if !ok {
				return fmt.Errorf("could not find gate function for %s", gateParts[1])
			}
			outputVar := strings.TrimSpace(parts[1])
			gates[outputVar] = NewGate(gateParts[0], gateParts[2], gateFunc)
			variables[outputVar] = true
		}
		return scanner.Err()
	})
	if err != nil {
		return nil, err
	}
	circuit := NewCircuit(initialStates, gates)
	return &Day24Solution{circuit, variables}, nil
}

func (s *Day24Solution) PartOneAnswer() (int, error) {
	return 0, nil
}

func (s *Day24Solution) PartTwoAnswer() (int, error) {
	return 0, nil
}
