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

func (c *BooksController) Index() {
	beego.ReadFromRequest(&c.Controller)

	books, err := booksManager.BooksFetcherInstance.GetAll()
	if err != nil {
		logrus.Error(err.Error())
	}

	c.Data["books"] = books

	err = c.Render()
	if err != nil {
		logrus.Error(err.Error())
	}
}

func (c *BooksController) Show() {
	bookId, err := c.GetInt("id")
	if err != nil {
		logrus.Error(err.Error())
	}

	book, err := booksManager.BooksFetcherInstance.GetById(uint(bookId))
	if err != nil {
		logrus.Error(err.Error())
	}

	c.Data["book"] = book

	err = c.Render()
	if err != nil {
		logrus.Error(err.Error())
	}
}

func (c *BooksController) New() {
	beego.ReadFromRequest(&c.Controller)

	c.Data["book"] = &models.Book{}

	err := c.Render()
	if err != nil {
		logrus.Error(err.Error())
	}
}

func (c *BooksController) Create() {
	book := models.Book{}

	err := c.ParseForm(&book)
	if err != nil {
		logrus.Error(err.Error())
	}

	err = booksManager.BooksCreatorInstance.Create(book.Name, book.Quantity)
	if err != nil {
		c.flashError(err.Error())
		c.Redirect(beego.URLFor("BooksController.New"), http.StatusFound)
	} else {
		c.flashSuccess("Book was created successfully")
		c.Redirect(beego.URLFor("BooksController.Index"), http.StatusFound)
	}
}

func (c *BooksController) Edit() {
	beego.ReadFromRequest(&c.Controller)

	bookId, err := c.GetInt("id")
	if err != nil {
		logrus.Error(err.Error())
	}

	book, err := booksManager.BooksFetcherInstance.GetById(uint(bookId))
	if err != nil {
		logrus.Error(err.Error())
	}

	c.Data["book"] = book

	err = c.Render()
	if err != nil {
		logrus.Error(err.Error())
	}
}

func (c *BooksController) Update() {
	bookId, err := c.GetInt("id")
	if err != nil {
		logrus.Error(err.Error())
	}

	book := models.Book{ID: uint(bookId)}

	err = c.ParseForm(&book)
	if err != nil {
		logrus.Error(err.Error())
	}

	err = booksManager.BooksUpdaterInstance.Update(book.ID, book.Name, book.Quantity)
	if err != nil {
		c.flashError(err.Error())
		c.Redirect(beego.URLFor("BooksController.Edit", "id", bookId), http.StatusFound)
	} else {
		c.flashSuccess("Book was updated successfully")
		c.Redirect(beego.URLFor("BooksController.Show", "id", bookId), http.StatusFound)
	}
}

func (c *BooksController) Destroy() {
	bookId, err := c.GetInt("id")
	if err != nil {
		logrus.Error(err.Error())
	}

	err = booksManager.BooksDestroyerInstance.Delete(uint(bookId))
	if err != nil {
		c.flashError(err.Error())
	} else {
		c.flashSuccess("Book was deleted")
	}

	c.Redirect(beego.URLFor("BooksController.Index"), http.StatusFound)
}
