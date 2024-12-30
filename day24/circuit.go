package day24

import (
	"advent/util"
	"fmt"
)

type Circuit struct {
	// initial state for each var. There is no initial state if the var is not in
	// initialStates
	initialStates map[string]bool
	// inputs maps an output var to its two input vars
	gates map[string]*Gate
	// stores the current state of each variable. Equal to initialStates
	// initially and when reset.
	states map[string]bool
}

func NewCircuit(initialStates map[string]bool, gates map[string]*Gate) *Circuit {
	states := util.CopyMap(initialStates)
	return &Circuit{initialStates, gates, states}
}

// Solve returns the value of the variable designated by v. Returns an error if
// any variable needed to solve v is not found.
func (c *Circuit) Solve(v string) (bool, error) {
	val, ok := c.states[v]
	if ok {
		return val, nil
	}
	gate, ok := c.gates[v]
	if !ok {
		return false, fmt.Errorf("no state or gate found for variable: %s", v)
	}
	leftVal, err := c.Solve(gate.Left)
	if err != nil {
		return false, err
	}
	rightVal, err := c.Solve(gate.Right)
	if err != nil {
		return false, err
	}
	val = gate.Compute(leftVal, rightVal)
	c.states[v] = val
	return val, nil
}

func (c *Circuit) Reset() {
	c.states = util.CopyMap(c.initialStates)
}
