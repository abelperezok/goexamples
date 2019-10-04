package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	const goroutines = 10

	for gr := 0; gr < goroutines; gr++ {
		go func(g int) {
			for i := 0; i < 10; i++ {
				num := g*10 + i
				fmt.Println("adding to channel", num)
				c <- num
			}
		}(gr)
	}

	for i := 0; i < goroutines*10; i++ {
		fmt.Println("pulling from channel", <-c)
	}
}
