package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	userid := this.GetSession("userid")
	if userid == nil || userid == 0 {
		this.TplName = "login.tpl"
		return
	}
	this.Data["Website"] = userid
	this.Data["Email"] = "tsaolee@yegrow.com"
	this.TplName = "index.tpl"
}

func (this *MainController) Post() {
	this.Data["msg"] = this.GetString("id")
	this.TplName = "post.tpl"
}
