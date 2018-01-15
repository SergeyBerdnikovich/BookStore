package apiv1

import (
	"github.com/astaxie/beego"
	"fmt"
)

// BaseAPIController - for all APIs Controller
type BaseAPIController struct {
	beego.Controller
}

func (c *BaseAPIController) AbortF(code, formatMessage string, args ...interface{}) {
	c.Ctx.Input.SetData("error", fmt.Sprintf(formatMessage, args...))
	c.Abort(code)
}

func (c *BaseAPIController) AbortE(code string, err error) {
	c.Ctx.Input.SetData("error", err.Error())
	c.Abort(code)
}