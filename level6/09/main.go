package main

import (
	"fmt"
)

func calculate(x, y int, f func(int, int) int) int {
	return f(x, y)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func main() {

	fmt.Println(calculate(1, 2, add))
	fmt.Println(calculate(1, 2, sub))
	fmt.Println(calculate(1, 2, mul))
}
