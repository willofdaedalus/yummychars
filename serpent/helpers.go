package serpent

func stripAnsiCodes(rawContent [][]rune) [][]rune {
	filteredContent := make([][]rune, len(rawContent))
	for i, line := range rawContent {
		filteredLine := make([]rune, 0, len(line))
		inEscapeSeq := false
		for _, char := range line {
			if char == '\033' {
				inEscapeSeq = true
				continue
			}
			if inEscapeSeq {
				if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
					inEscapeSeq = false
				}
				continue
			}
			filteredLine = append(filteredLine, char)
		}
		filteredContent[i] = filteredLine
	}
	return filteredContent
}
