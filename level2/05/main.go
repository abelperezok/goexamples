package main

import (
	"fmt"
)

func main() {

	str := `raw string literal \n \t"more things" in a new line
	this is a true new line `

	fmt.Println(str)

	str = "not-raw string literal \n \tmore things in a new line\nthis is a true new line"

	fmt.Println(str)
}
