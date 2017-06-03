package mound

// Mound is the main structure of
type Mound struct {
	G Grid
	T Turmite
}

// CreateMound builds the main structure that contains the grid and the turmite
func CreateMound(blockSize int, x int, y int, d Direction, tx int, ty int, r Rule) Mound {
	g := CreateGrid(blockSize, x, y)
	t := Turmite{
		D: d,
		X: tx,
		Y: ty,
		R: r,
	}
	m := Mound{
		G: g,
		T: t,
	}
	return m
}
