package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

// =================================================================================================
// GetBook godoc
// @Summary      Get all books
// @Description  Get all books in DB.
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Book
// @Router       /book [get]
func GetBook(response http.ResponseWriter, request *http.Request) {
	newbooks := models.GetAllBooks()
	response_body, _ := json.Marshal(newbooks)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(response_body)
}

// =================================================================================================
// GetBook godoc
// @Summary      Get books by Id.
// @Description  Get books by Id in DB.
// @Tags         book
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.Book
// @Router       /book/{bookId} [get]
func GetBookById(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	bookId, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing.")
	}
	bookDetails, _ := models.GetBookById(bookId)
	response_body, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(response_body)
}

// =================================================================================================
// CreateBook godoc
// @Summary      Create a Book.
// @Description  Create a Book in DB.
// @Tags         book-create
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.Book
// @Router       /book/{bookId} [post]
func CreateBook(response http.ResponseWriter, request *http.Request) {
	created_book := &models.Book{}
	utils.ParseBody(request, created_book)
	book := created_book.CreateBook()
	response_body, _ := json.Marshal(book)
	response.WriteHeader(http.StatusOK)
	response.Write(response_body)
}

// =================================================================================================
// DeleteBook godoc
// @Summary      Delete a Book.
// @Description  Delete a Book in DB.
// @Tags         book-delete
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.Book
// @Router       /book/{bookId} [delete]
func DeleteBook(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	bookId, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing.")
	}
	book := models.DeleteBook(bookId)
	response_body, _ := json.Marshal(book)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(response_body)
}

// =================================================================================================
// UpdateBook godoc
// @Summary      Update a Book.
// @Description  Update a Book in DB.
// @Tags         book-update
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.Book
// @Router       /book/{bookId} [put]
func UpdateBook(response http.ResponseWriter, request *http.Request) {
	var update_book = &models.Book{}
	utils.ParseBody(request, update_book)
	params := mux.Vars(request)
	bookId, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing.")
	}
	book_details, db := models.GetBookById(bookId)

	getBookDetails(update_book, book_details)
	db.Save(&book_details)

	response_body, _ := json.Marshal(book_details)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(response_body)
}

func getBookDetails(update_book *models.Book, book_details *models.Book) {
	if update_book.Name != "" {
		book_details.Name = update_book.Name
	}

	if update_book.Author != "" {
		book_details.Author = update_book.Author
	}

	if update_book.Publication != "" {
		book_details.Publication = update_book.Publication
	}
}
