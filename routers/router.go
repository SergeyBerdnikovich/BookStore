// @APIVersion 1.0.0
// @Title Gett {|PROJECTNAME|}
// @Description  Gett {|PROJECTNAME|}
// @Name Gett
// @Schemes https,http
// @Host {|PROJECTNAME|}.gett.com
package routers

import (
	_ "github.com/gtforge/services_common_go/gett-ops/gett-beego"
)

func init() {
	// Add new routes here
	// example
	//apiV1 := beego.NewNamespace("/api",
	//	beego.NSNamespace("/v1",
	//		beego.NSInclude(
	//			&controllers.DriverControllerV2{},
	//		),
	//	),
	//	)
	//beego.AddNamespace(apiV1)
}
