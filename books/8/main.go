package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func main() {
	TestBookToString_FormatsBookInfoAsString()
	fmt.Println("It's all good!")
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v - %v copies",
		book.Title, book.Author, book.Copies)
}

func TestBookToString_FormatsBookInfoAsString() {
	input := Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	want := "Sea Room by Adam Nicolson - 2 copies"
	got := BookToString(input)
	if want != got {
		panic("BookToString: unexpected result")
	}
}
