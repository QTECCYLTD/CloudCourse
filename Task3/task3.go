package main

import "fmt"

func main() {

	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Println("You are", age, "years old.")

	if age%2 == 0 {
		fmt.Println("Your age is even.")
	} else {
		fmt.Println("Your age is odd.")

	}

}
