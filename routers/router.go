package routers

import (
	"RaysGo/controllers"
	"RaysGo/controllers/admin"
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
	beego.Router("/myposts", post, "*:List")
	beego.Router("/post/new", post, "*:New;post:NewPost")
	beego.Router("/post/view/:id:int", post, "*:View")
	beego.Router("/post/edit/:id:int", post, "*:Edit;post:EditPost")
	beego.Router("/post/delete/:id:int", post, "*:Delete")
	beego.Router("/post/comment/:id:int", post, "post:Comment")
	beego.Router("/post/tag/:name", post, "*:Tag")

	// admin
	beego.Router("/admin", &controllers.AdminController{})

	config := new(admin.ConfigController)
	beego.Router("/admin/config", config, "*:Config;post:ConfigPost")
}
