package models

import (
       // "github.com/astaxie/beego/orm"
        "time"
)

// todo	
type Comment struct{
	Cid int
	Pid int `orm:"index"`
	Uid int `orm:"index"`
	Nid int `orm:"index"`
	Title string `orm:"size(255)"`
	Content string 
	Createtime time.Time `orm:"auto_now_add;type(datetime)"`
	Updatetime time.Time `orm:"auto_now_add;type(datetime)"`
	Host string `orm:"size(255)"`
	Name string `orm:"size(255)"`
	Email string `orm:"size(255)"`
	Homepage string `orm:"size(255)"`
}


func (m *Comment) TableName() string{
	return TableName("comment")
}

