package main

import "fmt"

const (
	_     = iota
	next1 = 2019 + iota
	next2
	next3
	next4
)

func main() {
	fmt.Println(next1)
	fmt.Println(next2)
	fmt.Println(next3)
	fmt.Println(next4)
}
