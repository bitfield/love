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

func main() {
	book := Book{
		Title:  "Engineering in Plain Sight",
		Author: "Grady Hillhouse",
		Copies: 2,
	}
	fmt.Println(bookToString(book))
}
