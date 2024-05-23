package main

import (
	"fmt"
)

func addition(first, last int, result chan<- int) {
	sum := 0
	for i := first; i <= last; i++ {
		sum += i
	}
	result <- sum
}

func main() {
	var N int
	fmt.Print("Enter the value of N: ")
	fmt.Scan(&N)
	result := make(chan int)
	num := 3
	step := N / num
	for i := 0; i < num; i++ {
		first := i*step + 1
		last := (i + 1) * step
		if i == num-1 {
			last = N
		}
		go addition(first, last, result)
	}
	totalSum := 0
	for i := 0; i < num; i++ {
		totalSum += <-result
	}

	fmt.Println("Sum from 1 to", N, "is:", totalSum)
}
