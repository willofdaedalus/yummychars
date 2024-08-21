package serpent

import (
	"fmt"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

const (
	HEAD_R = '>'
	HEAD_L = '<'
	HEAD_U = '^'
	HEAD_D = 'v'
	BODY   = 'o'
)

type coords struct {
	X, Y int
}

type Snake struct {
	MoveDir  int
	Head     rune
	Position coords
	Speed    float64
	Tail     []coords
}

func InitSnake(speed float64, x, y int) *Snake {
	// Initialize snake with the head position and an empty tail
	return &Snake{
		Head:     HEAD_R,
		Position: coords{x, y},
		Speed:    speed,
		Tail:     make([]coords, 4), // Make tail with a length of 4
	}
}

func (s *Snake) MoveSnake(dir int) {
	// move the tail segments
	for i := len(s.Tail) - 1; i > 0; i-- {
		s.Tail[i] = s.Tail[i-1]
	}

	// move the first tail segment to the previous position of the head
	if len(s.Tail) > 0 {
		s.Tail[0] = s.Position
	}

	// move the head based on direction
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

	s.MoveDir = dir
}

func (s *Snake) ClearScreen() {
	fmt.Printf("\033[2J\033[H")
}

func (s *Snake) DrawSnake() {
	// Draw the head of the snake
	fmt.Printf("\033[%d;%dH%c", s.Position.Y, s.Position.X, s.Head)

	// Draw the tail of the snake
	for _, segment := range s.Tail {
		fmt.Printf("\033[%d;%dH%c", segment.Y, segment.X, BODY)
	}
}

