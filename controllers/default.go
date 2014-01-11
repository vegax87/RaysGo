package controllers

import (
	"github.com/astaxie/beego/orm"
	"RaysGo/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	
	this.Layout = "layout/default.tpl"
	this.TplNames = "index.tpl"
}

func (this *MainController) Index() {
	var users []*models.User

	this.Layout = "layout/default.tpl"

    user := new(models.User)
	query := orm.NewOrm().QueryTable(user)
	count, _ := query.Count()
	query.OrderBy("-uid").All(&users)

	this.Data["users"] = users
	this.Data["count"] = count
	this.TplNames = "users.tpl"
}