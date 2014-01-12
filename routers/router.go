package routers

import (
	"RaysGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index", &controllers.MainController{}, "*:Index")

    beego.Router("/user/register", &controllers.UserController{}, "*:Register")
    beego.Router("/user/login", &controllers.UserController{}, "*:Login")
    beego.Router("/user/logout", &controllers.UserController{}, "*:Logout")
    beego.Router("/user/view/:id:int", &controllers.UserController{}, "*:View")
    beego.Router("/user/edit/:id:int", &controllers.UserController{}, "*:Edit")
}
