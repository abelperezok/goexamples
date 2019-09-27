package main

import (
	"fmt"
)

type cat int

var x cat
var y int

func main() {

	fmt.Printf("%T %v\n", x, x)

	x = 42

	fmt.Printf("%T %v\n", x, x)

	y = int(x)

	fmt.Printf("%T %v\n", y, y)

}