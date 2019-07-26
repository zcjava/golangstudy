package studygolang

import (
	"fmt"
	"testing"
)

//pointer practice

func zeroval(x int) {
	x = 0
	fmt.Println("zeroval x address:", &x)
}

func zeroptr(x *int) {
	*x = 0
}

func TestTwelve(t *testing.T) {
	i := 1
	fmt.Println("initial:", i, " address:", &i)

	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i, "  addr:", &i)

}
