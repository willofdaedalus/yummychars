package slideys

import "fmt"

type Slidey struct {
	content [][]rune
	sx, sy  int
}

func InitSlideys(sx, sy int, content [][]rune) *Slidey {
	return &Slidey{
		content: content,
		sx: sx,
		sy: sy,
	}
}

func (s *Slidey) PrintSlideyLeft() {
	_, right := s.splitRuneContent()
	for _, line := range right {
		fmt.Printf("%s\n\r", string(line))
	}
}
