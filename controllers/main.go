package controllers

import (
	"RaysGo/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"

	this.setMetas(map[string]string{
		"Title":       "Home",
		"Author":      "Raysmond",
		"Description": "Simple blog system in Go!",
		"Keywords":    "RaysGo, Raysmond",
	})

	this.TplNames = "index.html"
}

func (this *MainController) About() {
	this.TplNames = "about.html"
}

// Get implemented contact page
func (this *MainController) Contact() {
	this.TplNames = "contact.html"
}

// Post implemented contact action
func (this *MainController) ContactPost() {
	form := models.ContactForm{}

	if e := this.ParseForm(&form); e == nil {
		// send contact message
		this.FlashNotice("Thanks for your contact! We'll reply you as soon as possible!")
	} else {
		this.FlashError("All field in contact form is required! Please check it!")
	}

	this.Redirect("/contact", 302)
}
