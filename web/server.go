package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Route("/user/{uniqID}", func(w http.ResponseWriter, r *http.Request) {
		uniqID := chi.URLParam(r, "estateID")

	})
}
