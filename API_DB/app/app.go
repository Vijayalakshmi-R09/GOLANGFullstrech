package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/getAllEmployees", getAllEmployees).Methods("GET")
	router.HandleFunc("/createEmployee", createEmployee).Methods("POST")
	router.HandleFunc("/getEmployee/{EMP_ID}", getEmployeeByID).Methods("GET")
	router.HandleFunc("/getEmployeeByName/{EMP_NAME}", getEmployeeByName).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
