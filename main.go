package main

import (
	"fmt"
	"image/png"
	"os"
	"strconv"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/sdwalsh/turmites/mound"
)

func main() {
	r := mound.CreateRules()
	fmt.Println("creating mound")
	c1, _ := colorful.Hex("#1C2321")
	c2, _ := colorful.Hex("#7D98A1")
	r.AddRule(c1, c2, mound.L)
	r.AddRule(c2, c1, mound.R)

	def := mound.Move{
		C: c1,
		T: 1,
	}

	m := mound.CreateMound(5, 200, 200, mound.North, 124, *r, def)
	e := png.Encoder{
		CompressionLevel: -3,
	}
	num := 0

	for x := 0; x < 10000000; x++ {
		m.Next()
	}
	fmt.Println("begin rendering")
	img := m.Grid.GridToImage()
	toImg, _ := os.Create("color" + strconv.Itoa(num) + ".png")
	e.Encode(toImg, img)
}
