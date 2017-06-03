package mound

// Block is the unit of measurement in a grid
// x and y are measured in pixels
type Block struct {
	x int
	y int
}

// Grid is the basic structure for the turmites
// x and y are measured in blocks
// S is the state of the grid
type Grid struct {
	B Block
	X int
	Y int
	S []Color
}

// CreateGrid creates an empty grid
func CreateGrid(blockSize, x, y int) Grid {
	block := Block{x: blockSize, y: blockSize}
	grid := Grid{
		B: block,
		X: x,
		Y: y,
		S: make([]Color, x*y),
	}
	return grid
}

func grid() {

}
