package controllers

import (
	"github.com/astaxie/beego/orm"
	"RaysGo/models"
	"strconv"
)

type UserController struct{
	BaseController
}

func (this *UserController) Get(){
	this.Layout = "layout/default.tpl"
	this.TplNames = "user/index.tpl"
}


func (this *UserController) View(){
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	user := models.User{Uid: id}

	err := orm.NewOrm().Read(&user)

	if err != nil {
		this.Abort("404")
	}

	this.Data["user"] = user
	this.Layout = "layout/default.tpl"
	this.TplNames = "user/view.tpl"
}