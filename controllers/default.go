package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "yegrow.com"
	c.Data["Email"] = "tsaolee@yegrow.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Data["msg"] = c.GetString("id")
	c.TplName = "post.tpl"
}
