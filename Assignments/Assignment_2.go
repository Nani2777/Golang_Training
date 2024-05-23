/* package main

import (
	"fmt"
)

func producer(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("Data Read:", num)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
	var input string
	fmt.Scanln(&input)
	// close(ch)
}
 */

 package main

 import (
	 "fmt"
	 "sync"
 )
 
 func producer(ch chan<- int, wg *sync.WaitGroup) {
	 defer wg.Done()
	 for i := 1; i <= 100; i++ {
		 ch <- i // Send value to channel
	 }
	 close(ch) // Close the channel when done
 }
 
 func consumer(ch <-chan int, wg *sync.WaitGroup) {
	 defer wg.Done()
	 for num := range ch { // Iterate over channel until it's closed
		 fmt.Println("Consumed:", num)
	 }
 }
 
 func main() {
	 // Create a channel of type int
	 ch := make(chan int)
 
	 // Create a WaitGroup
	 var wg sync.WaitGroup
 
	 // Add two goroutines to the WaitGroup
	 wg.Add(2)
 
	 // Launch producer goroutine
	 go producer(ch, &wg)
 
	 // Launch consumer goroutine
	 go consumer(ch, &wg)
 
	 // Goroutine to listen for exit signal
	 done := make(chan struct{})
	 go func() {
		 defer close(done)
		 wg.Wait() // Wait for all goroutines to finish
	 }()
 
	 <-done // Block until all goroutines are done
 
	 fmt.Println("All goroutines finished. Exiting...")
 }
 