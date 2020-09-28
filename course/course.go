package course

import "fmt"

// Course type
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Lessons map[uint]string
}

// PrintLessons of a course
func (c *Course) PrintLessons() {
	text := "The lessons are: "
	for _, value := range c.Lessons {
		text += value + ", "
	}

	fmt.Println(text[:len(text)-2])
}

func (c *Course) updatePrice(p float64) {
	c.Price = p
}
