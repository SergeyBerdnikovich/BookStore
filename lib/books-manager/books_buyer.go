package booksManager

import (
	"github.com/gtforge/BookStore/store/books-store"
)

type BooksBuyerInterface interface {
	Buy(id uint) error
}

var BooksBuyerInstance BooksBuyerInterface = &booksBuyer{}

type booksBuyer struct{}

func (bm *booksBuyer) Buy(id uint) (err error) {
	return books_store.Instance.DecreaseQuantity(id)
}
