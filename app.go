package main

import (
	"RaysGo/models"
	"RaysGo/helpers"
	_ "RaysGo/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.CreateDb()

	helpers.AddViewFunc()
	
	beego.SessionOn = true
	
	beego.Run()
}
