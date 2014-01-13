package routers

import (
	"RaysGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index", &controllers.MainController{}, "*:Index")

    login := new(controllers.LoginController)
    beego.Router("/login", login, "get:Get;post:Login")
    beego.Router("/logout", login, "get:Logout")

    register := new(controllers.RegisterController)
    beego.Router("/register", register, "*:Get;post:Register")

    beego.Router("/user/register", &controllers.UserController{}, "*:Register")
    beego.Router("/user/logout", &controllers.UserController{}, "*:Logout")
    beego.Router("/user/view/:id:int", &controllers.UserController{}, "*:View")
    beego.Router("/user/edit/:id:int", &controllers.UserController{}, "*:Edit")
}
