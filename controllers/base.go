package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare(){
	beego.AddFuncMap("loadtimes", loadtimes)
	
	this.Data["PageStartTime"] = time.Now()
}

func loadtimes(t time.Time) int {
        return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

