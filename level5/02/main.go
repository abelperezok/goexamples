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

	// m := map[string]person{}
	// m[p1.lastName] = p1
	// m[p2.lastName] = p2

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName, v.lastName)
		for _, f := range v.favouriteIceCreamFlavours {
			fmt.Println("\t", f)
		}
	}

}
