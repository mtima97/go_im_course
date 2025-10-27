package main

import (
	"fmt"
	"gocourse/internal/service/files"
	"io"
	"log"
	"net/http"
)

//goland:noinspection ALL
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text, err := files.Read()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, text)
	})

	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer r.Body.Close()

		fmt.Fprintln(w, string(body))
	})

	// Make GET
	// Make POST

	if err := http.ListenAndServe(":7999", mux); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
