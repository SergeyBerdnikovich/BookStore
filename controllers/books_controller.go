package controllers

import (
	"fmt"
	"net/http"

	"github.com/gtforge/BookStore/lib/books-manager"
)

type BooksController struct {
	Base
}

//func (c *BooksController) Init(ct *context.Context, controllerName, actionName string, app interface{}) {
//	c.Controller.Init(ct, controllerName, actionName, app)
//}

// @router /books [get]
func (c *BooksController) Index() {
	books, err := booksManager.NewBooksFetcher().GetAll()

	if err != nil {
		errMsg := fmt.Sprintf("error while inserting PrioritiesGroups\nerror: %v", err.Error())
		c.Data["json"] = "{\"error\":\"" + errMsg + "\"}"
		c.ServeJSON()
		c.Ctx.Output.SetStatus(http.StatusOK)
	} else {
		c.Data["json"] = books
		c.ServeJSON()
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
}

// @router /books/:id [get]
func (c *BooksController) Show() {
	//book_id := c.GetString(":id")

	c.Ctx.Output.SetStatus(http.StatusOK)
}

// @router /books/new [get]
func (c *BooksController) New() {
	c.Ctx.Output.SetStatus(http.StatusOK)
}

// @router /books [post]
func (c *BooksController) Create() {
	//book_id := c.GetString(":id")

	c.Ctx.Output.SetStatus(http.StatusOK)
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
	//book_id := c.GetString(":id")

	c.Ctx.Output.SetStatus(http.StatusOK)
}
