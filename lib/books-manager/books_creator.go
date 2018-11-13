package booksManager

import (
	"github.com/gtforge/BookStore/store/books-store"
)

type BooksCreatorInterface interface {
	Create(name string, quantity int) error
}

var BooksCreatorInstance BooksCreatorInterface = &booksCreator{}

type booksCreator struct{}

func (bm *booksCreator) Create(name string, quantity int) (err error) {
	return books_store.Instance.Create(name, quantity)
}
