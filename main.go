package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"willofdaedalus/yummychars/serpent"

	"golang.org/x/term"
)

func main() {
	content, err := setupContent()
	if err != nil {
		log.Fatal(err)
	}

	oldState, fd, err := setupTerminal()
	if err != nil {
		log.Fatal(err)
	}
	// use the size of the terminal to determine the boundaries for the snake
	sx, sy, err := term.GetSize(fd)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanUp(fd, oldState)

	s := serpent.InitSnake(10, 5, sx, sy)
	s.TermContent = content
	dir := serpent.RIGHT
	s.MoveSnake(dir)

	buf := make([]byte, 1)
	for {
		// Draw the captured content
		for y, line := range s.TermContent {
			fmt.Printf("\033[%d;1H%s", y+1, string(line))
		}

		go func() {
			_, err := os.Stdin.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
		}()

		switch buf[0] {
		case 'w':
			if s.MoveDir != serpent.DOWN {
				dir = serpent.UP
			}
		case 'a':
			if s.MoveDir != serpent.RIGHT {
				dir = serpent.LEFT
			}
		case 's':
			if s.MoveDir != serpent.UP {
				dir = serpent.DOWN
			}
		case 'd':
			if s.MoveDir != serpent.LEFT {
				dir = serpent.RIGHT
			}
		case 'q':
			s.ClearScreen()
			return
		}
		s.MoveSnake(dir)
		s.DrawSnake()
		// exit the game when the snake collides with the edges
		if s.CheckBoundaries() {
			time.Sleep(time.Second * 2)
			s.ClearScreen()
			break
		}

		// add a short sleep to control the loop speed
		// this isn't the best but it works; might come back this
		time.Sleep(time.Second / time.Duration(s.Speed))

		// clear the previous frames to remove smears
		s.ClearScreen()
	}
}
