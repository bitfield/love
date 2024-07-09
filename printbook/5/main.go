package main

import "fmt"

func main() {
	fmt.Println("Books in stock:")

	var title = "The City & The City"
	var author = "China Miéville"
	printBook(title, author)
	title = "The Sugar Barons"
	author = "Matthew Parker"
	printBook(title, author)
}

func printBook(title, author string) {
	fmt.Println(title, "by", author)
}
