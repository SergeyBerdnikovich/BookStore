package v1

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
	"github.com/gtforge/BookStore/controllers/api"
	"github.com/gtforge/BookStore/lib/books-manager"
)

type BooksController struct {
	api.BaseAPIController
}

// @router /books [get]
func (c *BooksController) Index() {
	beego.ReadFromRequest(&c.Controller)

	pageNumber, err := c.GetInt("page")
	if err != nil {
		logrus.Error(err.Error())

		c.Data["json"] = "{\"error\":\"page parameter is missing\"}"
		c.ServeJSON()

		c.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)
		return
	}

	books, err := booksManager.BooksFetcherInstance.GetPerPage(pageNumber)
	if err != nil {
		logrus.Error(err.Error())
	}

	c.Data["json"] = books
	c.ServeJSON()

	c.Ctx.Output.SetStatus(http.StatusOK)
}

// @router /books/buy [post]
func (c *BooksController) Buy() {
	bookId, err := c.GetInt("book_id")
	if err != nil {
		logrus.Error(err.Error())

		c.Data["json"] = "{\"error\":\"book_id parameter is missing\"}"
		c.ServeJSON()

		c.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)
		return
	}

	err = booksManager.BooksBuyerInstance.Buy(uint(bookId))
	if err != nil {
		logrus.Error(err.Error())

		c.Data["json"] = "{\"error\":\"" + err.Error() + "\"}"
		c.ServeJSON()

		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		return
	}

	c.Data["json"] = "{\"message\":\"Enjoy your new book\"}"
	c.ServeJSON()
	c.Ctx.Output.SetStatus(http.StatusOK)
}
