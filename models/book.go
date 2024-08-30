package models

import (
	"database/sql"
	"example/bookstore/database"
	"fmt"
)

// Book represents the structure of our resource
type Book struct {
	ID     int64   `json:"id" gorm:"primaryKey"` // ID is the primary key
	Title  string  `json:"title"`                // Title of the book
	Author string  `json:"author"`               // Author of the book
	Price  float64 `json:"price"`                // Price of the book
}

func GetBooks() ([]Book, error) {
	rows, err := database.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func GetBookByID(id string) (Book, error) {
	var book Book
	query := "SELECT * FROM books WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price)
	if err == sql.ErrNoRows {
		return book, fmt.Errorf("Book not found")
	}
	return book, err
}

func (a *Book) AddBook() error {
	query := "INSERT INTO books (title, author, price) VALUES (?,?,?)"
	_, err := database.DB.Exec(query, a.Title, a.Author, a.Price)
	return err
}
