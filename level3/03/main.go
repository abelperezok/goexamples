package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1985
	for i <= time.Now().Year() {

		fmt.Println(i)

		i++
	}
}
