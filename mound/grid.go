package mound

// Block is the unit of measurement in a grid
// x and y are measured in pixels
type Block struct {
	x int
	y int
}

// Grid is the basic structure for the turmites
// x and y is measured in blocks
type Grid struct {
	B Block
	X int
	Y int
}

// CreateGrid ...
func CreateGrid(blockSize, x, y int) Grid {
	block := Block{x: blockSize, y: blockSize}
	grid := Grid{
		B: block,
		X: x,
		Y: y,
	}
	return grid
}

func grid() {

}
