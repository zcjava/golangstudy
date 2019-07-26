package studygolang

import (
	"fmt"
	"testing"
)

// close package feature

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

var fnum int = 1

func factorial(x int) {
	if x > 0 {
		fnum *= x
		factorial(x - 1)
	}

}

func TestEleven(t *testing.T) {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt = intSeq()
	fmt.Println(nextInt())
	factorial(7)
	fmt.Println(fnum)
}
