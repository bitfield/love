package main

import (
	"books"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: find <BOOK ID>")
		return
	}
	catalog := books.GetCatalog()
	ID := os.Args[1]
	book, ok := catalog.GetBook(ID)
	if !ok {
		fmt.Println("Sorry, I couldn't find that book in the catalog.")
		return
	}
	fmt.Println(book)
}
