package main

import (
	"fmt"
	"strings"
)

func main() {

	sport := "Football"

	switch favSport := strings.ToUpper(sport); favSport {
	case "FOOTBALL":
		fmt.Println(sport, "is great (", favSport, ")")
	case "BASEBALL":
		fmt.Println(sport, "is better(", favSport, ")")
	default:
		fmt.Println(sport, "is not a good choice :)")
	}

}
