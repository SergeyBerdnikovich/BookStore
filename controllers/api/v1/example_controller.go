package v1

import (
	"net/http"

	bContext "github.com/astaxie/beego/context"
	"github.com/gtforge/BookStore/controllers/api"
)

// ExampleController shows an example of controller creation
type ExampleController struct {
	api.BaseAPIController
}

// Init function to pass data to base class
func (c *ExampleController) Init(ct *bContext.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ct, controllerName, actionName, app)
}

// Example endpoint
// @Summary Some summary
// @Description Some description
// @Param   env    			 path    string  	true       "Gett env ru/uk/us/il"
// @Success 200 success return empty JSON
// @router /example [get]
func (c *ExampleController) Example() {
	// Get params like this:
	//
	// env := c.GetString(":env")

	// Handle errors like this:
	//
	// if err := models.TestFunc(); err != nil {
	// 	c.AbortF(http.StatusBadRequest, gettErrors.InputQueryError, "err: %s", err)
	// }

	result := make(map[string]interface{})
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = result
	c.ServeJSON()
}
