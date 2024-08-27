package utils

import "fmt"

type Coords struct {
	X, Y int
}

type Cell struct {
	Char     rune
	Colour   int
	BgColour int
	Bold     bool
	Position Coords
}

// convert a 2D array of runes to cells which are essentially runes but with properties
func RuneToCells() ([][]Cell, error) {
	content, err := GetTerminalContent()
	if err != nil {
		return nil, err
	}

	var finalColour, bgColour int
	var bold bool
	codedCells := make([][]Cell, len(content))

	for x := range content {
		codedCells[x] = make([]Cell, 0, len(content[x])) // pre-allocate
		for y := 0; y < len(content[x]); y++ {
			if content[x][y] == '\033' {
				var err error
				// extract our values and update
				finalColour, bgColour, bold, y, err = parseAnsiSequence(content[x], y)
				if err != nil {
					return nil, err
				}
			} else {
				c := Cell{
					Char:     content[x][y],
					Colour:   finalColour,
					BgColour: bgColour,
					Bold:     bold,
					Position: Coords{x, y},
				}
				codedCells[x] = append(codedCells[x], c)
			}
		}
	}

	return codedCells, nil
}

// get and extract the relevant values for the cell
func parseAnsiSequence(line []rune, start int) (fg, bg int, bold bool, newStart int, err error) {
	// confirm the start we're passing is actually valid before processing
	if start >= len(line) || line[start] != '\033' {
		return 0, 0, false, start, fmt.Errorf("invalid ANSI sequence start")
	}

	// skip to the [ character in the sequence
	newStart = start + 1
	if newStart >= len(line) || line[newStart] != '[' {
		return 0, 0, false, start, fmt.Errorf("invalid ANSI sequence format")
	}

	newStart++
	sequence := make([]int, 0, 4)
	current := 0

	for newStart < len(line) {
		if isNum(line[newStart]) {
			// accumulate digits into a single number
			current = current*10 + int(line[newStart]-'0')
		} else if line[newStart] == ';' {
			// add the current value to the sequence reset
			sequence = append(sequence, current)
			current = 0
		} else if line[newStart] == 'm' {
			// we're at the end of the sequence so we can exit the loop
			sequence = append(sequence, current)
			newStart++ // skip the m at the end of the sequences
			break
		} else {
			// handle unexpected characters in the ansi sequence
			return 0, 0, false, start, fmt.Errorf("unexpected character in ANSI sequence")
		}
		newStart++
	}

	// parse the sequence to extract fg, bg, and bold values
	for _, code := range sequence {
		switch {
		// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797#8-16-colors
		case code == 0:
			fg, bg, bold = 0, 0, false // reset all attributes
		case code == 1:
			bold = true
		case code >= 30 && code <= 37:
			fg = code - 30
		case code >= 40 && code <= 47:
			bg = code - 40
		// extended colour spaces or 256bit mode essentially
		// basically if the ANSI sequence is [38;5;123m, the code processes it as follows:
		// code == 38 (indicating an extended foreground color).
		// sequence = [38, 5, 123]
		// The condition len(sequence) >= 3 && sequence[1] == 5 is true.
		// fg = 123, meaning the foreground color is set to color 123 in the 256-color palette.
		case code == 38 || code == 48:
			if len(sequence) >= 3 && sequence[1] == 5 {
				if code == 38 {
					fg = sequence[2]
				} else {
					bg = sequence[2]
				}
			}
		}
	}

	return fg, bg, bold, newStart, nil
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}
