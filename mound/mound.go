package mound

// Mound is the main structure of
type Mound struct {
	Grid     Grid
	Turmites []*Turmite
	Default  Move
}

// CreateMound builds the main structure that contains the grid and a slice of turmites
func CreateMound(blockSize int, x int, y int, t []*Turmite, def Move) Mound {
	g := createGrid(blockSize, x, y)
	m := Mound{
		Grid:     g,
		Turmites: t,
		Default:  def,
	}
	return m
}

// Next mutates the mound and moves it forward one tick
func (m *Mound) Next() {
	for _, t := range m.Turmites {
		move, ok := t.findMove(m.Grid.currentColor(*t))
		if ok == false {
			move = m.Default
		}
		m.Grid.updateColor(t.Location, move.C)
		d, l := t.move(move.T, m.Grid)
		t.Location = l
		t.Direction = d
	}
}
