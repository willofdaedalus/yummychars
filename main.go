package main

import (
	"time"
	"willofdaedalus/yummychars/serpent"
)

func main() {
	s := serpent.InitSnake(1)

	for {
		s.MoveSnake(serpent.DOWN)
		s.DrawSnake()

		// add a short sleep to control the loop speed
		// this isn't the best but it works; might come back this
		time.Sleep(time.Second / time.Duration(s.Speed))
	}
}
