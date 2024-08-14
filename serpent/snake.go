package serpent

import "fmt"

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

const HEAD = ">"

type coords struct {
	X, Y int
}

type Snake struct {
	Position coords
	Speed    float64
	Tail     []string
}

func InitSnake(speed float64) *Snake {
	return &Snake{
		Position: coords{0, 0},
		Speed:    speed,
		Tail:     make([]string, 0),
	}
}

func (s *Snake) MoveSnake(dir int) {
	switch dir {
	case UP:
		s.Position.Y -= 1
	case DOWN:
		s.Position.Y += 1
	case LEFT:
		s.Position.X -= 1
	case RIGHT:
		s.Position.X += 1
	}
}

func (s *Snake) DrawSnake() {
	fmt.Printf("x: %d, y: %d", s.Position.X, s.Position.Y)
}
