package mound

import (
	"fmt"
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
	block := Block{1, 1}
	grid := Grid{
		B: block,
		X: x,
		Y: y,
		S: make([]colorful.Color, x*y),
	}
	return grid
}

// currentColor returns the color under the turmite
func (g Grid) currentColor(t Turmite) colorful.Color {
	l := t.Location
	return g.S[l]
}

// updateColor mutates the grid to update the color at the provided location
func (g *Grid) updateColor(location int, c colorful.Color) {
	fmt.Printf("Pre-update: #%v \n", g.S[location])
	g.S[location] = c
	fmt.Printf("Post-update: #%v \n", g.S[location])
	fmt.Printf("Should: #%v \n", c)
	fmt.Println("------------------------------------------------")
}

// GridToImage transforms a grid into an image.Image to encoded later
func (g *Grid) GridToImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, g.X*5, g.Y*5))
	location := 0

	for x := 0; x < g.X*5; x += 5 {
		for y := 0; y < g.Y*5; y += 5 {
			draw.Draw(img, image.Rect(x, y, x+5, y+5), &image.Uniform{g.S[location]}, image.ZP, draw.Src)
			fmt.Printf("Drawing: #%v \n", g.S[location])
			location++
		}
	}
	return img
}
