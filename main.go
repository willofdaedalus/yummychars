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
	fmt.Printf("\033[2J\033[H")
	fd := int(os.Stdin.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		log.Fatal("error occured putting terminal in raw mode\n")
	}
	defer term.Restore(fd, oldState)

	dir := serpent.RIGHT
	s := serpent.InitSnake(1)
	s.MoveSnake(dir)

	buf := make([]byte, 1)
	for {
		// non-blocking input with a goroutine
		go func() {
			_, err := os.Stdin.Read(buf)
			if err != nil {
				if os.IsTimeout(err) {
					return
				}
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
			fmt.Printf("\033[2J\033[H")
			return
		}
		s.MoveSnake(dir)
		s.DrawSnake()

		// add a short sleep to control the loop speed
		// this isn't the best but it works; might come back this
		time.Sleep(time.Second / time.Duration(s.Speed))
	}
}
