package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"willofdaedalus/yummychars/serpent"
	"willofdaedalus/yummychars/slideys"
	"willofdaedalus/yummychars/utils"

	"golang.org/x/term"
)

var (
	content    [][]rune
	oldState   *term.State
	sx, sy, fd int
	err        error
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("please provide an argument. see help")
		return
	}
	content, err = utils.SetupContent()
	if err != nil {
		log.Fatal(err)
	}
	oldState, fd, err = utils.SetupTerminal()
	if err != nil {
		log.Fatal(err)
	}
	sx, sy, err = term.GetSize(fd)
	if err != nil {
		log.Fatal(err)
	}
	defer utils.CleanUp(fd, oldState)

	game, err := initApp(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	game()
}

func initApp(cmd string) (func(), error) {
	switch cmd {
	case "snake":
		return snakeEnvironment, nil
	case "slidey":
		return slideyEnvironment, nil
	}

	return nil, fmt.Errorf("your command %s is not known. see help")
}

func slideyEnvironment() {
	s := slideys.InitSlideys(sx, sy, content)
	s.PrintSlideyLeft()
}

func snakeEnvironment() {
	s := serpent.InitSnake(10, sx, sy, content)
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

		// exit the game if the snake collides with the boundaries
		// don't know if checking the [][]rune every "frame" is efficient
		if s.CheckBoundaries() || s.WinConditionLogic() {
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

