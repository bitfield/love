package main

import (
	"books"
	"fmt"
)

func main() {
	catalog := books.GetCatalog()
	for _, book := range books.GetAllBooks(catalog) {
		fmt.Println(books.BookToString(book))
	}
}
