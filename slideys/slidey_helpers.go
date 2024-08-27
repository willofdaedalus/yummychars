package slideys

import "fmt"

func (s *Slidey) splitRuneContent() ([][]rune, [][]rune) {
	r1 := make([][]rune, len(s.content))
	r2 := make([][]rune, len(s.content))

	for _, line := range s.content {
		fmt.Println(len(line))
		half := len(line) / 2
		rr1 := make([]rune, half)
		rr2 := make([]rune, half)

		for j := range line {
			if j <= half {
				rr1 = append(rr1, line[j])
			} else {
				rr2 = append(rr2, line[j])
			}

		}
		r1 = append(r1, rr1)
		r2 = append(r2, rr2)
	}

	return r1, r2
}
