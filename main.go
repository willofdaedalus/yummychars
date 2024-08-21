package main

import (
	"log"
	"os"
	"time"
	"willofdaedalus/yummychars/serpent"
)

func main() {
	oldState, fd, err := setupTerminal()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanUp(fd, oldState)

	// get and spawn snake at cursor location
	// r, c, err := getCursorPosition()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	dir := serpent.RIGHT
	s := serpent.InitSnake(10, 0, 0)
	s.MoveSnake(dir)

	buf := make([]byte, 1)
	for {
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

		// add a short sleep to control the loop speed
		// this isn't the best but it works; might come back this
		time.Sleep(time.Second / time.Duration(s.Speed))

		// clear the previous frames to remove smears
		s.ClearScreen()
	}
}
