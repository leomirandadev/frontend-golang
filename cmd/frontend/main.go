package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leomirandadev/frontend/internal/handlers"
	"github.com/leomirandadev/frontend/internal/pages"
)

func main() {
	r := chi.NewRouter()

	handlers.Init(r, pages.List)

	log.Fatal(http.ListenAndServe(":3000", r))
}
