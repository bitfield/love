package main

import (
	"books"
	"fmt"
)

func main() {
	client := books.NewClient("localhost:3000")
	bookList, err := client.GetAllBooks()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, book := range bookList {
		fmt.Println(book)
	}
}
