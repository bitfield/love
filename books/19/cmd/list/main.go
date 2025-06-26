package main

import (
	"books"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:3000/v1/list")
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status %d", resp.StatusCode)
		return
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	bookList := []books.Book{}
	err = json.Unmarshal(data, &bookList)
	if err != nil {
		fmt.Printf("%v in %q", err, data)
		return
	}
	for _, book := range bookList {
		fmt.Println(book)
	}
}
