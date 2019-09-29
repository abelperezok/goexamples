package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck1 := truck{
		vehicle: vehicle{
			doors:  2,
			colour: "white",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors:  4,
			colour: "red",
		},
		luxury: true,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1)

	fmt.Println(truck1.colour)
	fmt.Println(truck1.doors)
	fmt.Println(sedan1.colour)
	fmt.Println(sedan1.doors)
}
