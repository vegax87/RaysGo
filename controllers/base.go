package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

var(
	session_username string
	session_uid int
	session_role_id int
	session_email string
)

const ViewFileExtension = ".tpl"

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (this *BaseController) userSession(){
	session_username, _ = this.GetSession("username").(string)
	session_uid, _ = this.GetSession("userid").(int)
	session_role_id, _ = this.GetSession("userrole").(int)
	session_email, _ = this.GetSession("useremail").(string)
	
	if session_role_id == 0 {
		this.Data["UserId"] = 0
        this.Data["UserName"] = ""
        this.Data["UserRole"] = 0
        this.Data["UserEmail"] = ""
        this.Data["IsLogin"] = false
	} else{
		this.Data["IsLogin"] = true
		this.Data["UserId"] = session_uid
        this.Data["UserName"] = session_username
        this.Data["UserRole"] = session_role_id
        this.Data["UserEmail"] = session_email
	}
}

func (this *BaseController) Prepare() {
	beego.AddFuncMap("loadtimes", loadtimes)
	
	this.controllerName, this.actionName = this.GetControllerAndAction()

	this.userSession()
	
	this.Data["PageStartTime"] = time.Now()
	this.Layout = beego.AppConfig.String("defaultLayout")
}

func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

func (this *BaseController) setMetas(metas map[string]string) {
	for key, value := range metas {
		this.Data[key] = value
	}
}

func (this *BaseController) GoView(view ...string) {
	if len(view) > 1 {
		this.Layout = view[1] + ViewFileExtension
	}
	if len(view) > 0 {
		this.TplNames = view[0] + ViewFileExtension
	}
}

