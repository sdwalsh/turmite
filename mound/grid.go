package mound

import (
	"errors"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"time"

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

// numberOfDigits calculates the number of digits in a given int using a lookup table
// less expensive than the alternative of math.Log10 calls
func numberOfDigits(x int) (int, error) {
	switch {
	case x < 10:
		return 1, nil
	case x < 100:
		return 2, nil
	case x < 1000:
		return 3, nil
	case x < 10000:
		return 4, nil
	case x < 100000:
		return 5, nil
	case x < 1000000:
		return 6, nil
	default:
		return -1, errors.New("Integer size not supported")
	}
}

func zeroString(zeros int) string {
	var zeroString string
	for x := 0; x < zeros; x++ {
		zeroString = zeroString + "0"
	}
	return zeroString
}

// BatchImages creates a temporary folder and takes
func (g *Grid) BatchImages(numberOfImages int, squareSize int) error {
	// Create directory to store images in before encoding
	directory := "batch" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	totalDigits, err := numberOfDigits(numberOfImages)
	if err != nil {
		return err
	}
	err = os.Mkdir(directory, os.FileMode(int(0777)))
	if err != nil {
		return err
	}
	e := png.Encoder{
		CompressionLevel: -3,
	}

	for x := 0; x < numberOfImages; x++ {
		digits, err := numberOfDigits(x)
		if err != nil {
			return err
		}
		zeros := totalDigits - digits
		filepath := filepath.Join(directory, zeroString(zeros)+"color.png")
		file, err := os.Create(filepath)
		if err != nil {
			return err
		}
		img := g.GridToImage(squareSize)
		err = e.Encode(file, img)
		if err != nil {
			return err
		}
	}
	return nil
}
