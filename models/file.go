package models

import (
        //"github.com/astaxie/beego/orm"
        "time"
)

// todo	
type File struct{
	Fid int
	Uid int `orm:"index"`
	Filename string `orm:"size(255)"`
	Uri string `orm:"size(255)"`
	Mimetype string `orm:"size(255)"`
	Filesize int
	status int
	timestamp time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *File) TableName() string{
	return TableName("file")
}

