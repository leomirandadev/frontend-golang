package home

import (
	"embed"
	"net/http"
	"text/template"
)

//go:embed home.html
var asset embed.FS

type Content struct {
	Username string
}

func Render(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(asset, "home.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	username := r.URL.Query().Get("username")

	if len(username) == 0 {
		username = "user not found"
	}

	err = tmpl.Execute(w, Content{Username: username})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
