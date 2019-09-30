package main

import "fmt"

func deferred() {
	fmt.Println("This function is deferred")
}

func main() {

	fmt.Println("first")
	defer deferred()
	fmt.Println("last")
}
