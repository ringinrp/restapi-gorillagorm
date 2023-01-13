package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", home).Methods("GET")

	fmt.Println("Server running on port 8000")
	http.ListenAndServe("localhost:8000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")

	fmt.Println("Hello semua")
}
