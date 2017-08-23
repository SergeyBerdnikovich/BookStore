package main

import (
	_ "github.com/gtforge/permission_enforcement_service/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
