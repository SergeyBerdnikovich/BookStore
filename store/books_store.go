package store

import . "github.com/gtforge/BookStore/models"

type IBooksStore interface {
	GetAllBooks() ([]Book, error)
	GetBooksPerPage(pageNumber int) ([]Book, error)
	GetBook(Id int) (Book, error)
	CreateBook(Name string, Quantity int) (Book, error)
	UpdateBook(Id int, Name string, Quantity int) error
	DeleteBook(Id int) error
}

var BooksStore IBooksStore
