package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1985
	for {

		if i > time.Now().Year() {
			break
		}

		fmt.Println(i)
		i++
	}
}
