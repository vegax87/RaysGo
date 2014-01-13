package routers

import (
	"RaysGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    main := new(controllers.MainController)
    beego.Router("/", main)
    beego.Router("/about", main, "*:About")
    beego.Router("/contact", main, "*:Contact;post:ContactPost")

    login := new(controllers.LoginController)
    beego.Router("/login", login, "get:Get;post:Login")
    beego.Router("/logout", login, "get:Logout")

    register := new(controllers.RegisterController)
    beego.Router("/register", register, "*:Get;post:Register")

    user := new(controllers.UserController)
    beego.Router("/user/view/:id:int", user, "*:View")
    beego.Router("/user/edit/:id:int", user, "*:Edit")

    post := new(controllers.PostController)
    beego.Router("/post/view/:id:int", post, "*:View")
    beego.Router("/myposts", post, "*:List")
    beego.Router("/post/new", post, "*:New;post:NewPost")
    beego.Router("/post/edit", post, "*:Edit;post:EditPost")
    beego.Router("/post/delete/:id:int", post, "*:Delete")
}
