package main

import (
	"fmt"
)

func main() {

	a := [5]int{1, 2, 3, 4, 5}

	for k, v := range a {
		fmt.Printf("index %v value %v (%T)\n", k, v, v)
	}

	fmt.Printf("array type %T\n", a)
}
