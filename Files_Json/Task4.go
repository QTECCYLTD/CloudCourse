package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type Book struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	PublicationYear int    `json:"publication_year"`
}

func readJSONFile(filePath string) ([]Book, error) {
	file, err := ioutil.ReadFile("info.json")
	if err != nil {
		return nil, err
	}

	var books []Book
	err = json.Unmarshal(file, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func filterBooksByYear(books []Book, year int) []Book {
	var filteredBooks []Book
	for _, book := range books {
		if book.PublicationYear > year {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return filteredBooks
}

func sortBooksByTitle(books []Book) {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Title < books[j].Title
	})
}

func markClassics(books []Book) []Book {
	for i, book := range books {
		if book.PublicationYear < 1950 {
			books[i].Title += " (Classic)"
		}
	}
	return books
}

// writeJSONFile writes the slice of Book back into a JSON file
func writeJSONFile(filePath string, books []Book) error {
	file, err := json.MarshalIndent(books, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, file, 0644)
}

func main() {
	filePath := "data.json"
	books, err := readJSONFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Perform operations
	filteredBooks := filterBooksByYear(books, 1950)
	sortBooksByTitle(filteredBooks)
	modifiedBooks := markClassics(filteredBooks)

	// Write to a new file
	newFilePath := "processed_data.json"
	err = writeJSONFile(newFilePath, modifiedBooks)
	if err != nil {
		log.Fatalf("Failed to write JSON file: %v", err)
	}

	fmt.Println("Processed data written to", newFilePath)
}
