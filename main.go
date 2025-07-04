package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var message = "welcome to index page"
	w.Write([]byte(message))
}

func hello(w http.ResponseWriter, r *http.Request) {
	var message = "HELLOW"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index", index)
	http.HandleFunc("/hello", hello)

	var address = "localhost:3000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
