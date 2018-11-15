package booksManager

import (
	"github.com/gtforge/BookStore/store/books-store"
)

type BooksUpdaterInterface interface {
	Update(id uint, name string, quantity int) error
}

var BooksUpdaterInstance BooksUpdaterInterface = &booksUpdater{}

type booksUpdater struct{}

func (bm *booksUpdater) Update(id uint, name string, quantity int) (err error) {
	return books_store.Instance.Update(id, name, quantity)
}
