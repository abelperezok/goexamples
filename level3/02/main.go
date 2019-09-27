package main

import "fmt"

func main() {

	for letter := 65; letter < 65+26; letter++ {
		fmt.Println(letter)
		for i := 0; i < 3; i++ {
			//fmt.Printf("\t%U '%v'\n", letter, string(letter))
			fmt.Printf("\t%#U\n", letter)
		}
	}
}
