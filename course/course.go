package course

import "fmt"

// Course type
type course struct {
	name    string
	price   float64
	isFree  bool
	userIDs []uint
	lessons map[uint]string
}

func (c *course) Name() string        { return c.name }
func (c *course) SetName(name string) { c.name = name }

func (c *course) Price() float64         { return c.price }
func (c *course) SetPrice(price float64) { c.price = price }

func (c *course) IsFree() bool          { return c.isFree }
func (c *course) SetIsFree(isFree bool) { c.isFree = isFree }

func (c *course) UserIDs() []uint       { return c.userIDs }
func (c *course) SetUserIDs(IDs []uint) { c.userIDs = IDs }

func (c *course) Lessons() map[uint]string           { return c.lessons }
func (c *course) SetLessons(lessons map[uint]string) { c.lessons = lessons }

// New course
func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 24.99
	}

	return &course{
		name:   name,
		price:  price,
		isFree: isFree,
	}
}

// PrintLessons of a course
func (c *course) PrintLessons() {
	text := "The lessons are: "
	for _, value := range c.lessons {
		text += value + ", "
	}

	fmt.Println(text[:len(text)-2])
}
