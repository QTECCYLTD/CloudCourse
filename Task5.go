package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Sum of slice elements: %d\n", sum)

	dictionary := make(map[string]int)

	dictionary["one"] = 1
	dictionary["two"] = 2
	dictionary["three"] = 3

	for key, value := range dictionary {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
