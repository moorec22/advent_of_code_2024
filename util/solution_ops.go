package util

// Solution is an interface to be used by a solution for any given day.
type Solution interface {
	PartOneAnswer() (int, error)
	PartTwoAnswer() (int, error)
}
