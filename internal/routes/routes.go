package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-manager/api/internal/updater"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux() // this is new server multiplexer

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

	var req updater.CreateTaskRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // optional but recommended

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// Validate basic fields
	if req.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	updater.UpdateTask(req)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "task created",
	})
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete task\n")
}
