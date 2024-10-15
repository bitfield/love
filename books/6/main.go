package main

import "fmt"

func main() {
	fmt.Println("Books in stock:")
	var title = "The City & The City"
	var author = "China Miéville"
	var copies = 1
	printBook(title, author, copies)
	title = "The Sugar Barons"
	author = "Matthew Parker"
	copies = 2
	printBook(title, author, copies)
}

func printBook(title, author string, copies int) {
	fmt.Println(title, "by", author, "-", copies, "copies")
}
