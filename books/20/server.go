package books

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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
			http.Error(w, fmt.Sprintf("%q not found", ID), http.StatusNotFound)
			return
		}
		err := json.NewEncoder(w).Encode(book)
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/v1/getcopies/{id}", func(w http.ResponseWriter, r *http.Request) {
		ID := r.PathValue("id")
		copies, err := catalog.GetCopies(ID)
		if err != nil {
			http.Error(w, fmt.Sprintf("%q not found", ID), http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(copies)
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/v1/addcopies/{id}/{copies}", func(w http.ResponseWriter, r *http.Request) {
		ID := r.PathValue("id")
		copies, err := strconv.Atoi(r.PathValue("copies"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		stock, err := catalog.AddCopies(ID, copies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = catalog.Sync()
		if err != nil {
			panic(err)
		}
		err = json.NewEncoder(w).Encode(stock)
		if err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/v1/subcopies/{id}/{copies}", func(w http.ResponseWriter, r *http.Request) {
		ID := r.PathValue("id")
		copies, err := strconv.Atoi(r.PathValue("copies"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		stock, err := catalog.SubCopies(ID, copies)
		if err != nil {
			if errors.Is(err, ErrNotEnoughStock) {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, err.Error(), http.StatusNotFound)
			}
			return
		}
		err = catalog.Sync()
		if err != nil {
			panic(err)
		}
		err = json.NewEncoder(w).Encode(stock)
		if err != nil {
			panic(err)
		}
	})
	return http.ListenAndServe(addr, mux)
}
