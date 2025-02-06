package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := []*books.Book{
		{
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
			ID:     "abc",
		},
		{
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
			ID:     "xyz",
		},
	}
	got := catalog.GetAllBooks()
	slices.SortFunc(got, func(a, b *books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})
	if !slices.EqualFunc(want, got, func(a, b *books.Book) bool {
		return *a == *b
	}) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_FindsBookInCatalogByID(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	want := &books.Book{
		ID:     "abc",
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	}
	got, ok := catalog.GetBook("abc")
	if !ok {
		t.Fatal("book not found")
	}
	if *want != *got {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("nonexistent ID")
	if ok {
		t.Fatal("want false for nonexistent ID, got true")
	}
}

func TestAddBook_AddsGivenBookToCatalog(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("123")
	if ok {
		t.Fatal("book already present")
	}
	catalog.AddBook(books.Book{
		ID:     "123",
		Title:  "The Prize of all the Oceans",
		Author: "Glyn Williams",
		Copies: 2,
	})
	_, ok = catalog.GetBook("123")
	if !ok {
		t.Fatal("added book not found")
	}
}

func TestChangeToNumberOfCopiesPersistsInCatalog(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	book, ok := catalog.GetBook("xyz")
	if !ok {
		t.Fatal("book not found before change")
	}
	if book.Copies != 2 {
		t.Fatalf("expected 2 copies before change, got %d", book.Copies)
	}
	book.Copies--
	book2, ok := catalog.GetBook("xyz")
	if !ok {
		t.Fatal("book not found after change")
	}
	if book2.Copies != 1 {
		t.Fatalf("expected 1 copy after change, got %d", book2.Copies)
	}
}

func getTestCatalog() books.Catalog {
	return books.Catalog{
		"abc": {
			Title:  "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
			ID:     "abc",
		},
		"xyz": {
			Title:  "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
			ID:     "xyz",
		},
	}
}
