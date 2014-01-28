package helpers

import (
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"time"
)

func Loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

func SiteTitle(title string) string {
	return title + " - " + AppName
}

func Pager(total int, page int, pagesize int) template.HTML {
	str := ""

	str += "<ul  class=\"pagination\">"
	prev, next, totalPage := page-1, page+1, total/pagesize
	if total%pagesize != 0 {
		totalPage++
	}
	if prev <= 0 {
		str += "<li class=\"disabled\"><a href=\"/admin/users?page=" + fmt.Sprintf("%d", prev) + "\">&laquo;</a></li>"
	} else {
		str += "<li><a href=\"/admin/users?page=" + fmt.Sprintf("%d", prev) + "\">&laquo;</a></li>"
	}
	for p := 0; p < totalPage; p++ {
		link := fmt.Sprintf("/admin/users?page=%d", p+1)
		if p == page-1 {
			str += "<li class=\"active\"><a href=\"" + link + "\">" + fmt.Sprintf("%d", p+1) + "</a></li>"
		} else {
			str += "<li><a href=\"" + link + "\">" + fmt.Sprintf("%d", p+1) + "</a></li>"
		}
	}
	if next > totalPage {
		str += "<li class=\"disabled\"><a href=\"/admin/users?page=" + fmt.Sprintf("%d", next) + "\">&laquo;</a></li>"
	} else {
		str += "<li><a href=\"/admin/users?page=" + fmt.Sprintf("%d", next) + "\">&laquo;</a></li>"
	}

	str += "</ul>"
	return beego.Str2html(str)
}

func AddViewFunc() {
	beego.AddFuncMap("loadtimes", Loadtimes)
	beego.AddFuncMap("site_title", SiteTitle)
	beego.AddFuncMap("markdown", MarkdownHtml)
	beego.AddFuncMap("pager", Pager)
}
