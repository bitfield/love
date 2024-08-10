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
		},
		{
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
		},
	}
	got := GetAllBooks()
	if !slices.Equal(want, got) {
		t.Errorf("want %#v, got %#v", want, got)
	}
}
