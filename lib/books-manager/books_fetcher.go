package booksManager

import (
	"github.com/gtforge/BookStore/models"
	"github.com/gtforge/BookStore/store"
)

type BooksFetcher interface {
	GetAll() ([]models.Book, error)
}

func NewBooksFetcher() BooksFetcher {
	return &booksFetcher{}
}

type booksFetcher struct{}

func (bm *booksFetcher) GetAll() (books []models.Book, err error) {
	books, err = store.BooksStore.GetAllBooks()

	return books, err
}
