package main

import (
	"fmt"
)

func sumfunc() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func main() {

	add := sumfunc()

	fmt.Println(add(1, 2))
}
