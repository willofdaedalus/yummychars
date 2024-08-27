package utils

type Coords struct {
	X, Y int
}

type Cell struct {
	c        rune
	colour   string
	position Coords
}

func NewCoords(x, y int) Coords {
	return Coords{x, y}
}
