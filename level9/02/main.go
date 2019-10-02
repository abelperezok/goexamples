package main

import "fmt"

type person struct {
	name string
	age  int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am", p.name, p.age)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "p1",
		age:  34,
	}

	saySomething(&p)
	//saySomething(p) doesn't work as it's a pointer receiver
}
