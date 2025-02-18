package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func main() {
	fmt.Println("Books in stock:")
	book := Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	printBook(book)
	// Output:
	// Sea Room by Adam Nicolson - 2 copies
}

func printBook(book Book) {
	fmt.Println(book.Title, "by", book.Author, "-", book.Copies, "copies")
}
