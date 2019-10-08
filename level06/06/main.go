package main

import (
	"fmt"
)

func main() {

	result := func(x, y int) int {
		return x + y
	}(1, 2)

	fmt.Println(result)
}
