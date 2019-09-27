package main

import (
	"fmt"
)

func main() {

	number := 7

	fmt.Printf("%v\t%b\t%X\n", number, number, number)

	shifted := number << 1

	fmt.Printf("%v\t%b\t%X\n", shifted, shifted, shifted)
}
