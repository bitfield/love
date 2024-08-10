package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

var books = []Book{
	{
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
	},
	{
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
	},
}

func GetAllBooks() []Book {
	return books
}

func main() {
	fmt.Println(GetAllBooks())
}
