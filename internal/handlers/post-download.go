package handlers

import (
	"net/http"
	"text/template"

	"github.com/mcreekmore/yoinker/pkg/models"
)

type DownloadHandler struct{}

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{}
}

func (h *DownloadHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	title := req.PostFormValue("title")
	director := req.PostFormValue("director")

	tmpl := template.Must(template.ParseFiles("../internal/templates/index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", models.Film{Title: title, Director: director})
}
