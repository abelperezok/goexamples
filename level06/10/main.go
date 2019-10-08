package main

import (
	"fmt"
)

func powersof(factor int) func() int {
	value := 1
	return func() int {
		value *= factor
		return value
	}
}

func main() {

	p1 := powersof(1)

	fmt.Println("calling p1")
	fmt.Println(p1())
	fmt.Println(p1())
	fmt.Println(p1())
	fmt.Println(p1())

	p2 := powersof(2)

	fmt.Println("calling p2")
	fmt.Println(p2())
	fmt.Println(p2())
	fmt.Println(p2())
	fmt.Println(p2())

	p3 := powersof(3)

	fmt.Println("calling p3")
	fmt.Println(p3())
	fmt.Println(p3())
	fmt.Println(p3())
	fmt.Println(p3())

}
