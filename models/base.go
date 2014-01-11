package models

import (
        "crypto/md5"
        "fmt"
        "github.com/astaxie/beego"
        "github.com/astaxie/beego/orm"
        _ "github.com/go-sql-driver/mysql"
)

type BaseModel struct{

}

func init() {
        dbhost := beego.AppConfig.String("dbhost")
        dbport := beego.AppConfig.String("dbport")
        dbuser := beego.AppConfig.String("dbuser")
        dbpassword := beego.AppConfig.String("dbpassword")
        dbname := beego.AppConfig.String("dbname")
        dbcharset := beego.AppConfig.String("dbcharset")
        if dbport == "" {
                dbport = "3306"
        }
        dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + dbcharset
        
        orm.RegisterDataBase("default", "mysql", dsn)
        

        orm.RegisterModel(new(User), new(UserRole), new(Variable))
}

func Md5(buf []byte) string {
        hash := md5.New()
        hash.Write(buf)
        return fmt.Sprintf("%x", hash.Sum(nil))
}

func TableName(str string) string {
        return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}