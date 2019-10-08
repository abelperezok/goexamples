package main

import "fmt"

func main() {

	s := struct {
		name string
		age  int
	}{
		name: "abel",
		age:  34,
	}

	fmt.Println(s.name)
	fmt.Println(s.age)
}
