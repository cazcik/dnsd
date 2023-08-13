package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/cazcik/utils/handler"
)

type DNSData struct {
	Host string
}

func main() {
	web := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.Execute(w, nil)
	}

	results := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/results.html"))
		domain := r.PostFormValue("domain")
		response := handler.GetDomain(domain)
		tmpl.Execute(w, response)
	}

	about := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/about.html"))
		tmpl.Execute(w, nil)
	}

	contact := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/contact.html"))
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", web)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/domain/", results)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
