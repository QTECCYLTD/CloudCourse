package main

import (
	"fmt"
	"time"
)

// factorialPart calculates a part of the factorial and sends the result to a channel
func factorialPart(start, end int, ch chan int) {
	result := 1
	for i := start; i <= end; i++ {
		result *= i
	}
	ch <- result
}

func main() {
	// Number to calculate the factorial of
	var n int = 10

	// Number of goroutines to use
	var numGoroutines int = 4

	// Channel to collect results
	ch := make(chan int, numGoroutines)

	// Calculate the workload for each goroutine
	workload := n / numGoroutines

	// Start timing
	startTime := time.Now()

	// Start goroutines
	for i := 0; i < numGoroutines; i++ {
		start := i*workload + 1
		end := (i + 1) * workload
		if i == numGoroutines-1 {
			end = n // Ensure the last goroutine goes up to n
		}
		go factorialPart(start, end, ch)
	}

	// Combine results
	total := 1
	for i := 0; i < numGoroutines; i++ {
		total *= <-ch
	}

	// Stop timing
	elapsedTime := time.Since(startTime)

	// Print results
	fmt.Printf("Factorial of %d is %d\n", n, total)
	fmt.Printf("Execution Time: %s\n", elapsedTime)
}
