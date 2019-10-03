package main

import (
	"fmt"
	"sync"
)

var incrementer int = 0

var wg sync.WaitGroup
var mu sync.Mutex

func increment() {
	mu.Lock()
	i := incrementer
	fmt.Println("initial incrementer", i)
	//runtime.Gosched()
	i++
	incrementer = i
	fmt.Println("new incrementer", i)
	mu.Unlock()
	wg.Done()
}

func main() {

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go increment()
	}

	wg.Wait()
}
