package day17

// ProgramState stores the needed state for a program: registers and the current
// instruction.
type ProgramState struct {
	Registers          Registers
	InstructionPointer int
}

func NewProgramState(registers Registers) *ProgramState {
	return &ProgramState{
		Registers:          registers,
		InstructionPointer: 0,
	}
}

func (p *ProgramState) Copy() *ProgramState {
	newState := NewProgramState(p.Registers)
	newState.InstructionPointer = p.InstructionPointer
	return newState
}

// Combo operands are defined in the problem as follows:
//   - Combo operands 0 through 3 represent literal values 0 through 3.
//   - Combo operand 4 represents the value of register A.
//   - Combo operand 5 represents the value of register B.
//   - Combo operand 6 represents the value of register C.
//   - Combo operand 7 is reserved and will not appear in valid programs.
func (s *ProgramState) GetComboOperand(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return s.Registers.A
	case 5:
		return s.Registers.B
	case 6:
		return s.Registers.C
	default:
		return 0
	}
}

// RegisterCount is the number of registers expected.
const RegisterCount = 3

// Registers stores three registers for the program
type Registers struct {
	A, B, C int
}

func NewRegisters() *Registers {
	return &Registers{}
}
