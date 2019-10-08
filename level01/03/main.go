package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%T %v\n%T \"%v\"\n%T %v\n", x, x, y, y, z, z)
	fmt.Printf(s)
}