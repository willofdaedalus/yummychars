package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"willofdaedalus/yummychars/serpent"

	"golang.org/x/term"
)

func cleanUp(fd int, orig *term.State) {
	term.Restore(fd, orig)
	// ensure the cursor is shown again when the program exits
	fmt.Print("\033[?25h") 
}

func setupTerminal() (*term.State, int, error) {
	fmt.Printf("\033[2J\033[H")
	fmt.Print("\033[?25l")
	fd := int(os.Stdin.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return nil, -1, err
	}

	return oldState, fd, nil
}

func main() {
	oldState, fd, err := setupTerminal()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanUp(fd, oldState)

	dir := serpent.RIGHT
	s := serpent.InitSnake(5)
	s.MoveSnake(dir)

	buf := make([]byte, 1)
	for {
		// non-blocking input with a goroutine
		go func() {
			_, err := os.Stdin.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
		}()

		switch buf[0] {
		case 'w':
			dir = serpent.UP
		case 'a':
			dir = serpent.LEFT
		case 's':
			dir = serpent.DOWN
		case 'd':
			dir = serpent.RIGHT
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
