package controllers

import(
	"github.com/astaxie/beego/context"
)

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

