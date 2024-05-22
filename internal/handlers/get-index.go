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
	tmpl := template.Must(template.ParseFiles("./internal/templates/index.html"))
	films := map[string][]models.DownloadResponse{
		"DownloadResponses": {},
	}
	tmpl.Execute(w, films)
}
