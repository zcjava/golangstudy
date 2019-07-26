package studygolang

import (
	"fmt"
	"testing"
)

var a1, a2 int = 22, 33

func TestThree(t *testing.T) {
	a3 := "three"
	fmt.Println("ye", a1, a3, a2)
}
