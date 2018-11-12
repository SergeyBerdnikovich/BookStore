package pg_store

import (
	. "github.com/gtforge/BookStore/models"
	"github.com/gtforge/BookStore/store"
	"github.com/gtforge/services_common_go/gett-storages"
)

func init() {
	if store.BooksStore != nil {
		panic("store.BooksStore already initialized")
	}
	store.BooksStore = &pgBooksStoreImpl{}
}

type pgBooksStoreImpl struct{}

const booksTableName = "books"

func (b *pgBooksStoreImpl) GetAllBooks() ([]Book, error) {
	var books []Book

	err := gettStorages.DB.Order("id ASC").Find(&books).Error

	return books, err
}

func (b *pgBooksStoreImpl) GetBooksPerPage(pageNumber int) ([]Book, error) {
	var books []Book
	itemsPerPage := 10
	offset := (pageNumber-1)*itemsPerPage + 1

	err := gettStorages.DB.Order("id ASC").Offset(offset).Limit(itemsPerPage).Find(&books).Error

	return books, err
}

func (b *pgBooksStoreImpl) GetBook(id int) (Book, error) {
	var book Book

	err := gettStorages.DB.Table(booksTableName).First(&book, "id = ?", id).Error

	return book, err
}

func (b *pgBooksStoreImpl) CreateBook(name string, quantity int) (Book, error) {
	book := Book{
		Name:     name,
		Quantity: quantity,
	}

	err := gettStorages.DB.Table(booksTableName).Create(&book).Error

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (b *pgBooksStoreImpl) UpdateBook(id int, name string, quantity int) error {
	return gettStorages.DB.Table(booksTableName).Where("id = ?", id).Updates(Book{Name: name, Quantity: quantity}).Error
}

func (b *pgBooksStoreImpl) DeleteBook(id int) error {
	book := Book{Id: id}

	return gettStorages.DB.Table(booksTableName).Delete(&book).Error
}
