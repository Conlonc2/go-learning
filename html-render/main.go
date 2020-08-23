package main

import (
	"log"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := struct {
		Msg string
	}{
		Msg: "test",
	}
	tmpl.Execute(w, data)
}

func routeHandler() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	routeHandler()
}
