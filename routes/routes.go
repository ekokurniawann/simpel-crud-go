package routes

import (
	"crud-simpel/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")

	router.HandleFunc("/products", handlers.CreateProductHandler).Methods("POST")
	router.HandleFunc("/products/{id}", handlers.GetProductHandler).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.UpdateProductHandler).Methods("PUT")
	router.HandleFunc("/products/{id}", handlers.DeleteProductHandler).Methods("DELETE")

	router.HandleFunc("/transactions", handlers.CreateTransactionHandler).Methods("POST")
	router.HandleFunc("/transactions/{id}", handlers.GetTransactionHandler).Methods("GET")
	router.HandleFunc("/transactions/{id}", handlers.UpdateTransactionHandler).Methods("PUT")
	router.HandleFunc("/transactions/{id}", handlers.DeleteTransactionHandler).Methods("DELETE")

	return router
}
