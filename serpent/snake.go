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
	MoveDir     int
	Speed       float64
	TermContent [][]rune

	head        rune
	actualChars [][]rune
	position    coords
	fieldSize   coords
	tail        []coords
}

func InitSnake(speed float64, length, fx, fy int, rawContent [][]rune) *Snake {
	// initialize snake with the head position and an empty tail
	return &Snake{
		head:        HEAD_R,
		position:    coords{0, 0},
		Speed:       speed,
		tail:        make([]coords, length), // make tail with a length of 4
		fieldSize:   coords{fx, fy},
		TermContent: rawContent,
		actualChars: stripAnsiCodes(rawContent),
	}
}

func (s *Snake) ClearScreen() {
	fmt.Printf("\033[2J\033[H")
}

func (s *Snake) CheckBoundaries() bool {
	if (s.position.x > s.fieldSize.x || s.position.x < 0) ||
		(s.position.y > s.fieldSize.y || s.position.y < 0) {
		s.ClearScreen()
		fmt.Printf("\033[%d;%dH%s", s.fieldSize.y/2, s.fieldSize.x/2, "game over!")
		return true
	}

	return false
}

func (s *Snake) MoveSnake(dir int) {
	// Move all segments of the tail except the first one
	for i := len(s.tail) - 1; i > 0; i-- {
		s.tail[i] = s.tail[i-1]
	}

	// Move the first tail segment to the previous position of the head
	if len(s.tail) > 0 {
		s.tail[0] = s.position
	}

	// Update the head's position based on direction
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

	if s.position.y >= 0 && s.position.y < len(s.actualChars) &&
		s.position.x >= 0 && s.position.x < len(s.actualChars[s.position.y]) {
		if s.actualChars[s.position.y][s.position.x] != ' ' {
			s.actualChars[s.position.y][s.position.x] = ' '

			// Update TermContent to reflect the change
			s.updateTermContent(s.position.y, s.position.x)
		}
	}
}

func (s *Snake) updateTermContent(y, x int) {
	actualX := 0
	inEscapeSeq := false
	for i := range s.TermContent[y] {
		if s.TermContent[y][i] == '\033' {
			inEscapeSeq = true
		} else if !inEscapeSeq {
			if actualX == x {
				s.TermContent[y][i] = ' '
				return
			}
			actualX++
		} else if (s.TermContent[y][i] >= 'A' && s.TermContent[y][i] <= 'Z') ||
			(s.TermContent[y][i] >= 'a' && s.TermContent[y][i] <= 'z') {
			inEscapeSeq = false
		}
	}
}

func (s *Snake) DrawScreenContent() {
	// Draw the captured terminal content
	for y, line := range s.TermContent {
		fmt.Printf("\033[%d;1H%s", y+1, string(line))
	}
}

func (s *Snake) DrawSnake() {
	// Draw the head of the snake
	fmt.Printf("\033[%d;%dH%c", s.position.y+1, s.position.x+1, s.head)
	// Draw the tail of the snake
	for _, segment := range s.tail {
		fmt.Printf("\033[%d;%dH%c", segment.y+1, segment.x+1, BODY)
	}
}
