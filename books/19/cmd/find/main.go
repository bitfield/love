package main

import (
	"books"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: find <BOOK ID>")
		return
	}
	ID := os.Args[1]
	resp, err := http.Get("http://localhost:3000/v1/find/" + ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status %d", resp.StatusCode)
		return
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	book := books.Book{}
	err = json.Unmarshal(data, &book)
	if err != nil {
		fmt.Printf("%v in %q", err, data)
		return
	}
	fmt.Println(book)
}
