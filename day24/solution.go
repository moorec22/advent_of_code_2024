// Advent of Code, 2024, Day 24
//
// https://adventofcode.com/2024/day/24
//
// Part 1: The biggest challenge initially I think is how to store these gates.
// My idea is to have a gate interface, for all types of gates. These gates do
// not store state. Then, we'll have a circuit class, that stores all gates and
// caches state for each gate. The circuit can compute the final output of all
// z gates.
//
// Part 2: We need to do four gate swaps for our solution. If we were to try
// every combo, we would have a runtime of O(n^8), which may be okay for a
// small problem. But perhaps there's something better. We know there are no
// loops, and it is trying to add two numbers together. We see that the digits
// are cummulative: z01 depends only on xs and ys below it (not x02 for example).
// So I decided to start by finding the first digit that is incorrect.
//
// This ended up being a lot of trial and error, and notebooking! I didn't end
// up with a software solution, and just fixed the gates as I found them. So
// it's messy!
package day24

import (
	"advent/util"
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Pair struct {
	one, two string
}

var BitMap = map[bool]int{
	false: 0,
	true:  1,
}

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
			outputVar := strings.TrimSpace(parts[1])
			gates[outputVar] = NewGate(gateParts[0], gateParts[2], gateParts[1])
			variables[outputVar] = true
		}
		return scanner.Err()
	})
	if err != nil {
		return nil, err
	}
	// this is a way to replace the current names in the input with readabale
	// names for intuitively understanding the machine
	// readableNames, oldNames := getReadableNames(variables, gates)
	// readableVariables := getReadableNamesSet(variables, readableNames)
	// readableInitialStates := getReadableNamesSet(initialStates, readableNames)
	// readableGates := getReadableGates(gates, readableNames)

	circuit := NewCircuit(initialStates, gates)
	return &Day24Solution{circuit, variables}, nil
}

func (s *Day24Solution) PartOneAnswer() (int, error) {
	return s.getAnswer(s.circuit, "z")
}

// 27133763813069 + 20654586710041 = 47788350523110
// 47788350523110 = 1010110111011010010111101010000100011011100110
//
// we know z11 needs to be switched: it takes x11 and y11 as input,
// rather than using the carry over digit as well.
//
// we can see that z24 is using its own carry over digit, and not
// c23. If we can find where c23 went, we can find another swap.
//
// Discovered Swaps:
//
//   - vkq and z11
//   - mmk and z24
//   - ftq and z28
//   - hqh and z38
func (s *Day24Solution) PartTwoAnswer() (int, error) {
	s.circuit.Reset()
	rightAnswer := "1010110111011010010111101010000100011011100110"
	swaps := []string{"vkq", "z11", "mmk", "z24", "pvb", "qdq", "hqh", "z38"}
	slices.Sort(swaps)
	s.swapGateOutputs(s.circuit, "vkq", "z11")
	s.swapGateOutputs(s.circuit, "mmk", "z24")
	s.swapGateOutputs(s.circuit, "pvb", "qdq")
	s.swapGateOutputs(s.circuit, "hqh", "z38")
	binaryStringAnswer, err := s.getBinaryStringAnswer(s.circuit, "z")
	if err != nil {
		return 0, err
	}
	fmt.Println(rightAnswer)
	fmt.Println(binaryStringAnswer)
	for i := len(rightAnswer) - 1; i >= 0; i-- {
		if rightAnswer[i] != binaryStringAnswer[i] {
			fmt.Println("wrong digit:", len(rightAnswer)-1-i)
		}
	}
	fmt.Println("Solution:", strings.Join(swaps, ","))
	return 0, nil
}

// getAnswer takes the circuit and returns the integer value of all digits that
// start with the prefix when treated as a binary number.
func (s *Day24Solution) getAnswer(circuit *Circuit, prefix string) (int, error) {
	vars := s.getVariablesWithPrefix(s.variables, prefix)
	num := 0
	for _, v := range vars {
		val, err := circuit.Solve(v)
		if err != nil {
			return 0, err
		}
		num = 2*num + BitMap[val]
	}
	return num, nil
}

func (s *Day24Solution) getBinaryStringAnswer(circuit *Circuit, prefix string) (string, error) {
	vars := s.getVariablesWithPrefix(s.variables, prefix)
	num := ""
	for _, v := range vars {
		val, err := circuit.Solve(v)
		if err != nil {
			return "", err
		}
		if val {
			num += "1"
		} else {
			num += "0"
		}
	}
	return num, nil
}

// getVariablesWithPrefix returns all variables that have the given prefix, in
// reverse alphabetical order.
func (s *Day24Solution) getVariablesWithPrefix(variables map[string]bool, prefix string) []string {
	variableList := make([]string, 0)
	for variable := range variables {
		if strings.HasPrefix(variable, prefix) {
			variableList = append(variableList, variable)
		}
	}
	slices.Sort(variableList)
	slices.Reverse(variableList)
	return variableList
}

func (s *Day24Solution) swapGateOutputs(circuit *Circuit, vOne, vTwo string) {
	gOne := circuit.gates[vOne]
	gTwo := circuit.gates[vTwo]
	circuit.gates[vOne] = gTwo
	circuit.gates[vTwo] = gOne
}

func (s *Day24Solution) printCircuitPath(circuit *Circuit, variable string) {
	variables := make([]string, 0)
	gate := circuit.gates[variable]
	if gate != nil {
		variables = append(variables, gate.Left, gate.Right)
		s.printCircuitPathHelper(circuit, variables, 1)
	}
	fmt.Println(variable)
}

func (s *Day24Solution) printCircuitPathHelper(circuit *Circuit, variables []string, level int) {
	newVariables := make([]string, 0)
	for _, variable := range variables {
		gate := circuit.gates[variable]
		if gate != nil {
			newVariables = append(newVariables, gate.Left, gate.Right)
		}
	}
	if len(newVariables) != 0 {
		s.printCircuitPathHelper(circuit, newVariables, level+1)
	}
	fmt.Printf("Level %d: %v\n", level, variables)
}

func (s *Day24Solution) printTwoCircuitPaths(circuit *Circuit, vOne, vTwo string) {
	vOnes := s.getNextLevel(circuit, []string{vOne})
	vTwos := s.getNextLevel(circuit, []string{vTwo})
	s.printTwoCircuitPathsHelper(circuit, vOnes, vTwos, 1)
	fmt.Println(vOne, vTwo)
}

func (s *Day24Solution) getNextLevel(circuit *Circuit, variables []string) []string {
	newVariables := make([]string, 0)
	for _, variable := range variables {
		gate := circuit.gates[variable]
		if gate != nil {
			newVariables = append(newVariables, gate.Left, gate.Right)
		}
	}
	return newVariables
}

func (s *Day24Solution) printTwoCircuitPathsHelper(circuit *Circuit, vOne, vTwo []string, level int) {
	newVOne := make([]string, 0)
	for _, variable := range vOne {
		gate := circuit.gates[variable]
		if gate != nil {
			newVOne = append(newVOne, gate.Left, gate.Right)
		}
	}
	newVTwo := make([]string, 0)
	for _, variable := range vTwo {
		gate := circuit.gates[variable]
		if gate != nil {
			newVTwo = append(newVTwo, gate.Left, gate.Right)
		}
	}
	if len(newVOne) != 0 || len(newVTwo) != 0 {
		s.printTwoCircuitPathsHelper(circuit, newVOne, newVTwo, level+1)
	}
	fmt.Printf("Level %d: %v      %v\n", level, vOne, vTwo)
}
