package main

import (
	"fmt"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/sdwalsh/turmites/mound"
)

func main() {
	r := mound.CreateRules()
	fmt.Println("creating mound")
	c1, _ := colorful.Hex("#FF1955")
	c2, _ := colorful.Hex("#00BFFF")
	c3, _ := colorful.Hex("#CCC014")
	r.AddRule(c1, c2, mound.L)
	r.AddRule(c2, c3, mound.R)
	r.AddRule(c3, c1, mound.N)

	def := mound.Move{
		C: c1,
		T: 1,
	}

	var t []*mound.Turmite
	t = append(t, mound.CreateTurmite(mound.North, 789, *r))

	m, _ := mound.CreateMound(5, 50, 50, t, def)
	directory, _ := m.BatchImages(50, 100)
	mound.ConvertPngToMp4(directory, "color2", 20)
}
