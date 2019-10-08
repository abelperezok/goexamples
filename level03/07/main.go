package main

import (
	"fmt"
)

func main() {

	password := "p455w0rd"

	if len(password) <= 0 {
		fmt.Println("password required")
	} else if len(password) > 20 {
		fmt.Println("password too long")
	} else {
		fmt.Println("password is good")
	}
}
