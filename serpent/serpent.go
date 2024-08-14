package serpent

type coords struct {
	X, Y int
}

type Snake struct {
	Position coords
	Tail []string
}

func InitSnake() *Snake {
	return &Snake{
		Position: coords{},
		Tail: make([]string, 0),
	}
}

func drawSnake() {
}
