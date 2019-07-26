package studygolang

import (
	"fmt"
	"testing"
)

func TestFive(t *testing.T) {
	var i = 1
	for i >= 3 {
		fmt.Println(i)
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	for {

		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

	if num := 9; num < 0 {
		fmt.Println("num < 0")

	} else {
		fmt.Println("num > 0")
	}
}
