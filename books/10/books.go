package books

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func BookToString(b Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)",
		b.Title, b.Author, b.Copies)
}
