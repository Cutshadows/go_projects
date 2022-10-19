package routes

import "net/http"

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world Tasks"))
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world Tasks"))
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world Tasks"))
}

func DeletetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world Tasks"))
}
