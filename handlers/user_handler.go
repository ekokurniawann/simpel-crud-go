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

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "invalid input")
		return
	}


	val := &validator.Validator{}
	val.Check(user.ID != "", "missing user ID")
	val.Check(user.Name != "", "missing user name")
	val.Check(user.Email != "", "missing user email")

	if val.HasErrors() {
		response.ErrorResponse(w, http.StatusBadRequest, val)
		return
	}

	datastore.CreateUser(user)
	response.SuccessResponse(w, http.StatusCreated, user)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, exists := datastore.GetUser(id)
	if !exists {
		response.ErrorResponse(w, http.StatusNotFound, "user not found")
		return
	}

	response.SuccessResponse(w, http.StatusOK, user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "invalid input")
		return
	}

	val := &validator.Validator{}
	val.Check(user.Name != "", "missing user name")
	val.Check(user.Email != "", "missing user email")
	val.Check(user.ID == id, "ID mismatch")

	if val.HasErrors() {
		response.ErrorResponse(w, http.StatusBadRequest, val)
		return
	}

	if !datastore.UpdateUser(user) {
		response.ErrorResponse(w, http.StatusNotFound, "user not found")
		return
	}

	response.SuccessResponse(w, http.StatusOK, user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if !datastore.DeleteUser(id) {
		response.ErrorResponse(w, http.StatusNotFound, "user not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
