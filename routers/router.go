// Package routers for service
// @APIVersion 1.0.0
// @Title Gett permission_enforcement_service
// @Description Gett permission_enforcement_service
// @Name Gett
// @Schemes http
package routers

import (
	"github.com/astaxie/beego"
	"github.com/gtforge/permission_enforcement_service/controllers"
	// import gett-beego package with share settings and endpoints
	_ "github.com/gtforge/services_common_go/gett-ops/gett-beego"
)

func init() {
	// Add new routes here
	// example
	//apiV1 := beego.NewNamespace("/api",
	//	beego.NSNamespace("/v1",
	//		beego.NSInclude(
	//			&apiv1.BaseAPIController{},
	//		),
	//	),
	//	)
	//beego.AddNamespace(apiV1)
	beego.Router("/", &controllers.Base{}, "Get:Index")
	beego.SetStaticPath("/assets", "assets")
	beego.SetStaticPath("/public/assets", "public/assets")
}
