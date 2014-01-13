package helpers

import(
	"time"
	"github.com/astaxie/beego"
)

func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

func Init(){
	beego.AddFuncMap("loadtimes", loadtimes)
}