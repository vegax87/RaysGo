package controllers

import (
	"RaysGo/models"
	"RaysGo/helpers"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	this.GoView("user/index")
}

func (this *UserController) View() {
	id,_ := helpers.Str2Int64(this.GetParam(":id"))

	var user = models.User{Id: int64(id)}
	has, err := models.Engine.Get(&user)

	if !has || err != nil {
		this.Abort("404")
	}

	this.Data["user"] = user
	this.setMeta("Title", user.Name)
	this.GoView("user/view")
}

// TODO
func (this *UserController) Edit() {
	//	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	this.Data["Title"] = "Edit - "
	this.TplNames = "user/edit.html"
}

// TODO
func (this *UserController) EditPost(){
	//	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	this.Data["Title"] = "Edit - "
	this.TplNames = "user/edit.html"
}
