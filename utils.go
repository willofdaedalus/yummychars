package main

import (
	// "atomicgo.dev/cursor"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

func cleanUp(fd int, orig *term.State) {
	term.Restore(fd, orig)
	// ensure the cursor is shown again when the program exits
	fmt.Print("\033[?25h")
}

func setupTerminal() (*term.State, int, error) {
	fmt.Printf("\033[2J\033[H") // clear screen
	fmt.Print("\033[?25l")      // hide cursor
	fd := int(os.Stdin.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return nil, -1, err
	}

	return oldState, fd, nil
}

func getCursorPosition() (int, int, error) {
	fmt.Print("\033[6n")
	var buf []byte
	n, err := os.Stdin.Read(buf)
	if err != nil {
		return -1, -1, err
	}

	// output comes as something like ^[[17;1R%
	res := string(buf[:n])

	if strings.HasPrefix(res, "\033[") && strings.HasSuffix(res, "R") {
		res = res[2:len(res) - 1]
		splits := strings.Split(res, ";")

		if len(splits) == 2 {
			row, err1 := strconv.Atoi(splits[0])
			col, err2 := strconv.Atoi(splits[1])

			if err1 != nil || err2 != nil {
				return -1, -1, fmt.Errorf("failed to get cursor position")
			}

			return row, col, nil
		}
	}

	return -1, -1, nil
}
