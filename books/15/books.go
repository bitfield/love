package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     string
}

func (b Book) String() string {
	return fmt.Sprintf("%v by %v (copies: %v)",
		b.Title, b.Author, b.Copies)
}

type Catalog map[string]Book

func (catalog Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(catalog))
}

func (catalog Catalog) GetBook(ID string) (Book, bool) {
	book, ok := catalog[ID]
	return book, ok
}

func (catalog Catalog) AddBook(book Book) {
	catalog[book.ID] = book
}

func GetCatalog() Catalog {
	return Catalog{
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
