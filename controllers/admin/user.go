package admin

import(
	"RaysGo/controllers"
	"RaysGo/helpers"
	"fmt"
)

type UserController struct{
	controllers.BaseController
}

// TODO
func (this *UserController) Delete(){
	id,_ := helpers.Str2Int64(this.GetParam(":id"))
	fmt.Println(id)
}

func (this *UserController) List(){

	this.Data["Title"] = "Users Administration"
	this.TplNames = "admin/user_list.html"
}

// Activate a user
func (this *UserController) Active(){

}

func (this *UserController) Block(){

}


