package main

import (
	"github.com/tavo1987/oop/course"
)

func main() {
	golang := course.New("Golang from scratch", 25, false)
	golang.SetUserIDs([]uint{1, 6, 3, 5})
	golang.SetLessons(map[uint]string{
		1: "Introduction",
		2: "Structures",
		3: "Maps",
	})

	golang.SetPrice(65)

	golang.PrintLessons()
}
