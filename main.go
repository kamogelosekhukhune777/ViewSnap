package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	//fmt.Fprintln(w, "<h1>Welcome to my ViewSnap!</h1>")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, "<h1>Frquently Asked Questions</h1><p>Here is a list commonly asked questions.</p>")
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "<h1>Sorry, but we couldn't find the page you were looking for.</h1>")
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.html")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("views/contact.html")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notfound)
	fmt.Println("starting server at port :8080")
	log.Fatal("error starting server", http.ListenAndServe(":8080", r))
}
