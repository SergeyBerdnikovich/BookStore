package books_store

import (
	. "github.com/gtforge/BookStore/models"
	"github.com/gtforge/BookStore/store/books-store/pg-store"
)

type BooksStoreInterface interface {
	GetAll() ([]Book, error)
	GetPerPage(pageNumber int) ([]Book, error)
	GetById(Id uint) (Book, error)
	Create(Name string, Quantity int) error
	Update(Id uint, Name string, Quantity int) error
	Delete(Id uint) error
	DecreaseQuantity(Id uint) error
}

var Instance BooksStoreInterface = &booksStore{
	db: pg_store.Instance,
}

type booksStore struct {
	db BooksStoreInterface
}

func (bs *booksStore) GetAll() ([]Book, error) {
	return bs.db.GetAll()
}

func (bs *booksStore) GetPerPage(pageNumber int) ([]Book, error) {
	return bs.db.GetPerPage(pageNumber)
}

func (bs *booksStore) GetById(Id uint) (Book, error) {
	return bs.db.GetById(Id)
}

func (bs *booksStore) Create(Name string, Quantity int) error {
	return bs.db.Create(Name, Quantity)
}

func (bs *booksStore) Update(Id uint, Name string, Quantity int) error {
	return bs.db.Update(Id, Name, Quantity)
}

func (bs *booksStore) Delete(Id uint) error {
	return bs.db.Delete(Id)
}

func (bs *booksStore) DecreaseQuantity(Id uint) error {
	return bs.db.DecreaseQuantity(Id)
}
