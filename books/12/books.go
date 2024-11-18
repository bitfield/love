package books

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
	ID     string
}

func BookToString(b Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)",
		b.Title, b.Author, b.Copies)
}

var catalog = []Book{
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

func GetAllBooks() []Book {
	return catalog
}

func GetBook(ID string) (Book, bool) {
	for _, book := range catalog {
		if book.ID == ID {
			return book, true
		}
	}
	return Book{}, false
}
