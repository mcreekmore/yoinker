package main

import (
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
				{Title: "Inglourious Basterds", Director: "Quentin Tarantino"},
				{Title: "Kill Bill", Director: "Quentin Tarantino"},
			},
		}
		tmpl.Execute(w, films)
	})

	mux.HandleFunc("POST /download/", func(w http.ResponseWriter, req *http.Request) {
		title := req.PostFormValue("title")
		director := req.PostFormValue("director")

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
