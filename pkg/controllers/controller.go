package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/ilyamudritskiy/bookstore_management/pkg/utils"
	"github.com/ilyamudritskiy/bookstore_management/pkg/models"
)

var NewBook models.Book

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	var CreateBook = &models.Book{}

	utils.ParseBody(request, CreateBook)

	var book = CreateBook.CreateBook()

	response, _ := json.Marshal(book)

	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	var newBooks = models.GetAllBooks()

	response, _ := json.Marshal(newBooks)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	var vars = mux.Vars(request)
	var bookID = vars["id"]

	ID, err := strconv.ParseInt(bookID, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing Int")
	}

	bookDetails, _ := models.GetBookById(ID)
	response, _ := json.Marshal(bookDetails)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var updateBook = &models.Book{}
	var vars = mux.Vars(request)
	var bookId = vars["id"]

	utils.ParseBody(request, updateBook)

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing Int")
	}

	bookDetails, db := models.GetBookById(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	response, _ := json.Marshal(bookDetails)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	var vars = mux.Vars(request)
	var bookID = vars["id"]

	ID, err := strconv.ParseInt(bookID, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing Int")
	}

	var book = models.DeleteBook(ID)

	response, _ := json.Marshal(book)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
