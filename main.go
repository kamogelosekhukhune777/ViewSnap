package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kamogelosekhukhune777/ViewSnap/views"
)

var (
	homeView     *views.View
	contactView  *views.View
	faqView      *views.View
	notfoundView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	if err := faqView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	if err := notfoundView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("views/home.html")
	contactView = views.NewView("views/contact.html")
	faqView = views.NewView("views/faq.html")
	notfoundView = views.NewView("views/notfound.html")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notfound)
	fmt.Println("starting server at port :8080")
	log.Fatal("error starting server", http.ListenAndServe(":8080", r))
}
