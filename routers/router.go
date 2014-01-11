package routers

import (
	"RaysGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index", &controllers.MainController{}, "*:Index")
    beego.Router("/user/view/:id:int", &controllers.UserController{}, "*:View")
}
