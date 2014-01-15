package admin

import (
	"RaysGo/controllers"
	"RaysGo/helpers"
	"RaysGo/models"
)

type UserController struct {
	controllers.BaseController
}

// TODO
func (this *UserController) Delete() {
	//id, _ := helpers.Str2Int64(this.GetParam(":id"))

}

func (this *UserController) List() {

	this.Data["Title"] = "Users Administration"
	this.TplNames = "admin/user_list.html"
}

// Activate a user
func (this *UserController) Active() {
	id, _ := helpers.Str2Int64(this.GetParam(":id"))
	user := models.User{}
	if _, err := models.Engine.Id(id).Get(&user); err == nil {
		user.Status = models.ACTIVE
		if _, e := models.Engine.Id(id).Update(&user); e == nil {
			// do something
		} else {
			// do something
		}
	}
}

func (this *UserController) Block() {
	id, _ := helpers.Str2Int64(this.GetParam(":id"))
	user := models.User{}
	if _, err := models.Engine.Id(id).Get(&user); err == nil {
		user.Status = models.BLOCKED
		if _, e := models.Engine.Id(id).Update(&user); e == nil {
			// do something
		} else {
			// do something
		}
	}
}
