package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var incrementer int64 = 0

var wg sync.WaitGroup

func increment() {

	atomic.AddInt64(&incrementer, 1)
	runtime.Gosched()
	fmt.Println("new incrementer", atomic.LoadInt64(&incrementer))
	wg.Done()
}

func main() {

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go increment()
	}

	wg.Wait()
}
