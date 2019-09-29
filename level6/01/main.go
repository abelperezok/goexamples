package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "Universal Answer"
}

func main() {

	x := foo()
	fmt.Println(x)

	y, z := bar()
	fmt.Println(y, z)
}
