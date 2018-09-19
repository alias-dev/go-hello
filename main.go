package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort     = "8080"
	defaultGreeting = "Hello, World!"
)

type templateData struct {
	Greeting string
}

func greetingHandler(greeting string, t *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, templateData{Greeting: greeting})
	}
}

func main() {
	port := os.Getenv("PORT")
	greeting := os.Getenv("GREETING")

	if port == "" {
		port = defaultPort
	}

	if greeting == "" {
		greeting = defaultGreeting
	}

	tmpl := template.Must(template.ParseFiles("greeting.tmpl.html"))

	http.HandleFunc("/", greetingHandler(greeting, tmpl))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
