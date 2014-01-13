package controllers

import (
	"RaysGo/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/i18n"
	"time"
)

var (
	session_username string
	session_uid      int
	session_role_id  int
	session_email    string
)

const ViewFileExtension = ".html"

type NestPreparer interface {
        NestPrepare()
}

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           models.User
	isLogin        bool
	flash          *beego.FlashData   // flash data from the last action
	newFlash       *beego.FlashData 
	//i18n.Locale
}

type AuthController struct{
	BaseController
}

type AdminController struct {
	BaseController
}

func (this *BaseController) userSession() {
	session_username, _ = this.GetSession("username").(string)
	session_uid, _ = this.GetSession("userid").(int)
	session_role_id, _ = this.GetSession("userrole").(int)
	session_email, _ = this.GetSession("useremail").(string)
	this.isLogin = false
	if session_role_id == 0 {
		this.Data["UserId"] = 0
		this.Data["UserName"] = ""
		this.Data["UserRole"] = 0
		this.Data["UserEmail"] = ""
		this.Data["IsLogin"] = false
	} else {
		this.isLogin = true
		this.Data["IsLogin"] = true
		this.Data["UserId"] = session_uid
		this.Data["UserName"] = session_username
		this.Data["UserRole"] = session_role_id
		this.Data["UserEmail"] = session_email
		this.user = models.User{
			Id : int64(session_uid),
			Name : session_username,
			Rid : int64(session_role_id),
			Email : session_email,
		}
	}
}

func (this *BaseController) Prepare() {

	this.controllerName, this.actionName = this.GetControllerAndAction()
	this.userSession()

	this.flash = beego.ReadFromRequest(&this.Controller)

	this.Data["PageStartTime"] = time.Now()
	this.Layout = beego.AppConfig.String("defaultLayout")

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
    }
}

func (this *BaseController) User() models.User{
	return this.user
}

func (this *BaseController) FlashError(message string, args ...interface{}){
	this._flash("error", message, args...)
}

func (this *BaseController) FlashWarning(message string, args ...interface{}){
	this._flash("warning", message, args...)
}

func (this *BaseController) FlashNotice(message string, args ...interface{}){
	this._flash("notice", message, args...)
}

func (this *BaseController) getFlash() *beego.FlashData{
	if this.newFlash == nil{
		this.newFlash = beego.NewFlash()
	}
	return this.newFlash
}

func (this *BaseController) _flash(key string, message string, args ...interface{}){
	if key != "error" && key != "warning" && key != "notice"{
		return
	}

	flash := this.getFlash()

	if oldMessage, ok := flash.Data[key]; ok {
		message = oldMessage + "<br/>" + message
	}

	switch key{
	case "error":
		flash.Error(message, args...)
	case "waring":
		flash.Warning(message, args...)
	case "notice":
		flash.Notice(message, args...)
	}
}

func (this* BaseController) SaveFlash(){
	if flash := this.newFlash; flash!= nil{
		flash.Store(&this.Controller)
	}
}

func (this *BaseController) setMetas(metas map[string]string) {
	for key, value := range metas {
		this.Data[key] = value
	}
}

func (this *BaseController) setMeta(name string, content string) {
	this.Data[name] = content
}

func (this *BaseController) GoView(view ...string) {
	if len(view) > 1 {
		this.Layout = view[1] + ViewFileExtension
	}
	if len(view) > 0 {
		this.TplNames = view[0] + ViewFileExtension
	}
}

func (this *BaseController) GetParam(key string) string{
	return this.Ctx.Input.Param(key)
}

func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}
