package mound

import colorful "github.com/lucasb-eyer/go-colorful"

// Direction determines the next location of the turmite
type Direction int

// [L]eft [R]igth [U]-turn [N]o-turn
const (
	N Direction = iota
	S
	E
	W
)

// Color is the array of colors
type Color struct {
	C colorful.Color
}

// Move defines the color to replace under the turmite and the direction the turmite should move
type Move struct {
	C Color
	D Direction
}

// Rule is a single direction in the form of a dictionary
type Rule struct {
	R map[Color]Move
}

// Turmite is an individual that exists inside a grid
// X and Y are positions of size defined in a the Block of a Grid
type Turmite struct {
	D Direction
	X int
	Y int
	R Rule
}
