package booksManager

import (
	"github.com/gtforge/BookStore/store/books-store"
)

type BooksDestroyerInterface interface {
	Delete(id uint) error
}

var BooksDestroyerInstance BooksDestroyerInterface = &booksDestroyer{}

type booksDestroyer struct{}

func (bm *booksDestroyer) Delete(id uint) (err error) {
	return books_store.Instance.Delete(id)
}
