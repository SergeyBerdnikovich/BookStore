package controllers

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gtforge/services_common_go/gett-auth"
	"github.com/gtforge/services_common_go/gett-env"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// Base - Base controller for all UI controller
type Base struct {
	gettAuth.AuthController
}

// Init - initialization for base controller with basic parameters
func (c *Base) Init(ct *context.Context, controllerName, actionName string, app interface{}) {
	// this part activates the authentication
	c.AuthController.Init(ct, controllerName, actionName, app)

	// this is the main html file
	c.Layout = "layout.html"

	// tabs in your navbar
	c.Data["menu"] = map[string]string{
		"Tab 1": beego.URLFor("Base.Index"),
		"Tab 2": beego.URLFor("Base.Index"),
	}

	// this part is about the env dropdown
	if env := c.GetString("env"); env != "" {
		c.Data["current_env"] = gettEnv.Envs[strings.ToUpper(env)]
	} else {
		c.Data["current_env"] = gettEnv.Envs["IL"]
	}
}

// Index - Main page
func (c *Base) Index() {
	err := c.Render()
	if err != nil {
		logrus.Error(err)
	}
}
