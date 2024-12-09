package interpreter

// ProgramState is a struct that holds the current state of the program
type ProgramState struct {
	// doCompute is a flag that determines whether the program should run multiplication operations
	doCompute bool
	// Answer is the current answer to the program, the sum of all multiplication operations
	Answer int
}

func NewProgramState() *ProgramState {
	return &ProgramState{true, 0}
}
