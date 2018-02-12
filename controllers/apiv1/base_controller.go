package apiv1

import (
	"fmt"

	"github.com/astaxie/beego"
)

// BaseAPIController - for all APIs Controller
type BaseAPIController struct {
	beego.Controller
}

// AbortF - aborts with formatted error message
func (c *BaseAPIController) AbortF(code, formatMessage string, args ...interface{}) {
	c.Ctx.Input.SetData("error", fmt.Sprintf(formatMessage, args...))
	c.Abort(code)
}

// AbortE - aborts with plain error message
func (c *BaseAPIController) AbortE(code string, err error) {
	c.Ctx.Input.SetData("error", err.Error())
	c.Abort(code)
}
