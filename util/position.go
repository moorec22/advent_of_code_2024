package util

type Position struct {
	Row, Col int
}

func NewPosition(row, col int) Position {
	return Position{Row: row, Col: col}
}

func (p *Position) Add(other Position) Position {
	return Position{Row: p.Row + other.Row, Col: p.Col + other.Col}
}

func (p *Position) GetManhattanDistance(other Position) ManhattanDistance {
	return NewManhattanDistance(p.Row-other.Row, p.Col-other.Col)
}

func (p *Position) AddManhattanDistance(distance ManhattanDistance) Position {
	return Position{Row: p.Row + distance.X, Col: p.Col + distance.Y}
}
