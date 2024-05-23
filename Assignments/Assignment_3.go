package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	numGoroutines := 10
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter Value: ", counter)
}