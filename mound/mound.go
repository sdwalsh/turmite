package mound

// Mound is the main structure of
type Mound struct {
	Grid    Grid
	Turmite Turmite
}

// CreateMound builds the main structure that contains the grid and the turmite
func CreateMound(blockSize int, x int, y int, d Direction, l int, r Rule) Mound {
	g := createGrid(blockSize, x, y)
	t := Turmite{
		Direction: d,
		Location:  l,
		Rule:      r,
	}
	m := Mound{
		Grid:    g,
		Turmite: t,
	}
	return m
}

func (m Mound) currentColor() Color {
	l := m.Turmite.Location
	return m.Grid.S[l]
}

func findMove(c Color, t Turmite) Move {
	return t.Rule[c]
}

func (m *Mound) next() {
	move := findMove(m.currentColor(), m.Turmite)
	m.Grid.updateColor(m.Turmite.Location, move.C)
	d, l := m.Turmite.move(move.T, m.Grid)
	m.Turmite.Location = l
	m.Turmite.Direction = d

}
