package main

import (
	"github.com/astaxie/beego"
	_ "github.com/gtforge/BookStore/routers"
)

func main() {
	beego.Run()
}
