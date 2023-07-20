package handler

import (
	"encoding/json"
	"net/http"

	"github.com/NopparootSuree/Hexagonal-Architechture-GO/service"
	"github.com/gorilla/mux"
)

type customerHandler struct {
	custSev service.CustomerService
}

func NewCustomerHandler(custSev service.CustomerService) customerHandler {
	return customerHandler{custSev: custSev}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customer, err := h.custSev.GetCustomers()
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	customer, err := h.custSev.GetCustomerByID(param)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
