// Package routers for service
// @APIVersion 1.0.0
// @Title Gett {|PROJECTNAME|}
// @Description Gett {|PROJECTNAME|}
// @Name Gett
// @Schemes http
package routers

import (
	"github.com/gtforge/{|PROJECTNAME|}/controllers"
	// "github.com/gtforge/{|PROJECTNAME|}/controllers/api/v1"
	
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
}
