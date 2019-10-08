package main

import (
	"fmt"
)

func main() {

	password := "p455w0rd"

	switch {
	case len(password) <= 0:
		fmt.Println("password required")
	case len(password) > 20:
		fmt.Println("password too long")
	default:
		fmt.Println("password is good")
	}

}
