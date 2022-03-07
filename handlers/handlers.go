package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleReq() {
	log.Println("Start development server localhost:8001")
	myRouter := mux.NewRouter().StrictSlash(true)
	handler := cors.AllowAll().Handler(myRouter)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/login", Login).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/user", CreateUser).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/users", ListUser).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/user/{id}", DetailUser).Methods("OPTIONS", "GET")
	log.Fatal(http.ListenAndServe(":8001", handler))
}
