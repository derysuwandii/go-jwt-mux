package main

import (
	"github.com/gorilla/mux"
	"go-jwt-mux/controllers/authcontroller"
	"go-jwt-mux/controllers/productcontroller"
	"go-jwt-mux/middlewares"
	"go-jwt-mux/models"
	"log"
	"net/http"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/product", productcontroller.Index).Methods("GET")
	api.HandleFunc("/product/{id}", productcontroller.Show).Methods("GET")
	api.HandleFunc("/product", productcontroller.Create).Methods("POST")
	api.HandleFunc("/product/{id}", productcontroller.Update).Methods("PUT")
	api.HandleFunc("/product", productcontroller.Delete).Methods("DELETE")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
