package helpers

import(
	"time"
	"github.com/astaxie/beego"
)

func Loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

func ShowFlashError(flash *beego.FlashData) string{
	return ""
}

func AddViewFunc(){
	beego.AddFuncMap("loadtimes", Loadtimes)
	beego.AddFuncMap("show_flash", ShowFlashError)
}
