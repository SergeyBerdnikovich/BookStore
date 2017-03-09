package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
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

// Make more test requests with less code
func MakeTestRequest(verb string, path string, data []byte) (w *httptest.ResponseRecorder) {
	payload := new(bytes.Buffer)
	payload.Write(data)

	r, _ := http.NewRequest(verb, path, payload)
	w = httptest.NewRecorder()

	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Info("testing", fmt.Sprintf("[%s] %s | %d\n%s", r.Method, r.URL, w.Code, w.Body.String()))

	return
}
