package mound

import (
	"image"
	"image/draw"

	colorful "github.com/lucasb-eyer/go-colorful"
)

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
	S []colorful.Color
}

// CreateGrid creates an empty grid
func createGrid(blockSize, x, y int) Grid {
	block := Block{x: blockSize, y: blockSize}
	grid := Grid{
		B: block,
		X: x,
		Y: y,
		S: make([]colorful.Color, x*y),
	}
	return grid
}

// updateColor mutates the grid to update the color at the provided location
func (g *Grid) updateColor(location int, c colorful.Color) {
	g.S[location] = c
}

// GridToImage transforms a grid into an image.Image to encoded later
func (g *Grid) GridToImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, g.X, g.Y))
	location := 0

	for x := 0; x < g.X; x++ {
		for y := 0; y < g.Y; y++ {
			draw.Draw(img, image.Rect(x, y, x+1, y+1), &image.Uniform{g.S[location]}, image.ZP, draw.Src)
			location++
		}
	}
	return img
}
