package main

import (
	"image/png"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/sdwalsh/turmites/mound"
)

func main() {
	r := mound.CreateRules()

	c1, _ := colorful.Hex("#1C2321")
	c2, _ := colorful.Hex("#7D98A1")
	r.AddRule(c1, c2, mound.L)
	r.AddRule(c2, c1, mound.R)

	m := mound.CreateMound(5, 600, 600, mound.North, 124, *r)
	e := png.Encoder{
		CompressionLevel: -3,
	}

	m.Next()

}
