package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(a[1:5])  // starting at position 1 (second) and ending at position 4 (fifth)
	fmt.Println(a[0:])   // the whole slice
	fmt.Println(a[:10])  // the whole slice
	fmt.Println(a[7:10]) // last 3 items
	fmt.Println(a[0:3])  // first 3 items
	fmt.Println(a[4:5])  // the 5th item only
}
