package interpreter

// ProgramState is a struct that holds the current state of the program
type ProgramState struct {
	doCompute bool
	Answer    int
}

func NewProgramState() *ProgramState {
	return &ProgramState{true, 0}
}
