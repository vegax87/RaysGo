package controllers

import (
	"RaysGo/helpers"
	"RaysGo/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strconv"
	"time"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	this.GoView("user/index")
}

func (this *UserController) View() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	var user = models.User{Id: int64(id)}
	has, err := models.Engine.Get(&user)

	if !has || err != nil {
		this.Abort("404")
	}

	this.Data["user"] = user
	this.setMeta("Title", user.Name)
	this.GoView("user/view")
}

func (this *UserController) Register() {
	valid := validation.Validation{}

	user := models.User{}
	if err := this.ParseForm(&user); err != nil {
		fmt.Println(err)
	} else {
		user.Status = 1
		user.CreateTime = time.Now()
		user.Rid = 2
	}

	ok, valid_err := valid.Valid(user)
	if valid_err != nil {
		// 
	}
	if !ok {
		for _, verr := range valid.Errors {
			fmt.Println(verr.Key + " : " + verr.Message)
		}
	} else {
		user.Password = helpers.EncryptPassword(user.Password, nil)
		if _, err := models.Engine.Insert(&user); err == nil {
			this.Redirect("/user/view/"+strconv.Itoa(int(user.Id)), 302)
		}
	}

	this.setMeta("Title", "Register")
	this.GoView("user/register")
}

func (this *UserController) Login() {

	name := this.GetString("name")
	if name != "" {
		user := models.User{
			Name: this.GetString("name"),
		}

		has, err := models.Engine.Get(&user)
		if has && err == nil && helpers.ValidatePassword(user.Password, this.GetString("password")) {
			this.SetSession("username", user.Name)
			this.SetSession("userid", int(user.Id))
			this.SetSession("userrole", int(user.Rid))
			this.SetSession("useremail", user.Email)

			this.Redirect("/user/view/"+fmt.Sprintf("%d", user.Id), 302)
		}
	}

	this.setMeta("Title", "Login")
	this.GoView("user/login")
}

func (this *UserController) Logout() {
	if session_uid != 0 {
		this.DestroySession()
	}

	this.Redirect("/", 302)
}

func (this *UserController) Edit() {
	//	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	this.GoView("edit")
}
