package mound

import (
	"fmt"

	colorful "github.com/lucasb-eyer/go-colorful"
)

// Mound is the main structure of
type Mound struct {
	Grid    Grid
	Turmite Turmite
	Default Move
}

// CreateMound builds the main structure that contains the grid and the turmite
func CreateMound(blockSize int, x int, y int, d Direction, l int, r Rule, def Move) Mound {
	g := createGrid(blockSize, x, y)
	t := Turmite{
		Direction: d,
		Location:  l,
		Rule:      r,
	}
	m := Mound{
		Grid:    g,
		Turmite: t,
		Default: def,
	}
	return m
}

// currentColor returns the color under the turmite
func (m Mound) currentColor() colorful.Color {
	l := m.Turmite.Location
	return m.Grid.S[l]
}

// Next mutates the mound and moves it forward one tick
func (m *Mound) Next() {
	move, ok := m.Turmite.findMove(m.currentColor())
	if ok == false {
		move = m.Default
	}
	m.Grid.updateColor(m.Turmite.Location, move.C)
	fmt.Printf("Loop: #%v \n", m.Turmite.Location)
	fmt.Printf("Loop: #%v \n", m.Grid.S[m.Turmite.Location])
	d, l := m.Turmite.move(move.T, m.Grid)
	m.Turmite.Location = l
	m.Turmite.Direction = d
}
