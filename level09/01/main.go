package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go goroutine1()

	wg.Add(1)
	go goroutine2()

	wg.Wait()
}

func goroutine1() {
	fmt.Println("Go routine #1")
	wg.Done()
}

func goroutine2() {
	fmt.Println("Go routine #2")
	wg.Done()
}
