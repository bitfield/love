package books

import (
	"encoding/json/v2"
	"fmt"
	"net/http"
)

func ListenAndServe(addr string, catalog *Catalog) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/list", func(w http.ResponseWriter, r *http.Request) {
		books := catalog.GetAllBooks()
		err := json.MarshalWrite(w, books)
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/v1/find/{id}", func(w http.ResponseWriter, r *http.Request) {
		ID := r.PathValue("id")
		book, ok := catalog.GetBook(ID)
		if !ok {
			http.Error(w, fmt.Sprintf("%q not found", ID), http.StatusNotFound)
			return
		}
		err := json.MarshalWrite(w, book)
		if err != nil {
			panic(err)
		}
	})
	return http.ListenAndServe(addr, mux)
}
