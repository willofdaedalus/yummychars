package serpent

import (
	"reflect"
	"testing"
)

func TestStripAnsiCodes(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]rune
		expected [][]rune
	}{
		{
			name:     "Simple green text",
			input:    [][]rune{[]rune("\033[32mThis is green\033[0m")},
			expected: [][]rune{[]rune("This is green")},
		},
		{
			name:     "Background color 0",
			input:    [][]rune{[]rune("\033[48;5;0mColor 0\033[0m")},
			expected: [][]rune{[]rune("Color 0")},
		},
		{
			name:     "Background color 1",
			input:    [][]rune{[]rune("\033[48;5;1mColor 1\033[0m")},
			expected: [][]rune{[]rune("Color 1")},
		},
		{
			name:     "Multiple lines with different colors",
			input:    [][]rune{[]rune("\033[31mRed\033[0m"), []rune("\033[32mGreen\033[0m"), []rune("\033[34mBlue\033[0m")},
			expected: [][]rune{[]rune("Red"), []rune("Green"), []rune("Blue")},
		},
		{
			name:     "Text with no ANSI codes",
			input:    [][]rune{[]rune("Plain text")},
			expected: [][]rune{[]rune("Plain text")},
		},
		{
			name:     "Empty input",
			input:    [][]rune{},
			expected: [][]rune{},
		},
		{
			name:     "Multiple ANSI codes in one line",
			input:    [][]rune{[]rune("\033[1m\033[31mBold Red\033[0m \033[32mGreen\033[0m")},
			expected: [][]rune{[]rune("Bold Red Green")},
		},
	}

	// cheers to claude.ai
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stripAnsiCodes(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("stripAnsiCodes() = %v, want %v", result, tt.expected)
			}
		})
	}
}
