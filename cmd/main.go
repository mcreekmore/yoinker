package main

import (
	"log"
	"net/http"

	"github.com/mcreekmore/yoinker/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.NewIndexHandler().ServeHTTP)

	mux.HandleFunc("POST /download/", handlers.NewDownloadHandler().ServeHTTP)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
