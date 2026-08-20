package books

import (
	"encoding/json/v2"
	"fmt"
	"maps"
	"os"
	"slices"
)

type Book struct {
	Title  string
	Author string
	Copies int
	ID     string
}

func (book Book) String() string {
	return fmt.Sprintf("%v by %v (copies: %v)",
		book.Title, book.Author, book.Copies)
}

func (book *Book) SetCopies(copies int) error {
	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	book.Copies = copies
	return nil
}

type Catalog map[string]Book

func OpenCatalog(path string) (*Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	catalog := &Catalog{}
	err = json.UnmarshalRead(file, catalog)
	if err != nil {
		return nil, err
	}
	return catalog, nil
}

func (catalog *Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(*catalog))
}

func (catalog *Catalog) GetBook(id string) (Book, bool) {
	book, ok := (*catalog)[id]
	return book, ok
}

func (catalog *Catalog) AddBook(book Book) error {
	_, ok := (*catalog)[book.ID]
	if ok {
		return fmt.Errorf("ID %q already exists", book.ID)
	}
	(*catalog)[book.ID] = book
	return nil
}

func (catalog *Catalog) SetCopies(id string, copies int) error {
	book, ok := (*catalog)[id]
	if !ok {
		return fmt.Errorf("ID %q not found", id)
	}
	err := book.SetCopies(copies)
	if err != nil {
		return err
	}
	(*catalog)[id] = book
	return nil
}

func (catalog *Catalog) Sync(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.MarshalWrite(file, catalog)
	if err != nil {
		return err
	}
	return nil
}
