package main

import (
	"fmt"
)

func main() {

	exp1 := 0 == 1
	exp2 := 0 <= 1
	exp3 := 0 >= 1
	exp4 := 0 != 1
	exp5 := 0 < 1
	exp6 := 0 > 1

	fmt.Println("0 == 1", exp1)
	fmt.Println("0 <= 1", exp2)
	fmt.Println("0 >= 1", exp3)
	fmt.Println("0 != 1", exp4)
	fmt.Println("0 < 1", exp5)
	fmt.Println("0 > 1", exp6)

}
