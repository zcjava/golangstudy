package studygolang

import (
	"fmt"
	"testing"
)

// func practice
func plus(nums ...int) (result int) {
	fmt.Println(nums)
	for _, num := range nums {
		fmt.Println(num)
		result += num
	}
	return result
}

func division(x int, y int) (result int, error string) {
	if y == 0 {
		error = "divisor must be greate than 0"
	} else {
		result = x / y
	}
	return
}

func TestTen(t *testing.T) {
	fmt.Println(plus(2, 3, 5, 5))
	fmt.Println(division(10, 0))
	fmt.Println(division(10, 2))
}
