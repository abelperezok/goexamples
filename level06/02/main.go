package main

import "fmt"

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	xfoo := []int{1, 2, 3, 4, 5}
	xbar := []int{6, 7, 8, 9, 10}

	foosum := foo(xfoo...)
	barsum := bar(xbar)

	fmt.Println(foosum)
	fmt.Println(barsum)
}
