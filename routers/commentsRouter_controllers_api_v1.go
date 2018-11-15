package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/gtforge/BookStore/controllers/api/v1:BooksController"] = append(beego.GlobalControllerRouter["github.com/gtforge/BookStore/controllers/api/v1:BooksController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/books`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/gtforge/BookStore/controllers/api/v1:BooksController"] = append(beego.GlobalControllerRouter["github.com/gtforge/BookStore/controllers/api/v1:BooksController"],
		beego.ControllerComments{
			Method:           "Buy",
			Router:           `/books/buy`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/gtforge/BookStore/controllers/api/v1:ExampleController"] = append(beego.GlobalControllerRouter["github.com/gtforge/BookStore/controllers/api/v1:ExampleController"],
		beego.ControllerComments{
			Method:           "Example",
			Router:           `/example`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

}
