package controllers

import (
//	"RaysGo/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"

	this.setMetas(map[string] string{
		"Title" : "Home",
		"Author" : "Raysmond",
		"Description" : "Simple blog system in Go!",
		"Keywords" : "RaysGo, Raysmond",
		})

	this.GoView("index","layout/default")
}

func (this *MainController) Index() {
	this.GoView("users")
}
