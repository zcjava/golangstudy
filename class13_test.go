package studygolang

import (
	"fmt"
	"testing"
)

// struct

type person struct {
	name string
	age  int
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func TestThirteen(t *testing.T) {
	p := person{name: "zhangsan", age: 11}
	fmt.Println(p, " address:", &p)
	s := &p
	i := 1
	fmt.Println(&s)
	fmt.Println(typeof(s))
	fmt.Println("i address:", &i)

}
