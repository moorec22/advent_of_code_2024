package interpreter

// Instruction is a single instruction in the "programming language" defined by the problem
type Instruction interface {
	// Execute with execute the instruction on the given ProgramState, and return the new ProgramState.
	Execute(*ProgramState) *ProgramState
}

// EmptyInstruction is an instruction that does nothing.
type EmptyInstruction struct{}

func NewEmptyInstruction() *EmptyInstruction {
	return &EmptyInstruction{}
}

func (i *EmptyInstruction) Execute(state *ProgramState) *ProgramState {
	return state
}

// MultiplyInstruction is an instruction that multiplies two numbers together.
type MultiplyInstruction struct {
	left  int
	right int
}

func NewMultiplyInstruction(left int, right int) *MultiplyInstruction {
	return &MultiplyInstruction{left, right}
}

func (i *MultiplyInstruction) Execute(state *ProgramState) *ProgramState {
	if state.doCompute {
		state.Answer += i.left * i.right
	}
	return state
}

// DoInstruction is an instruction that enables operations in the program state.
type DoInstruction struct{}

func NewDoInstruction() *DoInstruction {
	return &DoInstruction{}
}

func (i *DoInstruction) Execute(state *ProgramState) *ProgramState {
	state.doCompute = true
	return state
}

// DontInstruction is an instruction that disables operations in the program state.
type DontInstruction struct{}

func NewDontInstruction() *DontInstruction {
	return &DontInstruction{}
}

func (i *DontInstruction) Execute(state *ProgramState) *ProgramState {
	state.doCompute = false
	return state
}
