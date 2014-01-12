package main

import (
	"RaysGo/models"
	_ "RaysGo/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.CreateDb()
	beego.SessionOn = true
	beego.Run()
}
