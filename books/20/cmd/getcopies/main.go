package main

import (
	"books"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: getcopies <BOOK ID>")
		return
	}
	ID := os.Args[1]
	client := books.NewClient("localhost:3000")
	copies, err := client.GetCopies(ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d copies in stock", copies)
}
