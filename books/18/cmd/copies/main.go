package main

import (
	"books"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: copies <BOOK ID> <NUMBER OF COPIES>")
		return
	}
	catalog, err := books.OpenCatalog("testdata/catalog")
	if err != nil {
		fmt.Printf("opening catalog: %v\n", err)
		return
	}
	ID := os.Args[1]
	copies, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	err = catalog.SetCopies(ID, copies)
	if err != nil {
		fmt.Printf("updating book: %v\n", err)
		return
	}
	err = catalog.Sync()
	if err != nil {
		fmt.Printf("writing catalog: %v\n", err)
		return
	}
	fmt.Printf("Updated book %v to %d copies\n", ID, copies)
}
