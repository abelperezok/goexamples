package main

import "fmt"

type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v and I am %v years old", p.first, p.last, p.age)
}

func main() {

	p := person{
		first: "John",
		last:  "Doe",
		age:   34,
	}

	p.speak()
}
