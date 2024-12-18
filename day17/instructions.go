package day17

import "advent/util"

type Instruction interface {
	// Execute runs the instruction on the given state. It returns an error if
	// the instruction fails. It returns any output generated by the instruction.
	// If there is no output, an empty slice is returned.
	Execute(operand int, state *ProgramState) ([]int, error)
}

// AdvInstruction is a division operator. It raises the value at register A to
// 2^(combo operand) and stores the result in register A. It increments
// instruction pointer by 2.
type AdvInstruction struct{}

func (i *AdvInstruction) Execute(operand int, state *ProgramState) ([]int, error) {
	state.Registers.A = getDvValue(operand, state)
	state.InstructionPointer += 2
	return []int{}, nil
}

// BxlInstructions calculates the bitwise XOR of register B and the literal
// operand, and writes the result to register B. It increments instruction
// pointer by 2.
type BxlInstruction struct{}

func (i *BxlInstruction) Execute(operand int, state *ProgramState) ([]int, error) {
	state.Registers.B ^= operand
	state.InstructionPointer += 2
	return []int{}, nil
}

// BstInstruction mods its combo operator by 8 and stores the result in
// register B. It increments instruction pointer by 2.
type BstInstruction struct{}

func (i *BstInstruction) Execute(operand int, state *ProgramState) ([]int, error) {
	state.Registers.B = state.GetComboOperand(operand) % 8
	state.InstructionPointer += 2
	return []int{}, nil
}

// JnzInstruction jumps to the instruction at the offset of the literal operand
// if the value of register A is not zero. Otherwise, it increments the
// instruction pointer by 2.
type JnzInstruction struct{}

func (i *JnzInstruction) Execute(operand int, state *ProgramState) ([]int, error) {
	if state.Registers.A != 0 {
		state.InstructionPointer = operand
	} else {
		state.InstructionPointer += 2
	}
	return []int{}, nil
}

// BxcInstruction takes the bitwise XOR of register B and register C, and stores
// the result in register B. It increments the instruction pointer by 2.
type BxcInstruction struct{}

func (i *BxcInstruction) Execute(operand int, state *ProgramState) ([]int, error) {
	state.Registers.B ^= state.Registers.C
	state.InstructionPointer += 2
	return []int{}, nil
}

// OutInstruction returns the combo operand, modded by 8. It
// increments the instruction pointer by 2.
type OutInstruction struct{}

func (i *OutInstruction) Execute(operand int, state *ProgramState) ([]int, error) {
	state.InstructionPointer += 2
	return []int{state.GetComboOperand(operand) % 8}, nil
}

// BdvInstruction is a division operator. It raises the value at register A to
// 2^(combo operand) and stores the result in register B. It increments
// instruction pointer by 2.
type BdvInstruction struct{}

func (i *BdvInstruction) Execute(operand int, state *ProgramState) ([]int, error) {
	state.Registers.B = getDvValue(operand, state)
	state.InstructionPointer += 2
	return []int{}, nil
}

// CdvInstruction is a division operator. It raises the value at register A to
// 2^(combo operand) and stores the result in register C. It increments
// instruction pointer by 2.
type CdvInstruction struct{}

func (i *CdvInstruction) Execute(operand int, state *ProgramState) ([]int, error) {
	state.Registers.C = getDvValue(operand, state)
	state.InstructionPointer += 2
	return []int{}, nil
}

func getDvValue(operand int, state *ProgramState) int {
	return state.Registers.A / util.IntPow(2, state.GetComboOperand(operand))
}
