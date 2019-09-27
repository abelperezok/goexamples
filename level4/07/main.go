package main

import (
	"fmt"
)

func main() {

	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for k, v := range x {
		for kk, vv := range v {
			fmt.Println(k, kk, vv)
		}
	}
}
