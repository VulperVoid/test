package controllers

import (
	"test/models"
	_ "test/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Data["name"] = "Jack"
}

func (this *UserController) Post() {
	user := models.User{}
	err := user.Add(this.GetString("name"), this.GetString("password"))
	this.Data["id"] = user.Id
	this.Data["err"] = err
}

func (this *UserController) Login() {
	user := models.User{}
	user, err := user.Get(this.GetString("name"), this.GetString("password"))
	if user.Id > 0 {
		this.SetSession("userid", user.Id)
		this.Redirect("/", 200)
		return
	} else {
		if err == nil {
			this.Data["msg"] = "用户名或密码错误，请重试"
		} else {
			this.Data["msg"] = err
		}
		this.TplName = "login.tpl"
	}
}
