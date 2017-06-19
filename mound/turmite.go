package mound

import colorful "github.com/lucasb-eyer/go-colorful"

// Direction determines the next location of the turmite
type Direction int

// [N]orth [E]ast [S]outh [W]est
// Clockwise movement starting at N (noon)
const (
	North Direction = iota
	East
	South
	West
)

// Turn is relative to the direction of the turmite
type Turn int

// [L]eft [R]igth [U]-turn [N]o-turn
const (
	R Turn = iota
	L
	U
	N
)

// Move defines the color to replace under the turmite and the direction the turmite should move
type Move struct {
	C colorful.Color
	T Turn
}

// Rule is a single direction in the form of a dictionary
type Rule map[colorful.Color]Move

// Turmite is an individual that exists inside a grid
// X and Y are positions of size defined in a the Block of a Grid
type Turmite struct {
	Direction Direction
	Location  int
	Rule      Rule
}

// CreateRules returns a new Rule structure
func CreateRules() *Rule {
	m := Rule(make(map[colorful.Color]Move))
	return &m
}

// AddRule takes a color as a key and creates a move out of a second color and a turn and is added to the map
func (r Rule) AddRule(c1 colorful.Color, c2 colorful.Color, t Turn) {
	m := Move{
		C: c2,
		T: t,
	}
	r[c1] = m
}

// findMove returns the move associated with a color if it exists
func (t Turmite) findMove(c colorful.Color) (Move, bool) {
	m, ok := t.Rule[c]
	return m, ok
}

// Move takes a turmite, a turn, and a mound and returns the new direction and the update position of the turmite
func (t Turmite) move(turn Turn, grid Grid) (Direction, int) {
	direction := t.Direction
	switch turn {
	case R:
		direction = direction + 1
	case L:
		direction = direction + 3
	case U:
		direction = direction + 2
	}
	direction = direction % 4
	max := grid.X * grid.Y
	position := t.Location

	// Update position - wrap if needed
	switch direction {
	case North:
		if (position - grid.Y) >= 0 {
			position = position - grid.Y
		} else {
			position = ((grid.X - 1) * grid.Y) + position
		}
	case East:
		if position+1 < max {
			position = position + 1
		} else {
			position = ((grid.X - 1) * grid.Y) + 1
		}
	case South:
		if (position + grid.Y) < max {
			position = position + grid.Y
		} else {
			position = position + grid.Y - ((grid.X - 1) * grid.Y)
		}
	case West:
		if position != 0 {
			position = position - 1
		} else {
			position = (grid.X - 1) * grid.Y
		}
	}
	return direction, position
}
