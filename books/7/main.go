package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func main() {
	fmt.Println("Books in stock:")
	b := Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	printBook(b)
	// Output:
	// Sea Room by Adam Nicolson - 2 copies
}

func printBook(b Book) {
	fmt.Println(b.Title, "by", b.Author, "-", b.Copies, "copies")
}
