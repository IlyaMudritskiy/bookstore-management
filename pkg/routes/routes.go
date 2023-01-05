package routes

import (
	"github.com/gorilla/mux"
	"github.com/ilyamudritskiy/bookstore_management/pkg/controllers"
)

var RegisterRoutes = func(router *mux.Router) {

	// CreateBook
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")

	// GetBooks - Get all books in DB
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")

	// GetBookById
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")

	// UpdateBook
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")

	// DeleteBook
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
