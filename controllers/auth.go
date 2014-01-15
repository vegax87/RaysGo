package controllers

import (
	"RaysGo/helpers"
	"RaysGo/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"time"
)

type LoginController struct {
	BaseController
}

// Get implemented login page
func (this *LoginController) Get() {
	if this.isLogin {
		this.Redirect("/user/view/"+fmt.Sprintf("%d", this.User().Id), 302)
	}
	this.Data["Title"] = "Login"
	this.TplNames = "auth/login.html"

	//loginRedirect := strings.TrimSpace(this.GetString("to"))
}

// Post implemented login action
func (this *LoginController) Login() {
	var (
		user  models.User
		form  models.LoginForm
		valid validation.Validation
	)

	if err := this.ParseForm(&form); err != nil {
		fmt.Println(err)
	} else {
		if ok, valid_err := valid.Valid(form); ok && valid_err == nil {
			user.Name = form.UserName
			//user.Password = helpers.EncryptPassword(form.Password, nil)
			has, gerr := models.Engine.Get(&user)
			if has && gerr == nil && helpers.ValidatePassword(user.Password, form.Password) {
				this.SetSession("username", user.Name)
				this.SetSession("userid", int(user.Id))
				this.SetSession("userrole", int(user.Rid))
				this.SetSession("useremail", user.Email)

				this.Redirect("/user/view/"+fmt.Sprintf("%d", user.Id), 302)
			}
		} else {
			for _, e := range valid.Errors {
				this.FlashError(e.Key + " : " + e.Message)
			}
		}
	}

	this.SaveFlash()
	this.Data["Form"] = form
	this.Data["Title"] = "Login"
	this.TplNames = "auth/login.html"
	// this.Redirect("/login", 302)
}

func (this *LoginController) Logout() {
	if session_uid != 0 {
		this.DestroySession()
	}

	this.Redirect("/", 302)
}

type RegisterController struct {
	BaseController
}

func (this *RegisterController) Get() {
	this.Data["Title"] = "Register"
	this.TplNames = "auth/register.html"
}

// Post
func (this *RegisterController) Register() {
	var (
		user  models.User
		form  models.RegisterForm
		valid validation.Validation
		err   error
	)

	if err = this.ParseForm(&form); err != nil {
		fmt.Println(err)
	} else {
		if ok, e := valid.Valid(form); ok && e == nil {
			user.Name = form.UserName
			user.Email = form.Email
			user.Password = helpers.EncryptPassword(form.Password, nil)
			user.Status = models.ACTIVE
			user.Rid = models.ROLE_AUTHENTICATED
			user.CreateTime = time.Now()

			if _, err = models.Engine.Insert(&user); err == nil {
				this.FlashError(user.Name + " registered successfully. Please login now!")
				this.SaveFlash()
				this.Redirect("/login", 302)
				return
			}
		} else {
			for _, e := range valid.Errors {
				this.FlashError(e.Key + " : " + e.Message)
			}
		}
	}

	this.SaveFlash()
	this.Data["Form"] = form
	this.Data["Title"] = "Register"
	this.TplNames = "auth/register.html"
}
