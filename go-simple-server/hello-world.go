package main

import (
	"net/http"
)

func main() {

	//route
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	//init server
	http.ListenAndServe(":2000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world from Home"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world from Home"))
}
