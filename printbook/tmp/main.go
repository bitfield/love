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
	fmt.Println("It's all good!")
}
