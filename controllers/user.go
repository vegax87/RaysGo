package controllers

import (
	"RaysGo/helpers"
	"RaysGo/models"
	"fmt"
	"strconv"
	"time"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
	//	this.TplNames = "user/index.tpl"
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
	this.GoView("user/view")
}

func (this *UserController) Register() {
	valid := validation.Validation{}

	user := models.User{}
	if err := this.ParseForm(&user); err!=nil{
		fmt.Println(err)
	}else{
		
	}
	// user := models.User{
	// 	Name:       this.GetString("name"),
	// 	Email:      this.GetString("email"),
	// 	Password:   this.GetString("password"),
	// 	Status:     1,
	// 	Rid:        2,
	// 	CreateTime: time.Now(),
	// }
	user.Password = helpers.EncryptPassword(user.Password, nil)

	//	b, valid_err := valid.Valid(user)
	//	if valid_err != nil {
	//		// handle error
	//	}
	//	if !b {
	//		for _, verr := range valid.Errors {
	//			fmt.Println(verr.Key + " : " + verr.Message)
	//		}
	//	} else {
	//		user.Password = helpers.EncryptPassword(user.Password,nil)
	//		id, err := orm.NewOrm().Insert(&user)
	//		if err == nil {
	//			this.Redirect("/user/view/"+strconv.Itoa(int(id)), 302)
	//		}
	//	}

	_, err := models.Engine.Insert(&user)
	if err == nil {
		this.Redirect("/user/view/"+strconv.Itoa(int(user.Id)), 302)
	}

	this.setMetas(map[string]string{
		"Title": "Register",
	})

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

	this.setMetas(map[string]string{
		"Title": "Login",
	})
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
