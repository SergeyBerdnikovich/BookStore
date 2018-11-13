package controllers

import (
	"net/http"

	"github.com/astaxie/beego"

	"github.com/Sirupsen/logrus"

	"github.com/gtforge/BookStore/models"

	"github.com/gtforge/BookStore/lib/books-manager"
)

type BooksController struct {
	Base
}

// @router /books [get]
func (c *BooksController) Index() {
	books, _ := booksManager.BooksFetcherInstance.GetAll()

	c.Data["books"] = books

	err := c.Render()
	if err != nil {
		logrus.Error(err.Error())
	}
}

// @router /books/:id [get]
func (c *BooksController) Show() {
	bookId, _ := c.GetUint8(":id")
	book, _ := booksManager.BooksFetcherInstance.GetById(uint(bookId))

	c.Data["book"] = book

	err := c.Render()
	if err != nil {
		logrus.Error(err.Error())
	}
}

// @router /books/new [get]
func (c *BooksController) New() {
	beego.ReadFromRequest(&c.Controller)

	c.Data["book"] = &models.Book{}

	err := c.Render()
	if err != nil {
		logrus.Error(err.Error())
	}
}

// @router /books [post]
func (c *BooksController) Create() {
	book := models.Book{}
	err := c.ParseForm(&book)

	if err != nil {
		logrus.Error(err.Error())
	}

	if err := booksManager.BooksCreatorInstance.Create(book.Name, book.Quantity); err != nil {
		c.flashError(err.Error())
		c.Redirect("/books/new", http.StatusFound)
		return
	}

	c.flashSuccess("Book was created successfully")
	c.Redirect("/books", http.StatusFound)
}

// @router /books/:id/edit [get]
func (c *BooksController) Edit() {
	//book_id := c.GetString(":id")

	c.Ctx.Output.SetStatus(http.StatusOK)
}

// @router /books/:id [put]
func (c *BooksController) Update() {
	//book_id := c.GetString(":id")

	c.Ctx.Output.SetStatus(http.StatusOK)
}

// @router /books/:id [delete]
func (c *BooksController) Destroy() {
	bookId, _ := c.GetUint8(":id")

	if err := booksManager.BooksDestroyerInstance.Delete(uint(bookId)); err != nil {
		c.flashError(err.Error())
	} else {
		c.flashSuccess("Book was deleted")
	}

	c.Redirect("/books", http.StatusFound)
}
