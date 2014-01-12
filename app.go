package main

import (
	_ "RaysGo/routers"
	"github.com/astaxie/beego"
	"RaysGo/models"
)

func main() {
	models.CreateDb()
 	beego.SessionOn = true
	beego.Run()
}
