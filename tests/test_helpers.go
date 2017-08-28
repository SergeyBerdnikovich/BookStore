package tests

import (
	"path/filepath"
	"runtime"

	"github.com/astaxie/beego"
)

// init Beeego framework for testing
func initBeego() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}
