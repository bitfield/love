package main

import (
	"slices"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	want := []Book{
		{
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
		},
		{
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
		},
	}
	got := GetAllBooks()
	if !slices.Equal(want, got) {
		t.Errorf("want %#v, got %#v", want, got)
	}
}
