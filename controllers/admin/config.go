package admin

import (
	"RaysGo/controllers"
	"RaysGo/models"
	"fmt"
	"github.com/astaxie/beego/validation"
)

type ConfigController struct {
	controllers.AdminController
}

func (this *ConfigController) Config() {
	this.Data["Form"] = models.LoadConfig()
	this.Data["Title"] = "Website configuration"
	this.TplNames = "admin/config.html"
}

func (this *ConfigController) ConfigPost() {
	var (
		form  models.ConfigForm
		valid validation.Validation
	)

	if err := this.ParseForm(&form); err != nil {
		fmt.Println(err)
	} else {
		if ok, ve := valid.Valid(form); ok && ve == nil {
			if form.Save() {
				this.FlashNotice("Configuration saved successfully.")
			} else {
				this.FlashError("Configuration saved failed.")
			}
		} else {
			for _, e := range valid.Errors {
				this.FlashError(e.Key + " : " + e.Message)
			}
		}
	}

	this.SaveFlash()
	this.Data["Form"] = form
	this.Data["Title"] = "Website configuration"
	this.TplNames = "admin/config.html"
}
