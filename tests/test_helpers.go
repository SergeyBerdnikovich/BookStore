package tests

import (
	"path/filepath"
	"runtime"

	"github.com/astaxie/beego"
)

// init Beeego framework for testing
func initBeego() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(frame.File, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}
