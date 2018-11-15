package pg_store

import (
	"errors"

	. "github.com/gtforge/BookStore/models"
	"github.com/gtforge/services_common_go/gett-settings"
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
	DecreaseQuantity(Id uint) error
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
	itemsPerPage := gettSettings.EnvSetting("DEV", "Books", "maximum_rows_per_page").AsInt()
	offset := (pageNumber - 1) * itemsPerPage

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

func (b *pgBooksStore) DecreaseQuantity(id uint) error {
	for {
		var book Book

		err := gettStorages.DB.Table(booksTableName).First(&book, "id = ?", id).Error
		if err != nil {
			return err
		}

		newQuantity := book.Quantity - 1

		if newQuantity < 0 {
			return errors.New("Negative quantity")
		}

		result := gettStorages.DB.Table(booksTableName).Where("id = ? AND quantity = ?", id, book.Quantity).Updates(map[string]interface{}{"quantity": newQuantity})
		err = result.Error
		if err != nil {
			return err
		}
		if result.RowsAffected > 0 {
			return nil
		}
	}
}
