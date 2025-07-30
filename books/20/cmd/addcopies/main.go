package main

import (
	"books"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: addcopies <BOOK ID> <HOW MANY>")
		return
	}
	ID := os.Args[1]
	copies, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	client := books.NewClient("localhost:3000")
	stock, err := client.AddCopies(ID, copies)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d copies in stock", stock)
}
