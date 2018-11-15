// Package routers for service
// @APIVersion 1.0.0
// @Title Gett BookStore
// @Description Gett BookStore
// @Name Gett
// @Schemes http
package routers

import (
	"github.com/astaxie/beego"
	"github.com/gtforge/BookStore/controllers"
	"github.com/gtforge/BookStore/controllers/api/v1"
	_ "github.com/gtforge/services_common_go/gett-ops/gett-beego"
	_ "github.com/gtforge/services_common_go/gett-settings/controllers"
)

func init() {
	beego.Router("/", &controllers.Base{}, "Get:Index")
	beego.SetStaticPath("/assets", "assets")
	beego.SetStaticPath("/public/assets", "public/assets")

	beego.AutoRouter(&controllers.BooksController{})

	apiV1 := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSInclude(
				&v1.BooksController{},
			),
		),
	)
	beego.AddNamespace(apiV1)
}
