package main

import (
	"fmt"
	"os"
	"os/exec"
)

func _main() {
	// Run the tmux capture-pane command
	cmd := exec.Command("tmux", "capture-pane", "-p", "-e")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Write the raw output to a file
	err = os.WriteFile("tmux_out.txt", out, 0664)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Raw terminal output saved to tmux_output.txt")
}

