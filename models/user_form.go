package models

import (
	"github.com/astaxie/beego/validation"
)

// Login form
type LoginForm struct {
	UserName string `valid:"Required"`
	// Password string `form:"type(password)" valid:"Required"`
	Password string `valid:"Required"`
	//	Remember bool
}

// Register form
type RegisterForm struct {
	UserName   string `valid:"Required;AlphaDash;MinSize(5);MaxSize(30)"`
	Email      string `valid:"Required;Email;MaxSize(80)"`
	Password   string `valid:"Required;MinSize(4);MaxSize(30)"`
	PasswordRe string `valid:"Required;MinSize(4);MaxSize(30)"`
	//	Locale     i18n.Locale `form:"-"`
}

func (form *RegisterForm) Valid(v *validation.Validation) {

	if form.Password != form.PasswordRe {
		v.SetError("PasswordRe", "Password confirm is not matched with password field!")
		return
	}

	e1, e2, _ := CanRegistered(form.UserName, form.Email)

	if !e1 {
		v.SetError("UserName", "User name already exists.")
	}

	if !e2 {
		v.SetError("Email", "User email already exists.")
	}
}

// Forgot form
type ForgotForm struct {
	Email string `valid:"Required;Email;MaxSize(80)"`
	User  *User  `form:"-"`
}

func (form *ForgotForm) Valid(v *validation.Validation) {
	if HasUser(form.User, form.Email) == false {
		v.SetError("Email", "User email doesn't exists.")
	}
}

// Reset password form
type ResetPwdForm struct {
	Password   string `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
	PasswordRe string `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
}

func (form *ResetPwdForm) Valid(v *validation.Validation) {
	// Check if passwords of two times are same.
	if form.Password != form.PasswordRe {
		v.SetError("PasswordRe", "Passwords of two times are not matched.")
		return
	}
}
