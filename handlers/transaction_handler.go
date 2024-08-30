package handlers

import (
	"crud-simpel/datastore"
	"crud-simpel/models"
	"crud-simpel/response"
	"crud-simpel/validator"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "invalid input")
		return
	}


	val := &validator.Validator{}
	val.Check(transaction.ID != "", "missing transaction ID")
	val.Check(transaction.UserID != "", "missing user ID")
	val.Check(transaction.ProductID != "", "missing product ID")
	val.Check(transaction.Quantity > 0, "invalid quantity")
	val.Check(transaction.Total > 0, "invalid total")

	if val.HasErrors() {
		response.ErrorResponse(w, http.StatusBadRequest, val)
		return
	}

	datastore.CreateTransaction(transaction)
	response.SuccessResponse(w, http.StatusCreated, transaction)
}

func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	transaction, exists := datastore.GetTransaction(id)
	if !exists {
		response.ErrorResponse(w, http.StatusNotFound, "transaction not found")
		return
	}

	response.SuccessResponse(w, http.StatusOK, transaction)
}

func UpdateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "invalid input")
		return
	}

	val := &validator.Validator{}
	val.Check(transaction.UserID != "", "missing user ID")
	val.Check(transaction.ProductID != "", "missing product ID")
	val.Check(transaction.Quantity > 0, "invalid quantity")
	val.Check(transaction.Total > 0, "invalid total")
	val.Check(transaction.ID == id, "ID mismatch")

	if val.HasErrors() {
		response.ErrorResponse(w, http.StatusBadRequest, val)
		return
	}

	if !datastore.UpdateTransaction(transaction) {
		response.ErrorResponse(w, http.StatusNotFound, "transaction not found")
		return
	}

	response.SuccessResponse(w, http.StatusOK, transaction)
}

func DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if !datastore.DeleteTransaction(id) {
		response.ErrorResponse(w, http.StatusNotFound, "transaction not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
