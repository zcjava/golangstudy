package studygolang

import (
	"fmt"
	"testing"
)

// method of struct

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func TestFourTeen(t *testing.T) {
	rt := rect{width: 1, height: 2}
	fmt.Println(rt.area())
	fmt.Println(rt.perim())
}
