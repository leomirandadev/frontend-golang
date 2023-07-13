package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/leomirandadev/frontend/internal/pages"
)

func Init(r *chi.Mux, pagesList pages.PagesList) {
	r.Get("/", pagesList.Home)
	r.Get("/home", pagesList.Home)
}
