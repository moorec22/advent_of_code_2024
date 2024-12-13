package day04

// Tracker keeps track of the progress of a word in a given direction
type Tracker struct {
	// trackers for the word in the forward direction
	forwardHorizontal, forwardVertical, forwardDiagonalLeft, forwardDiagonalRight int
	// trackers for the word in the backward direction
	backwardHorizontal, backwardVertical, backwardDiagonalLeft, backwardDiagonalRight int
}

func NewTracker() Tracker {
	return Tracker{
		forwardHorizontal:     -1,
		forwardVertical:       -1,
		forwardDiagonalLeft:   -1,
		forwardDiagonalRight:  -1,
		backwardHorizontal:    -1,
		backwardVertical:      -1,
		backwardDiagonalLeft:  -1,
		backwardDiagonalRight: -1,
	}
}

// CountFrequencies returns the number of counters in t equal to i
func (t *Tracker) CountFrequencies(i int) int {
	count := 0
	if t.forwardHorizontal == i {
		count++
	}
	if t.forwardVertical == i {
		count++
	}
	if t.forwardDiagonalLeft == i {
		count++
	}
	if t.forwardDiagonalRight == i {
		count++
	}
	if t.backwardHorizontal == i {
		count++
	}
	if t.backwardVertical == i {
		count++
	}
	if t.backwardDiagonalLeft == i {
		count++
	}
	if t.backwardDiagonalRight == i {
		count++
	}
	return count
}
