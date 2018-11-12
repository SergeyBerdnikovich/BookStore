package booksManager

import (
	"github.com/gtforge/BookStore/models"
	"github.com/gtforge/BookStore/store"
	"github.com/k0kubun/pp"
)

type BooksFetcher interface {
	GetAll() ([]models.Book, error)
}

func NewBooksFetcher() BooksFetcher {
	return &booksFetcher{}
}

type booksFetcher struct{}

func (bm *booksFetcher) GetAll() (books []models.Book, err error) {


	pp.Println("store.BooksStore is ======>", store.BooksStore)

	books, err = store.BooksStore.GetAllBooks()

	return books, err
}
