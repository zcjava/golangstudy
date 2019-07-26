package studygolang

import (
	"fmt"
	"testing"
)

// range practice

func TestNine(t *testing.T) {
	arrs := []int{2, 3, 4}
	sum := 0
	for _, num := range arrs {
		sum += num
	}
	fmt.Println(arrs, " sum ", sum)

	m := make(map[int]string, 10)
	m[1] = "z"
	m[0] = "c"

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
