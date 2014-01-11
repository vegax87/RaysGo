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

	this.TplNames = "user/view.tpl"
}

func (this *UserController) Register(){
	user := models.User{
		Name: this.GetString("name"),
		Email: this.GetString("email"),
		Password: this.GetString("password"),
		Status: 1,
		Role_id: 2,
	}

	id, err := orm.NewOrm().Insert(&user)
	if err == nil {
		this.Ctx.Redirect(302, "/user/view/" + strconv.Itoa(int(id)))
	}
	this.setMetas(map[string] string{
		"Title" : "Register",
		})

	this.TplNames = "user/register.tpl"
}

func (this *UserController) Login(){
	this.setMetas(map[string] string{
		"Title" : "Login",
		})

	this.TplNames = "user/login.tpl"
}

func (this *UserController) Edit(){
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	user := models.User{Uid: id}

	err := orm.NewOrm().Read(&user)

	if err != nil {
		this.Abort("404")
	}

	this.Data["user"] = user

	this.setMetas(map[string] string{
		"Title" : "Edit - " + user.Name,
		})

	this.TplNames = "user/edit.tpl"
}