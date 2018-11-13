package pg_store

import (
	. "github.com/gtforge/BookStore/models"
	"github.com/gtforge/services_common_go/gett-storages"
	"github.com/jinzhu/gorm"
)

type PgBooksStoreInterface interface {
	GetAll() ([]Book, error)
	GetPerPage(pageNumber int) ([]Book, error)
	GetById(Id uint) (Book, error)
	Create(Name string, Quantity int) error
	Update(Id uint, Name string, Quantity int) error
	Delete(Id uint) error
}

var Instance PgBooksStoreInterface = &pgBooksStore{
	db: gettStorages.DB.New(),
}

type pgBooksStore struct {
	db *gorm.DB
}

const booksTableName = "books"

func (b *pgBooksStore) GetAll() ([]Book, error) {
	var books []Book

	err := gettStorages.DB.Order("id ASC").Find(&books).Error

	return books, err
}

func (b *pgBooksStore) GetPerPage(pageNumber int) ([]Book, error) {
	var books []Book
	itemsPerPage := 10
	offset := (pageNumber-1)*itemsPerPage + 1

	err := gettStorages.DB.Order("id ASC").Offset(offset).Limit(itemsPerPage).Find(&books).Error

	return books, err
}

func (b *pgBooksStore) GetById(id uint) (Book, error) {
	var book Book

	err := gettStorages.DB.Table(booksTableName).First(&book, "id = ?", id).Error

	return book, err
}

func (b *pgBooksStore) Create(name string, quantity int) error {
	book := Book{
		Name:     name,
		Quantity: quantity,
	}

	return gettStorages.DB.Table(booksTableName).Create(&book).Error
}

func (b *pgBooksStore) Update(id uint, name string, quantity int) error {
	return gettStorages.DB.Table(booksTableName).Where("id = ?", id).Updates(Book{Name: name, Quantity: quantity}).Error
}

func (b *pgBooksStore) Delete(id uint) error {
	book := Book{ID: id}

	return gettStorages.DB.Table(booksTableName).Delete(&book).Error
}
