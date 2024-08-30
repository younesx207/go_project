package handlers

import (
	"example/bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Homepage handles the root URL ("/") and returns a welcome message in JSON format
func Homepage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Bookstore RESTful API with Go, Gin and MySQL"})
}

// GetBooks handles the GET request to "/books" and returns all books in JSON format
func GetBooks(c *gin.Context) {
	// Retrieve the list of books from the database
	books, err := models.GetBooks()
	if err != nil {
		// If an error occurs while retrieving the books, return an internal server error response with the error message
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// If no error occurs, return the list of books with an HTTP 200 OK status
	c.IndentedJSON(http.StatusOK, books)
}

// GetBookByID handles the GET request to "/books/:id" and returns the book with the specified ID in JSON format
func GetBookByID(c *gin.Context) {
	// Extract the book ID from the URL parameter
	id := c.Param("id")

	// Retrieve the book from the database using the ID
	book, err := models.GetBookByID(id)
	if err != nil {
		// If the book is not found, return an HTTP 404 Not Found status with a corresponding message
		if err.Error() == "book not found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		} else {
			// If another error occurs, return an internal server error response with the error message
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	// If the book is found, return it with an HTTP 200 OK status
	c.IndentedJSON(http.StatusOK, book)
}

// PostBooks handles the POST request to "/books" and adds a new book to the database
func PostBooks(c *gin.Context) {
	var newBook models.Book

	// Bind the JSON request body to the newBook variable; return an error if the binding fails
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the database
	if err := newBook.AddBook(); err != nil {
		// If an error occurs while adding the book, return an internal server error response with the error message
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// If the book is added successfully, return the new book with an HTTP 201 Created status
	c.IndentedJSON(http.StatusCreated, newBook)
}
