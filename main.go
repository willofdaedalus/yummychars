package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func main() {
	// Get the file descriptor for stdin
	fd := int(os.Stdin.Fd())

	// Save the current terminal state
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		fmt.Println("Error setting raw mode:", err)
		return
	}

	// Ensure to restore the old state on exit
	defer func() {
		if err := term.Restore(fd, oldState); err != nil {
			fmt.Println("Error restoring terminal state:", err)
		}
	}()

	fmt.Println("Terminal is now in raw mode. Press Ctrl+C to exit.")

	// Reading from stdin in raw mode
	buf := make([]byte, 1)
	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			return
		}

		// Check if the input is Ctrl+C (ASCII code 3)
		if buf[0] == '\x03' {
			fmt.Println("Ctrl+C pressed. Exiting.")
			break
		}

		fmt.Printf("\rRead: %c\n", buf[0])
	}
}
