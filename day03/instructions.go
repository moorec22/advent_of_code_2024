package day03

// Instruction is a single instruction in the "programming language" defined by the problem
type Instruction interface {
	Execute(*ProgramState) *ProgramState
}

type StartInstruction struct{}

func NewStartInstruction() *StartInstruction {
	return &StartInstruction{}
}

func (i *StartInstruction) Execute(state *ProgramState) *ProgramState {
	return state
}

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

type DoInstruction struct{}

func NewDoInstruction() *DoInstruction {
	return &DoInstruction{}
}

func (i *DoInstruction) Execute(state *ProgramState) *ProgramState {
	state.doCompute = true
	return state
}

type DontInstruction struct{}

func NewDontInstruction() *DontInstruction {
	return &DontInstruction{}
}

func (i *DontInstruction) Execute(state *ProgramState) *ProgramState {
	state.doCompute = false
	return state
}
