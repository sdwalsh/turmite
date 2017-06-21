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

	var t []*mound.Turmite
	t = append(t, mound.CreateTurmite(mound.North, 40, *r))

	m := mound.CreateMound(5, 50, 50, t, def)
	e := png.Encoder{
		CompressionLevel: -3,
	}
	num := 0

	for x := 0; x < 100000; x++ {
		m.Next()
	}
	fmt.Println("begin rendering")
	img := m.Grid.GridToImage(100)
	toImg, _ := os.Create("color" + strconv.Itoa(num) + ".png")
	e.Encode(toImg, img)
}
