package main

import (
	"fmt"
	"math"
)

// Defining the shape interface with methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length, Width float64
}

// Area method for the Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Defining circle struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Permimeter ,ethod for the circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Function to calculate the area and  perimeter of the shape
func Calculate(shape Shape) {
	fmt.Printf("Area: %.2f\n", shape.Area())
	fmt.Printf("Perimeter: &.2f\n", shape.Perimeter())
}

func main() {
	r := Rectangle{Length: 10, Width: 5}
	c := Circle{Radius: 7}

	fmt.Println("Rectabgle: ")
	Calculate(r)

	fmt.Println("\nCircle:")
	Calculate(c)

}
