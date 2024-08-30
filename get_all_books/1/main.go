package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func bookToString(b Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", b.Title, b.Author, b.Copies)
}

var books = []Book{
	{
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	},
	{
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
	},
}

func GetAllBooks() []Book {
	return books
}

func main() {
	books := GetAllBooks()
	for _, book := range books {
		fmt.Println(bookToString(book))
	}
}
