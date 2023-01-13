package main

import (
	"fmt"
	"net/http"
	"restapi-gorilla/controllers/productcontroller"
	"restapi-gorilla/models"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	models.ConnectDatabase()

	route.HandleFunc("/products", productcontroller.Home).Methods("GET")
	route.HandleFunc("/product/{id}", productcontroller.Show).Methods("GET")
	route.HandleFunc("/product", productcontroller.Create).Methods("POST")
	route.HandleFunc("/product/{id}", productcontroller.Update).Methods("PUT")
	route.HandleFunc("/product", productcontroller.Delete).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	http.ListenAndServe("localhost:8080", route)
}
