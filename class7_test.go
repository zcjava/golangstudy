package studygolang

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

// slice practice

func TestSeven(t *testing.T) {
	s := make([]string, 3)
	for i := 0; i < 10; i++ {
		//s[i] = rand.Intn(100)
		s = append(s, strconv.Itoa(rand.Intn(100)))
	}
	fmt.Println(s)
	ss := s[3:]
	fmt.Println(ss)
}
