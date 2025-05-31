package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"net/http"
	"os"
	"slices"
	"sync"
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

type Catalog struct {
	mu   *sync.RWMutex
	data map[string]Book
}

func NewCatalog() *Catalog {
	return &Catalog{
		mu:   &sync.RWMutex{},
		data: map[string]Book{},
	}
}

func OpenCatalog(path string) (*Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	catalog := Catalog{
		mu:   &sync.RWMutex{},
		data: map[string]Book{},
	}
	err = json.NewDecoder(file).Decode(&catalog.data)
	if err != nil {
		return nil, err
	}
	return &catalog, nil
}

func (catalog *Catalog) GetAllBooks() []Book {
	catalog.mu.RLock()
	defer catalog.mu.RUnlock()
	return slices.Collect(maps.Values(catalog.data))
}

func (catalog *Catalog) GetBook(ID string) (Book, bool) {
	catalog.mu.RLock()
	defer catalog.mu.RUnlock()
	book, ok := catalog.data[ID]
	return book, ok
}

func (catalog *Catalog) GetCopies(ID string) (int, error) {
	catalog.mu.RLock()
	defer catalog.mu.RUnlock()
	book, ok := catalog.data[ID]
	if !ok {
		return 0, fmt.Errorf("ID %q not found", ID)
	}
	return book.Copies, nil
}

func (catalog *Catalog) Sync(path string) error {
	catalog.mu.RLock()
	defer catalog.mu.RUnlock()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(catalog.data)
	if err != nil {
		return err
	}
	return nil
}

func (catalog *Catalog) AddBook(book Book) error {
	catalog.mu.Lock()
	defer catalog.mu.Unlock()
	_, ok := catalog.data[book.ID]
	if ok {
		return fmt.Errorf("ID %q already exists", book.ID)
	}
	catalog.data[book.ID] = book
	return nil
}

func (catalog *Catalog) SetCopies(ID string, copies int) error {
	catalog.mu.Lock()
	defer catalog.mu.Unlock()
	book, ok := catalog.data[ID]
	if !ok {
		return fmt.Errorf("ID %q not found", ID)
	}
	err := book.SetCopies(copies)
	if err != nil {
		return err
	}
	catalog.data[ID] = book
	return nil
}

type APIServer struct {
	Addr       string
	Catalog    *Catalog
	httpServer *http.Server
}

func NewAPIServer(addr string, catalog *Catalog) *APIServer {
	return &APIServer{
		Addr:    addr,
		Catalog: catalog,
	}
}

func (srv *APIServer) ListenAndServe() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/books", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		books := srv.Catalog.GetAllBooks()
		err := json.NewEncoder(w).Encode(books)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	srv.httpServer = &http.Server{
		Addr:    srv.Addr,
		Handler: mux,
	}
	return srv.httpServer.ListenAndServe()
}

func (srv *APIServer) Close() {}
