package controllers

import (
	"RaysGo/models"
	"strconv"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	this.GoView("user/index")
}

func (this *UserController) View() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	var user = models.User{Id: int64(id)}
	has, err := models.Engine.Get(&user)

	if !has || err != nil {
		this.Abort("404")
	}

	this.Data["user"] = user
	this.setMeta("Title", user.Name)
	this.GoView("user/view")
}

func (this *UserController) Edit() {
	//	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	this.GoView("edit")
}
