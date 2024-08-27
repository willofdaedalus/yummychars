package serpent

import (
	"fmt"
	"willofdaedalus/yummychars/utils"
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

type Snake struct {
	MoveDir     int
	Speed       float64
	TermContent [][]rune

	colour      string
	head        rune
	actualChars [][]rune
	position    utils.Coords
	fieldSize   utils.Coords
	tail        []utils.Coords
}

func InitSnake(speed float64, fx, fy int, rawContent [][]rune) *Snake {
	// initialize snake with the head position and an empty tail
	return &Snake{
		head:        HEAD_R,
		position:   utils.Coords{X: fx, Y: fy},
		Speed:       speed,
		// make tail with a length of 4 so that I don't have to figure out self collision logic ;P
		tail:        make([]utils.Coords, 4), 
		fieldSize:   utils.Coords{X: fx, Y: fy},
		TermContent: rawContent,
		actualChars: stripAnsiCodes(rawContent),
	}
}

func (s *Snake) ClearScreen() {
	fmt.Printf("\033[2J\033[H")
}

func (s *Snake) WinConditionLogic() bool {
	for _, row := range s.actualChars {
		for _, c := range row {
			if c != ' ' {
				return false
			}
		}
	}

	fmt.Printf("\033[%d;%dH%s", s.fieldSize.Y/2, s.fieldSize.X/2, "you win!")
	return true
}

func (s *Snake) CheckBoundaries() bool {
	// allows the snake move along the boundaries without punishing the player
	if (s.position.X > s.fieldSize.X + 1 || s.position.X < -1) ||
		(s.position.Y > s.fieldSize.Y + 1 || s.position.Y < -1) {
		s.ClearScreen()
		fmt.Printf("\033[%d;%dH%s", s.fieldSize.Y/2, s.fieldSize.X/2, "game over!")
		return true
	}

	return false
}

func (s *Snake) MoveSnake(dir int) {
	// move all segments of the tail except the first one
	for i := len(s.tail) - 1; i > 0; i-- {
		s.tail[i] = s.tail[i-1]
	}

	// move the first tail segment to the previous position of the head
	if len(s.tail) > 0 {
		s.tail[0] = s.position
	}

	// update the head's position based on direction
	switch dir {
	case UP:
		s.position.Y -= 1
		s.head = HEAD_U
	case DOWN:
		s.position.Y += 1
		s.head = HEAD_D
	case LEFT:
		s.position.X -= 1
		s.head = HEAD_L
	case RIGHT:
		s.position.X += 1
		s.head = HEAD_R
	}

	s.MoveDir = dir

	if s.position.Y >= 0 && s.position.Y < len(s.actualChars) &&
		s.position.X >= 0 && s.position.X < len(s.actualChars[s.position.Y]) {
		if s.actualChars[s.position.Y][s.position.X] != ' ' {
			s.actualChars[s.position.Y][s.position.X] = ' '

			// update termcontent to reflect the change
			s.updateTermContent(s.position.Y, s.position.X)
		}
	}
}

func (s *Snake) DrawScreenContent() {
	// draw the captured terminal content
	for Y, line := range s.TermContent {
		fmt.Printf("\033[%d;1H%s", Y+1, string(line))
	}
}

func (s *Snake) DrawSnake() {
	// draw the head of the snake with the current colour
	fmt.Printf("%s\033[%d;%dH%c\033[0m", s.colour, s.position.Y+1, s.position.X+1, s.head)

	// draw the tail of the snake with the current colour
	for _, segment := range s.tail {
		fmt.Printf("%s\033[%d;%dH%c\033[0m", s.colour, segment.Y+1, segment.X+1, BODY)
	}
}
