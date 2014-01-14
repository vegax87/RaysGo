package controllers

type AuthInterface interface {
	AuthActions() []string
}

func (this *AdminController) Get() {

	this.Data["Title"] = "Administration"
	this.TplNames = "admin/admin.html"
}

func (this *AdminController) AuthPrepare() {
	if !this.isLogin {
		this.Redirect("/login", 302)
		return
	}

	if !this.User().IsAdmin() {
		this.FlashError("You don't have the permission to access the page.")
		this.SaveFlash()
		this.Redirect("/", 302)
	}
}

func (this *AuthController) AuthPrepare() {
	var actions []string
	if app, ok := this.AppController.(AuthInterface); ok {
		actions = app.AuthActions()
	}
	if len(actions) > 0 {
		for _, val := range actions {
			if this.actionName == val {
				if !this.isLogin {
					this.FlashError("Please login first!")
					this.SaveFlash()
					this.Redirect("/login", 302)
					return
				} else {
					break
				}
			}
		}
	} else {
		if !this.isLogin {
			// abort all actions if no auth actions are declared
			this.FlashError("Please login first!")
			this.SaveFlash()
			this.Redirect("/login", 302)
		}
	}
}
