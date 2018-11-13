// Package routers for service
// @APIVersion 1.0.0
// @Title Gett BookStore
// @Description Gett BookStore
// @Name Gett
// @Schemes http
package routers

import (
	"github.com/gtforge/BookStore/controllers"
	// "github.com/gtforge/BookStore/controllers/api/v1"

	// import gett-beego package with share settings and endpoints
	_ "github.com/gtforge/services_common_go/gett-ops/gett-beego"

	"github.com/astaxie/beego"
)

func init() {
	// Add new routes here, e.g.:
	//
	// apiV1 := beego.NewNamespace("/api",
	// 	beego.NSNamespace("/v1",
	// 		beego.NSInclude(
	// 			&v1.ExampleController{},
	// 		),
	// 	),
	// 	)
	// beego.AddNamespace(apiV1)
	beego.Router("/", &controllers.Base{}, "Get:Index")
	beego.SetStaticPath("/assets", "assets")
	beego.SetStaticPath("/public/assets", "public/assets")

	beego.Router("/books", &controllers.BooksController{}, "get:Index")
	beego.Router("/books/:id", &controllers.BooksController{}, "get:Show")
	beego.Router("/books/new", &controllers.BooksController{}, "get:New")
	beego.Router("/books", &controllers.BooksController{}, "post:Create")
	beego.Router("/books/:id/edit", &controllers.BooksController{}, "get:Edit")
	beego.Router("/books/:id", &controllers.BooksController{}, "put:Update")
	beego.Router("/books/:id", &controllers.BooksController{}, "post:Destroy")
}
