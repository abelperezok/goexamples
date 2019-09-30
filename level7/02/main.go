package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	p.name = "aaaaa"
	p.age = 33
}

func main() {

	p1 := person{
		name: "p1",
		age:  23,
	}

	fmt.Println(p1)

	changeMe(&p1)

	fmt.Println(p1)
}
