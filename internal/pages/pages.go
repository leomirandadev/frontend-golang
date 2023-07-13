package pages

import (
	"net/http"

	"github.com/leomirandadev/frontend/internal/pages/home"
)

type PagesList struct {
	Home func(w http.ResponseWriter, r *http.Request)
}

var List = PagesList{
	Home: home.Render,
}
