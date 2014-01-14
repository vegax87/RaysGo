package helpers

import (
	"github.com/astaxie/beego"
	"time"
)

func Loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

func SiteTitle(title string) string {
	return title + " - " + AppName
}

func AddViewFunc() {
	beego.AddFuncMap("loadtimes", Loadtimes)
	beego.AddFuncMap("site_title", SiteTitle)
	beego.AddFuncMap("markdown", MarkdownHtml)
}
