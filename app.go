package main

import (
	"RaysGo/helpers"
	"RaysGo/models"
	_ "RaysGo/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.CreateDb()

	helpers.AddViewFunc()

	beego.SessionOn = true

	beego.Run()
}
