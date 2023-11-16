package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your age: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		age, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		fmt.Printf("You entered: %d\n", age)
		break
	}
}
