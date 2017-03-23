package controllers

import (
	"test/models"
	_ "test/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["name"] = "Jack"
}

func (c *UserController) Post() {
	user := models.User{}
	user, err := user.Get(c.GetString("name"), c.GetString("password"))
	c.Data["id"] = user.Id
	c.Data["err"] = err
}

//func (c *UserController) Post() {
//	user := models.User{}
//	err := user.Add(c.GetString("name"), c.GetString("password"))
//	c.Data["err"] = err
//}
