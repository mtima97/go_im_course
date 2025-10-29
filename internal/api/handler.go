package api

import (
	"fmt"
	"net/http"
)

type Handler struct {
	//
}

//goland:noinspection ALL
func (h Handler) HandleForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	title := r.FormValue("title")

	fmt.Fprintf(w, "Title: %s", title)
}
