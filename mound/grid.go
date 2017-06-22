package mound

import (
	"image"
	"image/draw"

	colorful "github.com/lucasb-eyer/go-colorful"
	log "github.com/sirupsen/logrus"
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
	log.WithFields(log.Fields{
		"color": g.S[location].Hex(),
	}).Info("color under turmite before update")

	g.S[location] = c

	log.WithFields(log.Fields{
		"color": g.S[location].Hex(),
	}).Info("color under turmite after update")
}

// GridToImage transforms a grid into an image.Image to encoded later
func (g *Grid) GridToImage(squareSize int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, g.X*squareSize, g.Y*squareSize))
	location := 0

	for x := 0; x < g.X*squareSize; x += squareSize {
		for y := 0; y < g.Y*squareSize; y += squareSize {
			draw.Draw(img, image.Rect(x, y, x+squareSize, y+squareSize), &image.Uniform{g.S[location]}, image.ZP, draw.Src)
			log.WithFields(log.Fields{
				"location":    location,
				"color":       g.S[location].Hex(),
				"square size": squareSize,
			}).Info("drawing to image")
			location++
		}
	}
	return img
}
