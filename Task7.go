package main

import (
	"fmt"
)

// Book represents a book with title, author, and publication year.
type Book struct {
	title           string
	author          string
	publicationYear int
}

// DisplayDetails is a method for the Book struct that prints its details.
func (b Book) DisplayDetails() {
	fmt.Printf("Title: %s\n", b.title)
	fmt.Printf("Author: %s\n", b.author)
	fmt.Printf("Publication Year: %d\n", b.publicationYear)
	fmt.Println("-----------------------")
}

func main() {
	// Creating instances of the Book struct
	book1 := Book{
		title:           "The Great Gatsby",
		author:          "F. Scott Fitzgerald",
		publicationYear: 1925,
	}

	book2 := Book{
		title:           "Moby-Dick",
		author:          "Herman Melville",
		publicationYear: 1851,
	}

	// Calling the DisplayDetails method for each book
	book1.DisplayDetails()
	book2.DisplayDetails()
}
