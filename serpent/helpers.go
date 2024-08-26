package serpent

import "strings"

// func (s *Snake) updateTermContent(y, x int) {
// 	actualX := 0
// 	inEscapeSeq := false
// 	for i := range s.TermContent[y] {
// 		if s.TermContent[y][i] == '\033' {
// 			inEscapeSeq = true
// 		} else if !inEscapeSeq {
// 			// this conditional ensures we can't override the letters that are
// 			// part of the escape sequences like the m in \033[36m which is
// 			// essentially part of the escape sequence and not the underlying text
// 			if actualX == x {
// 				s.TermContent[y][i] = ' '
// 				return
// 			}
// 			actualX++
// 		} else if (s.TermContent[y][i] >= 'A' && s.TermContent[y][i] <= 'Z') ||
// 			(s.TermContent[y][i] >= 'a' && s.TermContent[y][i] <= 'z') {
// 			inEscapeSeq = false
// 		}
// 	}
// }

// basically run through the TermContent while checking if the character we're
// pointing to is part of an escape sequence
func (s *Snake) updateTermContent(y, x int) {
	actualX := 0
	inEscapeSeq := false
	var currentColor string
	var escapeSeq strings.Builder

	for i := range s.TermContent[y] {
		if s.TermContent[y][i] == '\033' {
			inEscapeSeq = true
			escapeSeq.Reset()
			escapeSeq.WriteRune(s.TermContent[y][i])
		} else if inEscapeSeq {
			escapeSeq.WriteRune(s.TermContent[y][i])
			if isLetter(s.TermContent[y][i]) {
				inEscapeSeq = false
				if strings.Contains(escapeSeq.String(), "[3") || strings.Contains(escapeSeq.String(), "[38;5;") {
					currentColor = escapeSeq.String()
				}
			}
		} else {
			if actualX == x {
				s.TermContent[y][i] = ' '
				if currentColor != "" {
					s.colour = currentColor
				}
				return
			}
			actualX++
		}
	}
}

func isLetter(char rune) bool {
	return char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z'
}

// the idea is to make keep two buffers; one that is actually printed and another
// for the snake to eat that way we don't mess up any ansi escaped sequences
func stripAnsiCodes(rawContent [][]rune) [][]rune {
	filteredContent := make([][]rune, len(rawContent))
	inEscapeSeq := false

	for i, line := range rawContent {
		filteredLine := make([]rune, 0, len(line))
		for _, c := range line {
			if c == '\033' {
				inEscapeSeq = true
				continue
			}
			if inEscapeSeq {
				if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
					inEscapeSeq = false
				}
				continue
			}
			filteredLine = append(filteredLine, c)
		}
		filteredContent[i] = filteredLine
	}
	return filteredContent
}
