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
	// create index.html template
	web := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.Execute(w, nil)
	}

	// create results.html template
	results := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/results.html"))
		domain := r.PostFormValue("domain")

		// if PostFormValue is not a domain, return invalid template
		if (!govalidator.IsDNSName(domain)) {
			log.Printf("[invalid lookup]: %s\n", domain)
			invStr := fmt.Sprintf("<div class='flex items-center justify-center'><p class='flex text-neutral-500'>invalid domain: %s</p></div>", domain)
			tmpl, _ := template.New("invalid").Parse(invStr)
			tmpl.Execute(w, invStr)
			return
		}

		// assuming valid domain call handler.GetDomain to retrieve the hosts data
		log.Printf("[lookup]: %s\n", domain)
		response, err := handler.GetDomain(domain)
		if err != nil {
			log.Fatal(err)
		}

		// execute and pass hosts data to results.html template
		tmpl.Execute(w, response)
	}

	// create about.html template
	about := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/about.html"))
		tmpl.Execute(w, nil)
	}

	// routing
	http.HandleFunc("/", web)
	http.HandleFunc("/about", about)
	http.HandleFunc("/domain", results)

	// start the http server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
