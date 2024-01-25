package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)

	//request matcher regex followed after :
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getAllCustomer)

	//method matcher .methods
	router.HandleFunc("/customers", createCustomers).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])

}
func createCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request recieved")
}
