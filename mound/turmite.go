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

// Color is the array of colors
type Color struct {
	C colorful.Color
}

// Move defines the color to replace under the turmite and the direction the turmite should move
type Move struct {
	C Color
	T Turn
}

// Rule is a single direction in the form of a dictionary
type Rule struct {
	R map[Color]Move
}

// Turmite is an individual that exists inside a grid
// X and Y are positions of size defined in a the Block of a Grid
type Turmite struct {
	D        Direction
	Location int
	R        Rule
}

// Move takes a turmite, a turn, and a mound and returns the new direction and the update position of the turmite
func (t Turmite) move(turn Turn, mound Mound) (Direction, int) {
	direction := t.D
	switch turn {
	case R:
		direction = direction + 1
	case L:
		direction = direction + 3
	case U:
		direction = direction + 2
	}

	direction = direction % 4
	max := mound.G.X * mound.G.Y
	position := t.Location

	// Update position - wrap if needed
	switch direction {
	case North:
		if (position - mound.G.Y) > 0 {
			position = position - mound.G.Y
		} else {
			position = (mound.G.X * (mound.G.Y - 1)) + position
		}
	case East:
		if position+1 < max {
			position = position + 1
		} else {
			position = (mound.G.X * (mound.G.Y - 1)) + 1
		}
	case South:
		if (position + mound.G.Y) < max {
			position = position + mound.G.Y
		} else {
			position = position - (mound.G.X * (mound.G.Y - 1))
		}
	case West:
		if position != 0 {
			position = position - 1
		} else {
			position = mound.G.Y
		}
	}
	return direction, position
}
