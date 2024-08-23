package main

import (
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
	defer cleanUp(fd, oldState)

	sx, sy, err := term.GetSize(fd)
	if err != nil {
		log.Fatal(err)
	}

	s := serpent.InitSnake(10, 5, sx, sy)
	s.TermContent = content
	dir := serpent.RIGHT

	buf := make([]byte, 1)
	for {
		s.DrawScreenContent()

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

		// Exit the game if the snake collides with the boundaries
		if s.CheckBoundaries() {
			time.Sleep(time.Second * 2)
			s.ClearScreen()
			break
		}

		// Add a short sleep to control the loop speed
		time.Sleep(time.Second / time.Duration(s.Speed))

		// Clear the previous frames to remove smears
		s.ClearScreen()
	}
}

