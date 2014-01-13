package controllers

import(
	"github.com/astaxie/beego/context"
)

type AuthInterface interface {
	AuthActions() []string
}

var FilterUser = func(ctx *context.Context) {
    _, ok := ctx.Input.Session("useruid").(int)
    if !ok {
        ctx.Redirect(302, "/login")
    }
}

var FilterAdmin = func(ctx *context.Context){
	_, okAdmin := ctx.Input.Session("userrole").(int)
	if !okAdmin {
		ctx.Redirect(302, "/")
	}
}

func (this *AuthController) NestPrepare(){
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
					this.Redirect("/login",302)
					return
				} else {
					break
				}
			}
		}
	} else {
		if !this.isLogin{
			// abort all actions if no auth actions are declared
			this.FlashError("Please login first!")
			this.SaveFlash()
			this.Redirect("/login",302)
		}
	}
}
