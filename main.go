package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	log.Println("running on port 8000")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "Godfather", Director: "Francis"},
				{Title: "Blade Runner", Director: "Ridley"},
				{Title: "Thing", Director: "John"},
			},
		}

		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
