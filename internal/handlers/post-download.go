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
	url := req.PostFormValue("url")

	tmpl := template.Must(template.ParseFiles("./internal/templates/index.html"))
	tmpl.ExecuteTemplate(w, "download-list", models.DownloadResponse{Url: url})
}
