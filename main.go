package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
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
	var message string
	var port int
	flag.StringVar(&message, "m", "Hello, Go!", "A friendly greeting")
	flag.IntVar(&port, "p", 80, "Bind port")
	flag.Parse()
	log.Printf("p: %v, m: %v", port, message)
	tmpl := template.Must(template.ParseFiles("greeting.tmpl.html"))

	http.HandleFunc("/", greetingHandler(message, tmpl))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}
