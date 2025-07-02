package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:3000")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(resp.StatusCode)
	}
	msg, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", msg)
}
