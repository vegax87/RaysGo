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
	page := 1
	pagesize := 10

	if v, err := helpers.Str2Int(this.GetParam(":page")); err == nil {
		page = v
		if page <= 0 {
			page = 1
		}
	}

	if v, err := helpers.Str2Int(this.GetParam(":pagesize")); err == nil {
		pagesize = v
	}

	users := make([]models.User, 0)
	models.Engine.OrderBy("id desc").Limit(pagesize, (page-1)*pagesize).Find(&users)
	count, _ := models.Engine.Count(new(models.User))

	this.Data["Page"] = page
	this.Data["PageSize"] = pagesize
	this.Data["Total"] = int(count)
	if len(users) > 0 {
		this.Data["Users"] = users
	}
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
