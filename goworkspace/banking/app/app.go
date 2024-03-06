package app

import (
	"bankingweb/domain"
	"bankingweb/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	//request matcher regex followed after :
	//router.HandleFunc("/customers/{customer_id:[0-9]+}", getAllCustomer).Methods(http.MethodGet)

	//method matcher .methods
	//router.HandleFunc("/customers", createCustomers).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
