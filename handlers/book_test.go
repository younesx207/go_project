package handlers

import (
	"bytes"
	"encoding/json"
	"example/bookstore/database"
	"example/bookstore/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// SetupRouter initializes the Gin router and defines the API endpoints.
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/books", GetBooks)        // Route to get all books
	r.POST("/books", PostBooks)      // Route to add a new book
	r.GET("/books/:id", GetBookByID) // Route to get a book by its ID
	return r
}

// TestGetBooks tests the GET /books endpoint to retrieve all books.
func TestGetBooks(t *testing.T) {
	// Connect to the database
	database.Connect()

	// Setup the router with the API routes
	router := SetupRouter()

	// Create a new HTTP GET request to the /books endpoint
	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Record the response using httptest
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert that the status code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK, got %v", rr.Code)

	// Unmarshal the response body into a slice of books
	var books []models.Book
	err = json.Unmarshal(rr.Body.Bytes(), &books)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Assert that the books slice is not nil
	assert.NotNil(t, books, "Expected books to be non-nil")
	// Assert that the response contains at least 0 books (no error expected if empty)
	assert.GreaterOrEqual(t, len(books), 0, "Expected at least 0 books in response")

	// Check that each book in the response has non-nil fields
	for _, book := range books {
		assert.NotNil(t, book.ID, "Expected book ID to be non-nil")
		assert.NotNil(t, book.Title, "Expected book title to be non-nil")
		assert.NotNil(t, book.Author, "Expected book author to be non-nil")
		assert.NotNil(t, book.Price, "Expected book price to be non-nil")
	}
}

// TestPostBooks tests the POST /books endpoint to add a new book.
func TestPostBooks(t *testing.T) {
	// Connect to the database
	database.Connect()

	// Setup the router with the API routes
	router := SetupRouter()

	// Create a new book to be added
	book := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		Price:  19.99,
	}

	// Marshal the book struct into JSON format
	jsonBook, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Failed to marshal book: %v", err)
	}

	// Create a new HTTP POST request with the JSON book data
	req, err := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(jsonBook))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json") // Set the content type to application/json

	// Record the response using httptest
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert that the status code is 201 Created
	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status Created, got %v", rr.Code)

	// Unmarshal the response body into a book struct
	var createdBook models.Book
	err = json.Unmarshal(rr.Body.Bytes(), &createdBook)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Assert that the created book is not nil and matches the input book
	assert.NotNil(t, createdBook, "Expected createdBook to be non-nil")
	assert.Equal(t, book.Title, createdBook.Title, "Expected title to match")
	assert.Equal(t, book.Author, createdBook.Author, "Expected author to match")
	assert.Equal(t, book.Price, createdBook.Price, "Expected price to match")
}

// TestGetBooksById tests the GET /books/:id endpoint to retrieve a specific book by its ID.
func TestGetBooksById(t *testing.T) {
	// Connect to the database
	database.Connect()

	// Setup the router with the API routes
	router := SetupRouter()

	// Define the book ID to retrieve
	bookID := 1

	// Create a new HTTP GET request for the specific book ID
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/books/%d", bookID), nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Record the response using httptest
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Assert that the status code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK, got %v", rr.Code)

	// Unmarshal the response body into a book struct
	var book models.Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Assert that the retrieved book is not nil and has the expected ID and non-nil fields
	assert.NotNil(t, book, "Expected book to be non-nil")
	assert.Equal(t, int64(bookID), book.ID, "Expected book ID to match")
	assert.NotNil(t, book.Title, "Expected book title to be non-nil")
	assert.NotNil(t, book.Author, "Expected book author to be non-nil")
	assert.NotNil(t, book.Price, "Expected book price to be non-nil")
}
