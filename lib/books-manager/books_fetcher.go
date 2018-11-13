package booksManager

import (
	. "github.com/gtforge/BookStore/models"
	"github.com/gtforge/BookStore/store/books-store"
)

type BooksFetcherInterface interface {
	GetAll() ([]Book, error)
	GetPerPage(pageNumber int) ([]Book, error)
	GetById(Id uint) (Book, error)
}

var BooksFetcherInstance BooksFetcherInterface = &booksFetcher{}

type booksFetcher struct{}

func (bm *booksFetcher) GetAll() (books []Book, err error) {
	return books_store.Instance.GetAll()
}

func (bm *booksFetcher) GetPerPage(pageNumber int) ([]Book, error) {
	return books_store.Instance.GetPerPage(pageNumber)
}

func (bm *booksFetcher) GetById(Id uint) (Book, error) {
	return books_store.Instance.GetById(Id)
}
