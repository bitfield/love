package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:3000", http.HandlerFunc(hello))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world")
}
