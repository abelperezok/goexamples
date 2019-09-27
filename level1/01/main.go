package main

import (
	"fmt"
)

func main() {

	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)

	fmt.Printf("%T %v\n%T %v\n%T %v\n", x, x, y, y, z, z)
}