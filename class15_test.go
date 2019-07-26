package studygolang

import (
	"fmt"
	"math"
)

// interface

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r *rect) perim() float64 {
	return 2*r.height + 2*r.width
}

func (r *rect) area() float64 {
	return r.width * r.height
}

type circle struct {
	redius float
}

func (c *circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (c *circle) perim() float64 {
	return 2 * c.redius * math.Pi
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func TestFifteen(t *testing.T) {
	r := rect{width: 12.1221, height: 212.311}
	measure(r)
}
