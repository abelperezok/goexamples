package main

import (
	"fmt"
)

type cat int

var x cat

func main() {

	fmt.Printf("%T %v\n", x, x)

	x = 42

	fmt.Printf("%T %v\n", x, x)

}