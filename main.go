package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/cazcik/utils/handler"
)

func main() {
	web := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.Execute(w, nil)
	}

	results := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/results.html"))
		domain := r.PostFormValue("domain")
		if (!govalidator.IsDNSName(domain)) {
			fmt.Printf("[invalid lookup]: %s\n", domain)
			invStr := fmt.Sprintf("<p class='flex text-center text-neutral-500'>invalid domain: %s</p>", domain)
			tmpl, _ := template.New("invalid").Parse(invStr)
			tmpl.Execute(w, invStr)
			return
		}
		fmt.Printf("[lookup]: %s\n", domain)
		response := handler.GetDomain(domain)
		tmpl.Execute(w, response)
	}

	about := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/about.html"))
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", web)
	http.HandleFunc("/about", about)
	http.HandleFunc("/domain/", results)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
