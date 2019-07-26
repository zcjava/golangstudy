package studygolang

import (
	"fmt"
	"testing"
)

// map practice

func TestEight(t *testing.T) {
	m := make(map[int]string)
	m[1] = "11"
	m[2] = "22"
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)

	_, value := m[3]
	fmt.Println(value)
}
