package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName string
}

func (this *BaseController) Prepare(){
	beego.AddFuncMap("loadtimes", loadtimes)
	
	this.controllerName, this.actionName = this.GetControllerAndAction()
	this.Data["PageStartTime"] = time.Now()
	this.Layout = beego.AppConfig.String("defaultLayout")
}

func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

func (this *BaseController) setMetas(metas map[string] string){
	for key, value := range metas{
		this.Data[key] = value
	}
}

