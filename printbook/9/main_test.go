package main

import "testing"

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := BookToString(input)
	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}
