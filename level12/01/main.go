package main

import (
	"fmt"

	"github.com/abelperezok/goexamples/level12/01/cat"
	"github.com/abelperezok/goexamples/level12/01/dog"
)

func main() {

	h := 30
	d := dog.Year(h)
	fmt.Println(h, "human years are", d, "dog years")
	c := cat.Year(h)
	fmt.Println(h, "human years are", c, "cat years")
}
