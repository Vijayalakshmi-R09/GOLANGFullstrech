package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/getAllEmployees", GetAllEmployees).Methods("GET")
	router.HandleFunc("/createEmployee", CreateEmployee).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
