package main

import (
	"books"
	"fmt"
)

func main() {
	for _, book := range books.GetAllBooks() {
		fmt.Println(books.BookToString(book))
	}
}
