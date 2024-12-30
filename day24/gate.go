package day24

type Gate struct {
	// variables that are input into a gate
	Left, Right, gateStr string
	Compute              GateFunc
}

func NewGate(left, right, gateStr string) *Gate {
	return &Gate{left, right, gateStr, GateFunctions[gateStr]}
}

// a simple, 2 input gate interface
type GateFunc func(x, y bool) bool

// AndGate returns true iff x && y
func AndGate(x, y bool) bool {
	return x && y
}

// OrGate returns true iff x || y
func OrGate(x, y bool) bool {
	return x || y
}

// XorGate returns true iff x != y
func XorGate(x, y bool) bool {
	return x != y
}

var GateFunctions = map[string]GateFunc{
	"AND": AndGate,
	"OR":  OrGate,
	"XOR": XorGate,
}
