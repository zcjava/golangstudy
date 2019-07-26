package studygolang

import (
	"fmt"
	"math/rand"
	"testing"
)

var six [5]int

func TestSix(t *testing.T) {
	fmt.Println(six)
	six[4] = 1
	fmt.Println(six, "six len:", len(six))

	six1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(six1)
	var multi [3][4]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			multi[i][j] = rand.Intn(30)
		}
	}
	fmt.Println("multi arrays: ", multi)
}
