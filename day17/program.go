package day17

import "fmt"

// instructionMap is a map of instruction IDs to their corresponding
// instructions.
var instructionMap = map[int]Instruction{
	0: &AdvInstruction{},
	1: &BxlInstruction{},
	2: &BstInstruction{},
	3: &JnzInstruction{},
	4: &BxcInstruction{},
	5: &OutInstruction{},
	6: &BdvInstruction{},
	7: &CdvInstruction{},
}

// Program is a program that can be executed. It contains state, which may be
// modified during execution. It contains instructions, which will not be
// modified.
type Program struct {
	Instructions []int
	State        *ProgramState
}

func NewProgram(instructions []int, state *ProgramState) *Program {
	return &Program{Instructions: instructions, State: state}
}

func (p *Program) Copy() *Program {
	newProgram := NewProgram(p.Instructions, p.State.Copy())
	return newProgram
}

// Run executes the program and returns the output. If loop is false, the
// program will terminate when a JNZ is encountered. Otherwise, the program
// is run.
func (p *Program) Run(loop bool) ([]int, error) {
	output := make([]int, 0)
	for p.State.InstructionPointer < len(p.Instructions) {
		if !loop && p.Instructions[p.State.InstructionPointer] == 3 {
			break
		}
		res, err := p.RunSingle()
		if err != nil {
			return nil, err
		}
		output = append(output, res...)
	}
	return output, nil
}

func (p *Program) RunSingle() ([]int, error) {
	instruction, ok := instructionMap[p.Instructions[p.State.InstructionPointer]]
	if !ok {
		return nil, fmt.Errorf("invalid instruction ID: %d", p.Instructions[p.State.InstructionPointer])
	}
	operand := p.Instructions[p.State.InstructionPointer+1]
	return instruction.Execute(operand, p.State)
}
