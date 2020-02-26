package error_handler

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error401() {
	c.TplName = "401.html"
}

func (c *ErrorController) Error403() {
	c.TplName = "403.html"
}

func (c *ErrorController) Error404() {
	c.TplName = "404.html"
}

func (c *ErrorController) Error500() {
	c.TplName = "500.html"
}


