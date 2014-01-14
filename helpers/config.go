package helpers

import (
	"github.com/astaxie/beego"
)

var (
	AppName           string
	AppDescription    string
	AppKeywords       string
	AppAuthor         string
	AppVer            string
	AppHost           string
	AppUrl            string
	IsProMode         bool
	MailUser          string
	MailFrom          string
	ActiveCodeLives   int
	ResetPwdCodeLives int
	LoginRememberDays int
	Langs             []string
	DefaultLayout     string
)

func LoadConf() {
	AppName = beego.AppConfig.String("AppName")
	AppDescription = beego.AppConfig.String("AppDescription")
	AppKeywords = beego.AppConfig.String("AppKeywords")
	AppAuthor = beego.AppConfig.String("AppAuthor")
	AppVer = beego.AppConfig.String("AppVer")
	IsProMode = beego.AppConfig.String("runmode") != "dev"
	MailUser = beego.AppConfig.String("MailUser")
	MailFrom = beego.AppConfig.String("MailFrom")
	ActiveCodeLives, _ = beego.AppConfig.Int("ActiveCodeLives")
	ResetPwdCodeLives, _ = beego.AppConfig.Int("ResetPwdCodeLives")
	LoginRememberDays, _ = beego.AppConfig.Int("LoginRememberDays")
	DefaultLayout = beego.AppConfig.String("DefaultLayout")
}
