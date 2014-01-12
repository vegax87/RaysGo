package controllers

import (
//	"strconv"
//	"RaysGo/models"
//	"RaysGo/helpers"
)

type UserController struct {
	BaseController
}

func (this *UserController) Get() {
//	this.TplNames = "user/index.tpl"
	this.GoView("user/index")
}

func (this *UserController) View() {
//	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

//	user := models.User{Uid: id}
//
//	err := orm.NewOrm().Read(&user)
//
//	if err != nil {
//		this.Abort("404")
//	}
//
//	this.Data["user"] = user

	this.GoView("user/view")
}

func (this *UserController) Register() {
//	valid := validation.Validation{}
//	user := models.User{
//		Name:     this.GetString("name"),
//		Email:    this.GetString("email"),
//		Password: this.GetString("password"),
//		Status:   1,
//		Role_id:  2,
//	}
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

	this.setMetas(map[string]string{
		"Title": "Register",
	})

	this.GoView("user/register")
}

func (this *UserController) Login() {
	
//	user := models.User{
//		Name :     this.GetString("name"),
//		Password : this.GetString("password"),
//	}
//	
//	fmt.Printf(helpers.EncryptPassword(user.Password, nil))
	this.setMetas(map[string]string{
		"Title": "Login",
	})
	this.GoView("user/login")
}

func (this *UserController) Edit() {
//	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

//	user := models.User{Uid: id}
//
//	err := orm.NewOrm().Read(&user)
//
//	if err != nil {
//		this.Abort("404")
//	}
//
//	this.Data["user"] = user
//
//	this.setMetas(map[string]string{
//		"Title": "Edit - " + user.Name,
//	})

	this.GoView("edit")
}
