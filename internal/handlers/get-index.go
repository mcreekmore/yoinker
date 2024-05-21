package handlers

import (
	"net/http"
	"text/template"

	"github.com/mcreekmore/yoinker/pkg/models"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]models.Film{
		"Films": {
			{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
			{Title: "Inglourious Basterds", Director: "Quentin Tarantino"},
			{Title: "Kill Bill", Director: "Quentin Tarantino"},
		},
	}
	tmpl.Execute(w, films)
}
