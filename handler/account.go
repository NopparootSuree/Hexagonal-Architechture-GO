package handler

import (
	"encoding/json"
	"net/http"

	"github.com/NopparootSuree/Hexagonal-Architechture-GO/errs"
	"github.com/NopparootSuree/Hexagonal-Architechture-GO/service"
	"github.com/gorilla/mux"
)

type accountHandler struct {
	accServ service.AccontService
}

func NewAccountHandler(accServ service.AccontService) accountHandler {
	return accountHandler{accServ: accServ}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["id"]

	if r.Header.Get("content-type") != "application/json" {
		handleError(w, errs.NewValidationError("request body incorect format"))
		return
	}

	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorect format"))
		return
	}

	response, err := h.accServ.NewAccount(customerID, request)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["id"]

	response, err := h.accServ.GetAccount(customerID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
