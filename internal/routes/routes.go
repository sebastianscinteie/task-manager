package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hello", helloHandler)

	// TASKS
	mux.HandleFunc("GET /task", getTaskHandler)
	mux.HandleFunc("POST /task", postTaskHandler)
	// mux.HandleFunc("PUT /task", putTaskHandler)
	mux.HandleFunc("DELETE /task", deleteTaskHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!!!\n")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get task\n")
}

func postTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post task\n")
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete task\n")
}
