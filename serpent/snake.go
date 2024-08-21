package serpent

import "fmt"

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

const (
	HEAD_R = '>'
	HEAD_L = '<'
	HEAD_U = 'ʌ'
	HEAD_D = 'V'
	BODY   = 'o'
)

type coords struct {
	X, Y int
}

type Snake struct {
	Head     rune
	Position coords
	Speed    float64
	Tail     []string
}

func InitSnake(speed float64) *Snake {
	return &Snake{
		Head: HEAD_R,
		Position: coords{0, 0},
		Speed:    speed,
		Tail:     make([]string, 0),
	}
}

func (s *Snake) MoveSnake(dir int) {
	switch dir {
	case UP:
		s.Position.Y -= 1
		s.Head = HEAD_U
	case DOWN:
		s.Position.Y += 1
		s.Head = HEAD_D
	case LEFT:
		s.Position.X -= 1
		s.Head = HEAD_L
	case RIGHT:
		s.Position.X += 1
		s.Head = HEAD_R
	}
}

func (s *Snake) ClearScreen() {
	fmt.Printf("\033[2J\033[H")
}

func (s *Snake) DrawSnake() {
	// apparently the format is actually (y, x) and not (x, y)
	fmt.Printf("\033[%d;%dH%c", s.Position.Y, s.Position.X, s.Head)
}
