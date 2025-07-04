package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func index(w http.ResponseWriter, r *http.Request) {
	var data = M{"name": "Mr. Ard"}
	var tmpl = template.Must(template.ParseFiles(
		"views/index.html",
		"views/_header.html",
		"views/_footer.html",
	))

	var err = tmpl.ExecuteTemplate(w, "index", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	var data = M{"name": "Mr. Ard"}
	var tmpl = template.Must(template.ParseFiles(
		"views/about.html",
		"views/_header.html",
		"views/_footer.html",
	))

	var err = tmpl.ExecuteTemplate(w, "about", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	var message = "HELLOW"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/hello", hello)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	var address = "localhost:3000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
