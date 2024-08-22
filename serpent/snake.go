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
	x, y int
}

type Snake struct {
	MoveDir  int
	Speed    float64
	head     rune
	position coords
	fieldSize coords
	tail     []coords
}

func InitSnake(speed float64, length, fx, fy int) *Snake {
	// initialize snake with the head position and an empty tail
	return &Snake{
		head:     HEAD_R,
		position: coords{0, 0},
		Speed:    speed,
		tail:     make([]coords, length), // make tail with a length of 4
		fieldSize: coords{fx, fy},
	}
}

func (s *Snake) MoveSnake(dir int) {
	// basically move all segments but the first at the position of the one
	// infront of it and only update the first segment to the head's position
	for i := len(s.tail) - 1; i > 0; i-- {
		s.tail[i] = s.tail[i-1]
	}

	// move the first tail segment to the previous position of the head
	if len(s.tail) > 0 {
		s.tail[0] = s.position
	}

	// move the head based on direction
	switch dir {
	case UP:
		s.position.y -= 1
		s.head = HEAD_U
	case DOWN:
		s.position.y += 1
		s.head = HEAD_D
	case LEFT:
		s.position.x -= 1
		s.head = HEAD_L
	case RIGHT:
		s.position.x += 1
		s.head = HEAD_R
	}

	s.MoveDir = dir
}

func (s *Snake) ClearScreen() {
	fmt.Printf("\033[2J\033[H")
}

func (s *Snake) CheckBoundaries() bool {
	if (s.position.x > s.fieldSize.x || s.position.x < 0) || (s.position.y > s.fieldSize.y || s.position.y < 0) {
		s.ClearScreen()
		fmt.Printf("\033[%d;%dH%s", s.fieldSize.y/2, s.fieldSize.x/2, "game over!")
		return true
	}

	return false
}

func (s *Snake) DrawSnake() {
	// draw the head of the snake
	fmt.Printf("\033[%d;%dH%c", s.position.y, s.position.x, s.head)
	// draw the tail of the snake
	for _, segment := range s.tail {
		fmt.Printf("\033[%d;%dH%c", segment.y, segment.x, BODY)
	}
}

