package books

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ListenAndServe(addr string, catalog *Catalog) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/list", func(w http.ResponseWriter, r *http.Request) {
		books := catalog.GetAllBooks()
		err := json.NewEncoder(w).Encode(books)
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/v1/find/{id}", func(w http.ResponseWriter, r *http.Request) {
		ID := r.PathValue("id")
		book, ok := catalog.GetBook(ID)
		if !ok {
			panic(fmt.Sprintf("ID %q not found", ID))
		}
		err := json.NewEncoder(w).Encode(book)
		if err != nil {
			panic(err)
		}
	})
	return http.ListenAndServe(addr, mux)
}
