package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("hello world!")
}

func health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
