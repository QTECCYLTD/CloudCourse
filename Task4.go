package main

import "fmt"

// factorial function calculates the factorial of n
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var num int
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Please enter a positive integer.")
		return
	}

	result := factorial(num)
	fmt.Printf("Factorial of %d is %d\n", num, result)
}
