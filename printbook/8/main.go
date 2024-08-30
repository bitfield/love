package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func main() {
	TestBookToStringReturnsExpectedValue()
	fmt.Println("It's all good!")
}

func bookToString(b Book) string {
	return fmt.Sprintf("%v by %v - %v copies", b.Title, b.Author, b.Copies)
}

func TestBookToStringReturnsExpectedValue() {
	input := Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	want := "Sea Room by Adam Nicolson - 2 copies"
	got := bookToString(input)
	if want != got {
		panic("bookToString: unexpected result")
	}
}
