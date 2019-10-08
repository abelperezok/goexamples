package main

import "fmt"

type person struct {
	firstName                 string
	lastName                  string
	favouriteIceCreamFlavours []string
}

func main() {

	p1 := person{
		firstName:                 "Name1",
		lastName:                  "Surname1",
		favouriteIceCreamFlavours: []string{"flavour1", "flavour2"},
	}

	p2 := person{
		firstName:                 "Name2",
		lastName:                  "Surname2",
		favouriteIceCreamFlavours: []string{"flavour3", "flavour4", "flavour5"},
	}

	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.favouriteIceCreamFlavours {
		fmt.Println("\t", v)
	}

	fmt.Println(p1.firstName, p2.lastName)
	for _, v := range p2.favouriteIceCreamFlavours {
		fmt.Println("\t", v)
	}
}
