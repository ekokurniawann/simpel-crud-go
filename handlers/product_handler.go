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

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "invalid input")
		return
	}

	val := &validator.Validator{}
	val.Check(product.ID != "", "missing product ID")
	val.Check(product.Name != "", "missing product name")
	val.Check(product.Description != "", "missing product description")
	val.Check(product.Price > 0, "invalid product price")

	if val.HasErrors() {
		response.ErrorResponse(w, http.StatusBadRequest, val)
		return
	}

	datastore.CreateProduct(product)
	response.SuccessResponse(w, http.StatusCreated, product)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	product, exists := datastore.GetProduct(id)
	if !exists {
		response.ErrorResponse(w, http.StatusNotFound, "product not found")
		return
	}

	response.SuccessResponse(w, http.StatusOK, product)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "invalid input")
		return
	}

	val := &validator.Validator{}
	val.Check(product.Name != "", "missing product name")
	val.Check(product.Description != "", "missing product description")
	val.Check(product.Price > 0, "invalid product price")
	val.Check(product.ID == id, "ID mismatch")

	if val.HasErrors() {
		response.ErrorResponse(w, http.StatusBadRequest, val)
		return
	}

	if !datastore.UpdateProduct(product) {
		response.ErrorResponse(w, http.StatusNotFound, "product not found")
		return
	}

	response.SuccessResponse(w, http.StatusOK, product)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if !datastore.DeleteProduct(id) {
		response.ErrorResponse(w, http.StatusNotFound, "product not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
