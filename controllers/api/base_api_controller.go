package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gtforge/services_common_go/gett-env"
	gettErrors "github.com/gtforge/services_common_go/gett-ops/gett-beego/controllers"

	"github.com/astaxie/beego"
)

// BaseAPIController - controller for all APIs
type BaseAPIController struct {
	Env gettEnv.Env

	beego.Controller
}

// AbortF - aborts with formatted error message
func (c *BaseAPIController) AbortF(code int, status, formatMessage string, args ...interface{}) {
	c.Ctx.Input.SetData("error", fmt.Sprintf(formatMessage, args...))
	if status == "" {
		status = http.StatusText(code)
	}
	c.CustomAbort(code, status)
}

// AbortE - aborts with plain error message
func (c *BaseAPIController) AbortE(code int, status string, err error) {
	c.Ctx.Input.SetData("error", err.Error())
	if status == "" {
		status = http.StatusText(code)
	}
	c.CustomAbort(code, status)
}

// EnsureEnv - ensure that request has proper env code
func (c *BaseAPIController) EnsureEnv() {
	envCode := strings.ToUpper(c.GetString(":env"))
	if env, ok := gettEnv.Envs[envCode]; ok {
		c.Env = env
	} else {
		c.AbortF(http.StatusOK, gettErrors.InputQueryError, "%s is not a valid env code", envCode)
	}
}
